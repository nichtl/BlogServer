package api

import (
	"blogServe/business/service"
)

type BaseAPI struct{}

var (
	userService     = service.UserService{}
	categoryService = service.CategoryService{}
	articleService  = service.ArticleService{}
	tagService      = service.TagService{}
)
