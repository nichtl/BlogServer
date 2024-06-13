package request

import "blogServe/business/model"

type TagDto struct {
	model.Page
	model.Tag
}
