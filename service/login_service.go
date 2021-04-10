package service

import (
	"../model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginService(loginDTO dto.Login, err error) (int, interface{}) {

	if err == nil {
		var isUser = checkLoginInfo(loginDTO)
		if isUser == false {
			return 500, gin.H{"message": "Account not exist"}
		}
		token := GenerateToken(loginDTO.Username, isUser)
		if token != "" {
			return http.StatusOK, gin.H{"token": token}
		} else {
			return http.StatusUnauthorized, gin.H{"message": "Permission is denied"}
		}
	} else {
		//log.Fatal(err)
		return 500, gin.H{"error": err.Error()}
	}

}

func checkLoginInfo(loginInfo dto.Login) bool {
	if loginInfo.Username == "user1" && loginInfo.Password == "123" || loginInfo.Username == "admin" && loginInfo.Password == "admin" {
		return true
	} else {
		return false
	}
}
