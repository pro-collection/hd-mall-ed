package addressController

import "hd-mall-ed/packages/common/pkg/e"

// @title     	更新地址
// @description
// @auth      	晴小篆  331393627@qq.com
// @param		addressController.handleUpdateAddressHelperOptions
// @return    	bool 如果是返回true, 说明有问题；返回 false 说明没有问题
func handleUpdateAddressHelper(options handleUpdateAddressHelperOptions) bool {
	api := options.api
	model := options.model
	params := options.updateParams
	var err error

	// 首先看是否是存在当前用户
	userId := api.GetUserId()
	err = model.FindAddressById(params.ID, uint(userId))
	if err != nil {
		//此时说明没有查询到相关地址， 或者是相关地址查询错误
		api.ResFail(e.AddressNotExist)
		return true
	}

	// 更新逻辑
	err = model.UpdateAddress(params)
	if err != nil {
		//此时说明没有查询到相关地址， 或者是相关地址查询错误
		api.ResFail(e.AddressUpdateFail)
		return true
	}
	return false
}
