package controllers

import (
	"../model/dto"
	"../service"
	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginDTO dto.Login
	var err = c.ShouldBindJSON(&loginDTO)
	httpStatus, message = service.LoginService(loginDTO, err)
	c.JSON(httpStatus, message)
}
