package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	"blogServe/business/model/request"
	"errors"
	"github.com/jinzhu/copier"
)

type ArticleService struct{}

func (a *ArticleService) QueryArticleList(dto request.ArticleDto) (articles []*model.Article, total int64, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	if dto.Article.Title != "" {
		db = db.Where("title LIKE ?", "%"+dto.Article.Title+"%")
	}
	if dto.Article.UserID > 0 {
		db = db.Where("user_id = ?", dto.Article.UserID)
	}
	if dto.Article.UserAccount != "" {
		db = db.Where("user_account = ?", dto.Article.UserAccount)
	}
	if dto.Article.UpdateTime != nil {
		db = db.Where("update_time = ?", dto.Article.UpdateTime)
	}
	if dto.PageNum <= 0 || dto.PageSize <= 0 {
		dto.PageSize = global.NumTen
		dto.PageNum = global.NumOne
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

func (a *ArticleService) CreateArticle(dto request.ArticleDto) (id int64, err error) {

	db := global.DefaultDb
	article := &model.Article{}
	err = copier.CopyWithOption(&article, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return 0, err
	}
	r := db.Model(&article).Create(&article)
	if r != nil {
		return 0, r.Error
	}
	return article.ID, nil
}
func (a *ArticleService) UpdateArticle(dto request.ArticleDto) (count int64, err error) {
	var user model.User
	err = copier.CopyWithOption(&user, &dto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	db := global.DefaultDb
	r := db.Model(&user).Updates(user)
	if r.Error != nil {
		return 0, errors.New("UpdateUser failed" + r.Error.Error())
	}
	return r.RowsAffected, nil
}

func (a *ArticleService) DelByID(id int64) (count int64, err error) {
	article := model.Article{}
	article.ID = id
	db := global.DefaultDb
	err = db.Delete(&article).Error
	if err != nil {
		return 0, err
	}
	return 0, nil
}
