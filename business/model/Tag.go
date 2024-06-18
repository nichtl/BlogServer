package model

type Tag struct {
	ExtendModel

	Name     string `gorm:"column:name;type:varchar(128);" json:"name"`
	FatherID int64  `gorm:"column:father_id;type:bigint;not null;default:0;" json:"fatherId"`
	Path     string `gorm:"column:path;type:varchar(128);" json:"path"`
}
