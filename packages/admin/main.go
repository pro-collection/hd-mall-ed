package main

import (
	"hd-mall-ed/packages/admin/router"
	"hd-mall-ed/packages/common/config"
	"hd-mall-ed/packages/common/config/cache"
	"hd-mall-ed/packages/common/database"
)

func init() {
	config.SetUp()
	database.SetUp()
	cache.SetUp()
}

func main() {
	appRouter := router.SetUpRouter()
	_ = appRouter.Run(":8081")
}
