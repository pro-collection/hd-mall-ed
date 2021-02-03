package app

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/pkg/e"
	"net/http"
)

// 公用的返回
func Response(code int, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}
