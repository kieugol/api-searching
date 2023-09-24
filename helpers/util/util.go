package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ShoudBindHeader(c *gin.Context) bool {
	platform := c.Request.Header.Get("X-PLATFORM")
	deviceType := c.Request.Header.Get("X-DEVICE-TYPE")
	deviceId := c.Request.Header.Get("X-DEVICE-ID")
	lang := c.Request.Header.Get("X-LANG")
	channel := c.Request.Header.Get("X-CHANNEL")

	if platform == "" || deviceType == "" || deviceId == "" || lang == "" || channel == "" {
		return false
	}

	return true
}

func LogPrint(jsonData interface{}) {
	prettyJSON, _ := json.MarshalIndent(jsonData, "", "")
	fmt.Printf("%s\n", strings.ReplaceAll(string(prettyJSON), "\n", ""))
}

func GetKeyFromContext(ctx context.Context, key string) (interface{}, bool) {
	if v := ctx.Value(key); v != nil {
		return v, true
	}

	return nil, false
}

func ParseJSON(data []byte, target interface{}, modelName string) error {
	if err := json.Unmarshal(data, target); err != nil {
		logrus.WithField("model", modelName).WithField("error", err.Error()).Error("PARSE_DATA_ERROR")
		return err
	}
	return nil
}

func ReadFile(path string) []byte {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	value, _ := ioutil.ReadAll(jsonFile)

	return value
}
