package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	"blogServe/business/model/request"
)

type ArticleService struct{}

func (a *ArticleService) QueryAritcleList(dto request.ArticleDto) (articles []*model.Article, total int64, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	if dto.Article.Title != "" {
		db = db.Where("title LIKE ?", "%"+dto.Article.Title+"%")
	}
	if dto.Article.UserId > 0 {
		db = db.Where("user_id = ?", dto.Article.UserId)
	}
	if dto.Article.UserAccount != "" {
		db = db.Where("user_account = ?", dto.Article.UserAccount)
	}
	if dto.Article.UpdateTime != nil {
		db = db.Where("update_time = ?", dto.Article.UpdateTime)
	}
	if dto.PageNum <= 0 || dto.PageSize <= 0 {
		dto.PageSize = global.NUM_TEN
		dto.PageNum = global.NUM_ONE
	}
	db = db.Limit(dto.PageSize).Offset((dto.PageNum - 1) * dto.PageSize)
	db = db.Order("update_time DESC")
	err = db.Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Model(&articles).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}
