package model

type TagMapper struct {
	ExtendModel
	MapperType int   `json:"mapper_type" gorm:"column:mapper_type;type:int;"`
	ArticleId  int64 `json:"article_id" gorm:"column:article_id;type:bigint;"`
}
