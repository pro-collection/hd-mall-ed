package main

import (
	"hd-mall-ed/packages/client/config"
	"hd-mall-ed/packages/client/config/cache"
	"hd-mall-ed/packages/client/database"
	"hd-mall-ed/packages/client/router"
)

func init() {
	config.SetUp()
	database.SetUp()
	cache.SetUp()
}

func main() {
	appRouter := router.SetUpRouter()
	_ = appRouter.Run()
}
