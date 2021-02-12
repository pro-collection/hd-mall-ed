package database

import (
	"hd-mall-ed/packages/common/database/tableModel"
	"log"
)

//初始化表
func autoMigrateTable() {
	err := DataBase.AutoMigrate(&tableModel.User{})
	err = DataBase.AutoMigrate(&tableModel.Address{})
	err = DataBase.AutoMigrate(&tableModel.Category{})

	if err != nil {
		log.Fatalf("get DataBase.AutoMigrate() error: %v", err)
	}
}
