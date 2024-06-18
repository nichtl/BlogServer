package request

type CategoryDto struct {
	FatherID int64  `json:"fatherId" gorm:"column:father_id;type:bigint;not null;default:0;comment:父id"`
	Name     string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:类别名称"`
	UserID   int64  `json:"user_id" gorm:"column:user_id;type:int;not null;comment:用户 id"`
}
