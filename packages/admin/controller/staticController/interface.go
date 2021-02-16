package staticController

import "hd-mall-ed/packages/admin/models/staticModel"

type createStaticsParams struct {
	List []staticModel.Static `json:"list" valid:"required"`
}
