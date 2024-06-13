package request

type LoginDto struct {
	Id       int64  `json:"id"`
	UserName string `gorm:"user_name" json:"userName" valid:"Required; MaxSize(50)" `
	Password string `gorm:"password" json:"password"  valid:"Required; MaxSize(50)" `
}
