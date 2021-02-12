package adminApp

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/common/pkg/app"
)

func ApiInit(c *gin.Context) *ApiFunction {
	return &ApiFunction{app.ApiFunction{C: c}}
}
