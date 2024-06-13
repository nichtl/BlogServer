package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	"errors"
)

type CategoryService struct{}

func (c *CategoryService) FindCategoryById(id int) (category *model.Category, err error) {
	if id <= 0 {
		return nil, errors.New("无效 id")
	}
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Where(" id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) FindFirstAll() (category []*model.Category, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Where("father_id = 0 ").Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) FindChildrenByFatherId(id int) (category []*model.Category, err error) {
	if id <= 0 {
		return nil, errors.New("无效 id")
	}
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
