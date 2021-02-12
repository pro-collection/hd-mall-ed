package router

import "github.com/gin-gonic/gin"

// @title     	商品分类路由
// @description 不需要分类
// @auth      	晴小篆  331393627@qq.com
// @param
// @return
func categoryRouter(router *gin.RouterGroup) {
	category := router.Group("/category")
	{
		category.POST("/create")
	}
}
