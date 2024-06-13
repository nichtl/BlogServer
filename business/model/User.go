package model

type User struct {
	ExtendModel
	RealName         string `json:"realName" gorm:"column:real_name;type:varchar(128);not null;default:'';comment:真实姓名"`
	UserType         int8   `json:"userType"  gorm:"column:user_type;type:int;not null;default:0;comment:用户类型"`
	NickName         string `json:"nickName" gorm:"column:nick_name;type:varchar(256);not null;default:'';comment:昵称"`
	Account          string `json:"account" gorm:"column:account;type:varchar(128);not null;default:'';comment:账户名"`
	Password         string `json:"password" gorm:"column:password;type:varchar(256);not null;default:'';comment:密码"`
	Email            string `json:"email" gorm:"column:email;type:varchar(128);not null;default:'';comment:邮箱"`
	NameIdentityCode string `json:"nameIdentityCode" gorm:"column:name_identity_code;type:varchar(128);not null;default:'';comment:用户身份码 userId 有可能变动 这个做备用"`
	ExtendInfo       string `json:"extendInfo" gorm:"column:extend_info;type:varchar(1024);not null;default:'';comment:扩展信息"`
	LinkAccount      string `json:"linkAccount" gorm:"column:link_account;type:varchar(256);not null;default:'';comment:关联账户"`
	Gender           int8   `json:"gender" gorm:"column:gender;type:tinyint(3);not null,default:0;comment:性别:1女,2男,0未知 "`
	Signature        string `json:"signature" gorm:"column:signature;type:varchar(256);not null;default:'';comment:个性签名"`
	WeChat           string `json:"weChat" gorm:"column:we_chat;type:varchar(128);not null;default:'';comment:微信"`
	RegisterTime     *Time  `json:"registerTime" gorm:"column:register_time;type:datetime;not null;default:current_timestamp;comment:注册时间"`
	Qq               int64  `json:"qq" gorm:"column:qq;type:bigint;not null;default:0;comment:qq"`
	Birthday         *Time  `json:"birthday" gorm:"column:birthday;type:date;not null;default:'1997-01-01';comment:生日"`
}

func (User) TableName() string {
	return "user"
}
