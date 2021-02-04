package app

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/pkg/e"
	"net/http"
)

type ApiFunction struct {
	C *gin.Context
}

// 公用的返回
func (api *ApiFunction) Response(code int, data interface{}) {
	api.C.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
		"data":    data,
	})
}

// 失败的返回
func (api *ApiFunction) ResFail(code int) {
	api.C.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": e.GetMsg(code),
	})
}

// 返回失败
func (api *ApiFunction) ResFailMessage(code int, message string) {
	api.C.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}
