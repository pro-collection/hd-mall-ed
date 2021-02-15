package static

import (
	"hd-mall-ed/packages/common/database"
	"hd-mall-ed/packages/common/database/tableModel"
)

type Static tableModel.Static

func (static *Static) Create() error {
	return database.DataBase.Create(static).Error
}

func (static *Static) Update() error {
	return database.DataBase.Where("id = ?", static.ID).Updates(*static).Error
}

func (static *Static) GetListByQuery(query map[string]string) (*[]Static, error) {
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
