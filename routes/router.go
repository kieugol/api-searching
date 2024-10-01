package routes

import (
	"net/http"

	"github.com/coding-challenge/api-searching/services"
	"github.com/gin-gonic/gin"

	"github.com/coding-challenge/api-searching/controllers"
	"github.com/coding-challenge/api-searching/helpers/api"
	"github.com/coding-challenge/api-searching/middleware"
)

func RouteInit(engine *gin.Engine) {
	var c *gin.Context
	userSrv := services.NewUserService(c, api.NewHttClient())
	userCtr := controllers.NewUserController(userSrv)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API fetching data")
	})
	engine.Use(middleware.Recovery()) // Customize panic error

	apiV1 := engine.Group("/v1")
	apiV1.Use(middleware.ValidateHeader())
	apiV1.Use(middleware.RequestLog()) // format log request -response
	{
		apiV1.GET("/users/:id", userCtr.Detail)
		apiV1.GET("/users-v2/:id", userCtr.Detail)
		apiV1.GET("/detail/:id", userCtr.Detail)
		apiV1.GET("/detail-v2/:id", userCtr.Detail)
	}
}
