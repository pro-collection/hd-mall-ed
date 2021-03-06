package addressController

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/addressModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"strings"
)

func Create(c *gin.Context) {
	api := app.ApiFunction{C: c}
	model := &addressModel.Address{}
	var err error

	_ = c.ShouldBindJSON(model)
	_, err = govalidator.ValidateStruct(model)
	if err != nil {
		api.ResFailMessage(e.Error, strings.Split(err.Error(), ";")[0])
		return
	}

	user, _ := api.GetUser()
	model.UserId = user.ID
	addressList, _ := model.FindAddressByUserId(user.ID)
	if len(*addressList) > 20 {
		api.ResFail(e.CreateAddressLimit)
		return
	}

	err = model.CreateAddress()
	if err != nil {
		api.ResFail(e.CreateAddressFail)
		return
	}

	api.ResponseNoData()
}
