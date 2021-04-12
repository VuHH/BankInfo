package controllers

import (
	"bankinfo.com/model/dto"
	"bankinfo.com/service"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginDTO dto.Login
	var err = c.ShouldBindJSON(&loginDTO)
	httpStatus, message = service.LoginService(loginDTO, err)
	c.JSON(httpStatus, message)
}
