package productController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/admin/models/staticModel"
	"hd-mall-ed/packages/common/pkg/adminApp"
	"hd-mall-ed/packages/common/pkg/e"
)

func Create(c *gin.Context) {
	api := adminApp.ApiInit(c)
	model := &productModel.Product{}
	staticModelInstance := &staticModel.Static{}

	params := &createParamsStruct{}
	var err error

	err = c.ShouldBindJSON(params)

	// 参数绑定
	err = deepcopier.Copy(params).To(model)
	if api.ValidateHasError(model) {
		return
	}

	// todo 创建之前校验是是否有重复商品
	err = model.Create()
	if err != nil {
		api.ResFailMessage(e.Fail, err.Error())
		return
	}

	// 创建成功之后， 可以在这个地方添加各种场景的图片
	for i, _ := range params.ProductImageList {
		params.ProductImageList[i].ProductId = model.ID
	}

	for i, _ := range params.ProductDetailImageList {
		params.ProductDetailImageList[i].ProductId = model.ID
	}

	go staticModelInstance.CreateStatics(&params.ProductImageList)
	go staticModelInstance.CreateStatics(&params.ProductDetailImageList)



	api.ResponseNoData()
}
