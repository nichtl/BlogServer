package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	Req "blogServe/business/model/request"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/copier"
)

type TagService struct{}

func (t *TagService) QueryTagList(dto Req.TagDto) (tag []*model.Tag, total int64, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")

	if dto.Name == "" {
		db = db.Where("name = ?", dto.Name)
	}

	//if dto.UpdateTime != nil {
	//	db = db.Where("update_time = ?", dto.UpdateTime)
	//}
	//
	//if dto.AddTime != nil {
	//	db = db.Where("add_time = ?", dto.AddTime)
	//}

	if dto.FatherID > 0 {
		db = db.Where("father_id = ?", dto.FatherID)
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

func (t *TagService) UpdateTagByID(dto Req.TagDto) (count int64, err error) {
	if dto.Name == "" || dto.ID <= 0 {
		return 0, errors.New("tag name or id can not empty")
	}
	tag := model.Tag{}
	_ = copier.CopyWithOption(&tag, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	db := global.DefaultDb
	r := db.Model(&tag).Updates(tag)
	if r.Error != nil {
		return 0, errors.New("CreateUpdateTagUser failed" + r.Error.Error())
	}
	return r.RowsAffected, nil
}

func (t *TagService) DelTagByID(id int64) (count int64, err error) {

	tag := model.Tag{}
	tag.ID = id
	db := global.DefaultDb
	err = db.Delete(&tag).Error
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (t *TagService) CreateTag(dto Req.TagDto) (id int64, err error) {
	valid := validation.Validation{}
	_, err = valid.Valid(&dto)
	if err != nil {
		return 0, err
	}

	tag := model.Tag{}
	err = copier.CopyWithOption(&tag, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return 0, errors.New("copier.Copy failed" + err.Error())
	}
	db := global.DefaultDb
	r := db.Create(&tag)
	if r.Error != nil {
		return 0, errors.New("CreateUser failed" + r.Error.Error())
	}
	return tag.ID, nil
}
