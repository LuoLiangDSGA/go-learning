package apis

import (
	"blog/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 获取多个文章标签
func GetTags(ctx *gin.Context) {
	log.Printf("enter getTags")
	name := ctx.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	state := -1
	if arg := ctx.Query("state"); arg != "" {
		maps["state"] = state
	}
	data["lists"] = model.GetTags(1, 1, maps)
	data["total"] = 0
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// 获取单个文章标签
func GetTag(ctx *gin.Context) {

}

// 添加单个标签
func AddTag(ctx *gin.Context) {

}

// 删除单个标签
func DeleteTag(ctx *gin.Context) {

}

// 更新单个标签
func UpdateTag(ctx *gin.Context) {

}
