package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	"blogServe/business/model/request"
	"errors"
	"github.com/jinzhu/copier"
)

type CategoryService struct{}

func (c *CategoryService) FindCategoryByID(id int) (category *model.Category, err error) {
	if id <= 0 {
		return nil, errors.New("无效 id")
	}
	db := global.DefaultDb
	err = db.Where(" id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) FindFirstAll() (category []*model.Category, err error) {
	db := global.DefaultDb
	err = db.Where("father_id = 0 ").Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) FindChildrenByFatherID(id int) (category []*model.Category, err error) {
	if id <= 0 {
		return nil, errors.New("无效 id")
	}
	db := global.DefaultDb
	err = db.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryService) CreateCategory(dto request.CategoryDto) (id int64, err error) {

	if dto.Name == "" {
		return 0, errors.New("name can not be empty")
	}
	category := &model.Category{}
	err = copier.CopyWithOption(&category, &dto, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return 0, err
	}
	db := global.DefaultDb
	err = db.Create(&category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (c *CategoryService) DelByID(id int64) (count int64, err error) {
	if id <= 0 {
		return 0, errors.New("id can not be empty")
	}
	db := global.DefaultDb

	r := db.Model(&model.Category{}).Where("id = ?", id).Delete(&model.Category{})

	if r.Error != nil {
		return 0, r.Error
	}
	count = r.RowsAffected
	return count, nil
}
