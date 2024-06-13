package api

import (
	base "blogServe/business/api/BaseApi"
)

type ApiGroup struct {
	base.BaseApi
	base.CategoryApi
	base.UserApi
	base.ArticleApi
	base.TagApi
}

var ApiGroupApp = new(ApiGroup)
