package productController

import (
	"hd-mall-ed/packages/admin/models/productModel"
	"hd-mall-ed/packages/admin/models/staticModel"
)

type createParamsStruct struct {
	productModel.Product
	ProductImageList []staticModel.Static `json:"product_image_list"`
	ProductDetailImageList []staticModel.Static `json:"product_detail_image_list"`
}

type deleteParamsStruct struct {
	Id int `json:"id" valid:"require"`
}
