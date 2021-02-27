package main

import (
	"fmt"
	"hd-mall-ed/packages/common/config"
	"hd-mall-ed/packages/common/database"
)

func init() {
	config.SetUp()

}

func main() {
	database.SetUp()
	sqlDb, err := database.DataBase.DB()

	err = sqlDb.Close()
	if err != nil {
		fmt.Println("初始化数据库失败")
	}
	fmt.Println("初始化数据库成功")
}
