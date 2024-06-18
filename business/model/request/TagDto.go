package request

import (
	"blogServe/business/model"
	"github.com/astaxie/beego/validation"
)

type TagDto struct {
	model.Page
	ID       int64  `json:"id" gorm:"column:id;type:bigint"`
	Name     string `gorm:"column:name;type:varchar(128);" json:"name"`
	FatherID int64  `gorm:"column:father_id;type:bigint;not null;default:0;" json:"fatherId"`
	Path     string `gorm:"column:path;type:varchar(128);" json:"path"`
}

func (u *TagDto) Valid(v *validation.Validation) {
	v.Required(u, "Name").Message("名称不能为空")
}
