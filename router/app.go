package router

import (
	"gin-gorm-ol/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	// 問題リストを取得する
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// 利用者情報を取得する
	r.GET("/user-detail", service.GetUserDetail)

	// サブミットログ情報を取得する
	r.GET("/submit-list", service.GetSubmitList)

	// login
	r.POST("/login", service.Login)

	return r
}
