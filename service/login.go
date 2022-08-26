package service

import (
	"fmt"
	"gin-gorm-ol/helper"
	"gin-gorm-ol/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Login ログイン
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "please input the required fields",
		})
		return
	}

	// md5
	password = helper.GetMd5(password)
	fmt.Printf(password)
	var data models.UserBasic

	err := models.DB.Where("name = ? and password = ?", username, password).Find(&data).Error
	if err != nil {
		log.Println("Login Error: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "can not get the user info",
		})
		return
	}
	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		log.Println("Get Token Error: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "can not get the token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
	})
}
