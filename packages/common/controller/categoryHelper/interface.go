package categoryHelper

import "hd-mall-ed/packages/common/database/tableModel"

type ListItemResultStruct struct {
	ID              uint                       `json:"id" gorm:"primarykey"`
	Name            string                     `json:"name" gorm:"not null;comment:'商品分类'" valid:"required~缺少商品类别名称"`
	Type            int                        `json:"type" gorm:"comment:'商品分类的类型'"`
	Avatar          string                     `json:"avatar" gorm:"comment:'缩略图''"`
	BannerImage     string                     `json:"banner_image" gorm:"content:banner图片;type:longtext"`
	BannerProductId uint                       `json:"banner_product_id" gorm:"content:banner对应的商品id, 可以实现跳转"`
	BannerLink      string                     `json:"banner_link" gorm:"content:banner对应商品的跳转链接"`
	Children        []*tableModel.CategoryBase `json:"children"`
}
