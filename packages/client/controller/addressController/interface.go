package addressController

import (
	"hd-mall-ed/packages/client/models/addressModel"
	"hd-mall-ed/packages/common/pkg/app"
)

type handleUpdateAddressHelperOptions struct {
	api          *app.ApiFunction
	model        *addressModel.Address
	updateParams *addressModel.UpdateRequestParamsStruct
}
