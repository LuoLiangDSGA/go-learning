package router

import (
	"blog/router/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/ping", apis.Ping) // 服务检测接口

	tagApi := r.Group("/api")
	{
		tagApi.GET("/tags", apis.GetTags)
		tagApi.GET("/tag", apis.GetTag)
		tagApi.POST("/tag", apis.AddTag)
		tagApi.PUT("/tag", apis.UpdateTag)
		tagApi.DELETE("/tag", apis.DeleteTag)
	}

	return r
}
