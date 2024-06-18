package request

import (
	"blogServe/business/model"
	"github.com/astaxie/beego/validation"
	"regexp"
)

type RegisterUserDto struct {
	RealName         string      `json:"realName" comment:"真实姓名" `
	UserType         int8        `json:"userType" `
	NickName         string      `json:"nickName" comment:"昵称" `
	Account          string      `json:"account" comment:"账户名"  `
	Password         string      `json:"password" comment:"密码 9-30个字符 。"`
	Email            string      `json:"email" comment:"邮箱" `
	NameIdentityCode string      `json:"nameIdentityCode" comment:"用户身份码 userId 有可能变动 这个做备用"`
	ExtendInfo       string      `json:"extendInfo" comment:"扩展信息"`
	LinkAccount      string      `json:"linkAccount" comment:"关联账户"`
	Gender           int8        `json:"gender" comment:"性别:0女,1男,2未知 "`
	Signature        string      `json:"signature" comment:"个性签名"`
	WeChat           string      `json:"weChat" comment:"微信"`
	RegisterTime     *model.Time `json:"registerTime" comment:"注册时间"`
	Qq               int64       `json:"qq" comment:"qq"`
	Birthday         *model.Time `json:"birthday" comment:"生日"`
}

func (u *RegisterUserDto) Valid(v *validation.Validation) {
	if u.Email != "" {
		v.Email(u, "Email").Message("邮箱格式有误")
	}
	v.Required(u, "Account").Message("账户不能为空")
	v.MinSize(u.Account, 6, "Account").Message("长度必须在6到30之间")
	v.MaxSize(u, 30, "Account").Message("长度必须在6到30之间")
	v.Required(u, "Password").Message("密码不能为空")
	v.MinSize(u, 9, "Password").Message("长度必须在9到30之间")
	v.MaxSize(u, 30, "Password").Message("长度必须在9到30之间")
	v.Match(u, regexp.MustCompile("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\\\d)(?=.*[^\\\\da-zA-Z]).+$"), "Password").Message("包含至少一个小写字母、一个大写字母、一个数字和一个特殊字符")

}
