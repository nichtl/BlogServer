package model

type Article struct {
	ExtendModel
	Title       string `gorm:"column:title; type:varchar(256); default:'';" json:"title" `
	Content     string `gorm:"column:content; type:longtext;" json:"content"  `
	Intro       string `gorm:"column:intro; type:varchar(512);" json:"intro"  `
	UserID      int64  `gorm:"column:user_id;type:bigint;" json:"userId"`
	UserAccount string `gorm:"column:user_account;type:varchar(128);" json:"userAccount" `
}

func (a Article) TableName() string { return "article" }
