package controllers

import (
	"net/http"

	"github.com/coding-challenge/api-searching/helpers/respond"
	request "github.com/coding-challenge/api-searching/request/user"
	"github.com/coding-challenge/api-searching/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserSrv services.IUserService
}

func NewUserController(userSrv services.IUserService) *UserController {
	return &UserController{
		UserSrv: userSrv,
	}
}

func (userCtrl UserController) Detail(c *gin.Context) {
	var req request.DetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, respond.MissingParams())
		return
	}
	data, sttCode := userCtrl.UserSrv.HandleDetail(req)

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

func (userCtrl UserController) PublishTopic(c *gin.Context) {

}
