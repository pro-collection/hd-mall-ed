package database

import (
	"hd-mall-ed/packages/client/database/tableModal"
	"log"
)

//初始化表
func autoMigrateTable() {
	err := DataBase.AutoMigrate(&tableModal.User{})

	if err != nil {
		log.Fatalf("get DataBase.AutoMigrate() error: %v", err)
	}
}
