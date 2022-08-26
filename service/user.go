package service

import (
	"gin-gorm-ol/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail 利用者情報を取得する
func GetUserDetail(c *gin.Context) {
	userIdentity := c.Query("user_identity")
	if userIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "user_identity can not be empty",
		})
		return
	}
	var data models.UserBasic
	err := models.DB.Omit("password").Where("identity = ?", userIdentity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "can not find user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
