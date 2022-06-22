package routes

import (
	"net/http"
	"preview/app/conf"

	"github.com/gin-gonic/gin"
)

func DefaultRoute(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.GET("/get", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    "请求成功",
				"data":   "请求成功-GET",
			})
		})

		api.POST("/post", func(c *gin.Context) {
			json := conf.Default{}
			c.Bind(&json)

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    "请求成功-POST",
				"data": conf.Default{
					Test:  "模拟数据1",
					Other: "模拟数据2",
				},
			})

		})

	}
}
