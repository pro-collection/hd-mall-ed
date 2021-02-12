package app

import (
	"github.com/asaskevich/govalidator"
	"hd-mall-ed/packages/common/pkg/e"
	"strings"
)

// 验证
func (api *ApiFunction) ValidateHasError(paramsStruct interface{}) bool {
	_, err := govalidator.ValidateStruct(paramsStruct)
	if err != nil {
		api.ResFailMessage(e.InvalidParams, strings.Split(err.Error(), ";")[0])
		return true
	}
	return false
}
