package controllers

import (
	"net/http"

	"github.com/coding-challenge/api-searching/helpers/api"
	"github.com/coding-challenge/api-searching/helpers/respond"
	request "github.com/coding-challenge/api-searching/request/user"
	"github.com/coding-challenge/api-searching/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (userCtrl UserController) Detail(c *gin.Context) {
	var req request.DetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	userCtrl.UserService = services.NewUserService(api.NewHttClient())

	data, sttCode := userCtrl.UserService.HandleDetail(req)

	switch sttCode {
	case http.StatusNotFound:
		c.JSON(sttCode, respond.NotFound())
		return
	case http.StatusInternalServerError:
		c.JSON(sttCode, respond.InternalServerError())
		return
	}

	c.JSON(http.StatusOK, respond.Success(data, "Success"))
}
