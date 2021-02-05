package database

import (
	"hd-mall-ed/packages/client/database/tableModel"
	"log"
)

//初始化表
func autoMigrateTable() {
	err := DataBase.AutoMigrate(&tableModel.User{})
	err = DataBase.AutoMigrate(&tableModel.Address{})

	if err != nil {
		log.Fatalf("get DataBase.AutoMigrate() error: %v", err)
	}
}
