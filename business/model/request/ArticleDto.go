package request

import (
	"blogServe/business/model"
	"github.com/astaxie/beego/validation"
)

type ArticleDto struct {
	model.Page
	model.Article
}

func (a *ArticleDto) Valid(v *validation.Validation) {
	if !v.Min(int(a.UserID), 1, "UserID").Ok {
		_ = v.SetError("UserID", "UserID cannot be zero")
	}
	if !v.Required(a, a.Content).Ok {
		_ = v.SetError("Content", "Content cannot be zero")
	}
}
