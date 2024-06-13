package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	Req "blogServe/business/model/request"
)

type TagService struct{}

func (t *TagService) QueryTagList(dto Req.TagDto) (tag []*model.Tag, total int64, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")

	if dto.Name == "" {
		db = db.Where("name = ?", dto.Name)
	}

	if dto.UpdateTime != nil {
		db = db.Where("update_time = ?", dto.UpdateTime)
	}

	if dto.AddTime != nil {
		db = db.Where("add_time = ?", dto.AddTime)
	}

	if dto.FatherId > 0 {
		db = db.Where("father_id = ?", dto.FatherId)
	}

	if dto.Path != "" {
		db = db.Where("path LIKE ?", "%"+dto.Path+"%")
	}

	db = db.Limit(dto.PageSize).Offset((dto.PageNum - 1) * dto.PageSize)
	db = db.Order("update_time DESC")
	err = db.Find(&tag).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Model(&model.Tag{}).Count(&total).Error

	if err != nil {
		return nil, 0, err
	}
	return tag, total, nil
}
