package app

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/pkg/e"
	"net/http"
)

// 公用的返回
func (api *ApiFunction) Response(data interface{}) {
	api.C.JSON(http.StatusOK, gin.H{
		"code":    e.Success,
		"message": e.GetMsg(e.Success),
		"data":    data,
	})
}

// 公共返回没有data的场景
func (api *ApiFunction) ResponseNoData() {
	api.C.JSON(http.StatusOK, gin.H{
		"code":    e.Success,
		"message": e.GetMsg(e.Success),
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
