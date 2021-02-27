package bannerController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

// 获取所有 banner
func GetBannerList(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &staticModel.Static{}
	query := map[string]string{
		"type": "3", // banner 主图类型
	}
	list, err := model.GetListByQuery(query)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	api.Response(list)
}
