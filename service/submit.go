package service

import (
	"gin-gorm-ol/define"
	"gin-gorm-ol/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetSubmitList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefauPage))
	size, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultSize))
	if err != nil {
		log.Println("GetSubmitList page strconv Error: ", err)
		return
	}
	page = (page - 1) * size
	var count int64
	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))
	var data []models.SubmitBasic

	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("Get Submit List Error: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取提交信息错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": count,
		},
	})
}
