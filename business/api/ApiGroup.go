package api

import (
	base "blogServe/business/api/BaseApi"
)

type ApiGroup struct {
	base.BaseAPI
	base.CategoryAPI
	base.UserAPI
	base.ArticleAPI
	base.TagAPI
}

var ApiGroupApp = new(ApiGroup)
