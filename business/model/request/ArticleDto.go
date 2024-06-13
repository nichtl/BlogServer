package request

import "blogServe/business/model"

type ArticleDto struct {
	model.Page
	model.Article
}
