package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/coding-challenge/api-searching/config"
	"github.com/coding-challenge/api-searching/helpers/mytime"
	"github.com/coding-challenge/api-searching/middleware"
	"github.com/coding-challenge/api-searching/routes"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(middleware.Recovery())
}

func main() {
	env := flag.String("e", os.Getenv("APP_ENV"), "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*env)
	cfg := config.GetConfig()

	ok := mytime.SetTimezone(cfg.GetString("timezone"))
	if ok != nil {
		fmt.Println("Fatal error timezone", ok)
	}
	port := cfg.GetString("server.port")

	routes.RouteInit(engine)
	_ = engine.Run(":" + port)
}
