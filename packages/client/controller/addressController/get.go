package addressController

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/common/database/tableModel"
	"hd-mall-ed/packages/admin/models/addressModel"
	"hd-mall-ed/packages/common/pkg/app"
	"hd-mall-ed/packages/common/pkg/e"
	"strconv"
)

// 通过 userid 查找所有的address
func GetAddressList(c *gin.Context) {
	api := app.ApiFunction{C: c}

	model := &addressModel.Address{}
	// 获取userId
	user, err := api.GetUser()
	if err != nil {
		api.ResFail(e.FailLoginStatus)
		return
	}

	addressList, err2 := model.FindAddressByUserId(user.ID)
	if err2 != nil {
		api.ResFailMessage(e.Fail, err2.Error())
		return
	}

	api.Response(addressList)
}

func GetAddressById(c *gin.Context) {
	api := app.ApiFunction{C: c}
	var err error

	// 获取参数
	id := c.DefaultQuery("id", "")
	addressId, _ := strconv.Atoi(id)

	// 获取user 信息
	user, _ := api.GetUser()
	model := &addressModel.Address{}
	err = model.FindAddressById(uint(addressId), user.ID)
	if err != nil {
		api.ResFail(e.FindAddressFail)
		return
	}

	address := tableModel.AddressBase{}
	_ = deepcopier.Copy(model).To(&address)
	api.Response(address)
}
