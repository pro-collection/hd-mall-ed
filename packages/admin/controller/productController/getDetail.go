package productController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

func GetDetail(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	static := &staticModel.Static{}

	detail := &detailResponseStruct{}

	id, _ := strconv.Atoi(c.DefaultQuery("id", ""))
	err := model.GetById(id)
	if err != nil {
		api.ResFail(e.Fail)
		return
	}

	_ = deepcopier.Copy(model).To(detail)

	// 查询静态资源
	staticList := &[]staticModel.Static{}
	staticQueryMap := make(map[string]interface{})
	staticQueryMap["product_id"] = id
	staticList, _ = static.GetListByQuery(staticQueryMap)

	for index, static := range *staticList {
		if static.Type == 1 {
			detail.ProductImageList[index] = static
		}
		if static.Type == 2 {
			detail.ProductDetailImageList[index] = static
		}
	}

	api.Response(detail)
}
