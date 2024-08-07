package api

import (
	"blogServe/business/global"
	"blogServe/business/model"
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"blogServe/business/utils"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

type UserAPI struct{}

func (base *BaseAPI) RegisterUser(c *gin.Context) {
	var registerDto Req.RegisterUserDto
	err := c.ShouldBindJSON(&registerDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	validate := validation.Validation{}
	vr, _ := validate.Valid(&registerDto)
	if !vr {
		Res.ErrorWithMsg("invalid param", c)
		return
	}

	id, err := userService.CreateUser(registerDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	user, err := userService.SelectUserByID(int(id))
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	data := map[string]interface{}{
		"user": user,
	}
	Res.OkData(data, c)
}

func (base *BaseAPI) Login(c *gin.Context) {
	var loginDto Req.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&loginDto)
	if !ok {
		Res.ErrorWithMsg("invalid param", c)
	}

	user, err := userService.SelectByNameAndPass(loginDto.UserName, loginDto.Password)

	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}

	token, err := base.CreateToken(strconv.Itoa(int(user.ID)), loginDto.UserName, loginDto.Password)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
	}
	data := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	Res.OkData(data, c)
}

func (base *BaseAPI) CreateToken(id, userName, password string) (token string, err error) {
	user, isExist := userService.CheckAuthUser(userName, password)
	if !isExist {
		return "", errors.New("not exist user")
	}
	token, err = utils.GenerateToken(id, userName, password, user.Account)
	if err != nil {
		return "", err
	}
	err = utils.Set(global.TokenPrefix+id, token, global.NumFive, time.Hour)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (base *BaseAPI) UpdateUser(c *gin.Context) {
	var loginDto Req.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	if loginDto.ID <= 0 {
		Res.ErrorWithMsg("invalid param", c)
		return
	}
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	var user model.User
	err = copier.CopyWithOption(&user, &loginDto, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	count, err := userService.UpdateUser(user)

	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	Res.OkData(count, c)
}
func (base *BaseAPI) LogoutByUser(c *gin.Context) {
	var loginDto Req.LoginDto
	err := c.ShouldBindJSON(&loginDto)
	if loginDto.ID <= 0 {
		Res.ErrorWithMsg("invalid param", c)
		return
	}
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	user, err := userService.SelectUserByID(int(loginDto.ID))
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	if user == nil || user.Account == "" {
		Res.ErrorWithMsg("用户账户不存在", c)
		return
	}
	_ = utils.Del(global.TokenPrefix + strconv.Itoa(int(user.ID)))
	Res.Ok(c)
}
