package apis

import (
	"blog/model"
	setting "blog/pkg"
	"blog/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 获取多个文章标签
func GetTags(ctx *gin.Context) {
	log.Printf("enter GetTags")
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
	data["lists"], _ = model.GetTags(util.GetPage(ctx), setting.PageSize, maps)
	data["total"] = model.GetTagTotal(maps)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// 获取单个文章标签
func GetTag(ctx *gin.Context) {
	log.Printf("enter GetTag")
	id, _ := strconv.Atoi(ctx.Query("id"))
	tag, _ := model.GetTag(id)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": tag,
	})
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
