package addressController

import (
	"github.com/gin-gonic/gin"
	"hd-mall-ed/packages/client/models/addressModel"
	"hd-mall-ed/packages/client/pkg/app"
	"hd-mall-ed/packages/client/pkg/e"
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

}
