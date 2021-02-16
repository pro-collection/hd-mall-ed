package staticModel

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type Static tableModel.Static

func (static *Static) Create() error {
	return database.DataBase.Create(static).Error
}

// 插入多个
func (static *Static) CreateStatics(statics *[]Static) error {
	return database.DataBase.Create(statics).Error
}

func (static *Static) Update() error {
	return database.DataBase.Where("id = ?", static.ID).Updates(*static).Error
}

// 按照条件获取
func (*Static) GetListByQuery(query interface{}) (*[]Static, error) {
	list := &[]Static{}
	err := database.DataBase.Where(query).Find(list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (static *Static) GetAllList() (*[]Static, error) {
	list := &[]Static{}
	err := database.DataBase.Find(list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func (static *Static) Delete() error {
	return database.DataBase.Delete(static).Error
}

// 按照条件批量删除
func (*Static) Deletes(query map[string]interface{}) error {
	return database.DataBase.Where(query).Delete(&Static{}).Error
}
