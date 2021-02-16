package database

import (
	"hd-mall-ed/packages/common/database/tableModel"
	"log"
)

//初始化表
func autoMigrateTable() {
	err := DataBase.AutoMigrate(
		&tableModel.User{},
		&tableModel.Address{},
		&tableModel.Category{},
		&tableModel.AdminUser{},
		&tableModel.Product{},
		&tableModel.ProductCategory{},
		&tableModel.Static{},
	)

	if err != nil {
		log.Fatalf("get DataBase.AutoMigrate() error: %v", err)
	}
}
