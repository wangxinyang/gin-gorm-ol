package service

import (
	"gin-gorm-ol/define"
	"gin-gorm-ol/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList 問題リストを取得する
func GetProblemList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefauPage))
	size, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList page strconv Error: ", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	list := make([]models.Problem, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Offset(page).Omit("content", "deleted_at").Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error: ", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取问题错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail 問題情報を取得する
func GetProblemDetail(c *gin.Context) {
	problemIdentity := c.Query("problem_identity")
	if problemIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	var data models.Problem
	err := models.DB.Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("identity = ?", problemIdentity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题数据不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get Problem Detail Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
