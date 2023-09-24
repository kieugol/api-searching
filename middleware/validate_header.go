package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-challenge/api-searching/helpers/respond"
	"github.com/coding-challenge/api-searching/helpers/util"
)

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.ShoudBindHeader(c) {
			c.JSON(http.StatusBadRequest, respond.MissingHeader())
			c.Abort()
			return
		}

		c.Next()
	}
}
