package api

import (
	"blogServe/business/service"
)

type BaseApi struct{}

var (
	userService     = service.UserService{}
	categoryService = service.CategoryService{}
	articleService  = service.ArticleService{}
	tagService      = service.TagService{}
)
