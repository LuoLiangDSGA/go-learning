package util

import (
	setting "blog/pkg"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		panic("page param error")
	}
	result = (page - 1) * setting.PageSize
	return result
}
