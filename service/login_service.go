package service

import (
	"../model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginService(c *gin.Context) {
	var loginDTO dto.Login
	if err := c.ShouldBindJSON(&loginDTO); err == nil {
		var isUser = checkLoginInfo(loginDTO)
		token := GenerateToken(loginDTO.Username, isUser)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	} else {
		//TODO log error err.Error()
		c.JSON(500, gin.H{"error": "System error"})
	}

}

func checkLoginInfo(loginInfo dto.Login) bool {
	if loginInfo.Username == "user1" && loginInfo.Password == "123" || loginInfo.Username == "admin" && loginInfo.Password == "admin" {
		return true
	} else {
		return false
	}
}
