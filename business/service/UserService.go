package service

import (
	"blogServe/business/global"
	"blogServe/business/model"
	"blogServe/business/model/request"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/copier"
)

type UserService struct{}

func (s *UserService) CreateUser(register request.RegisterUserDto) (id int64, err error) {
	valid := validation.Validation{}
	_, err = valid.Valid(&register)
	if err != nil {
		return 0, err
	}
	user := model.User{}
	err = copier.Copy(&register, &user)
	if err != nil {
		return 0, errors.New("copier.Copy failed" + err.Error())
	}
	db := global.DefaultDb
	r := db.Create(&user)
	if r.Error != nil {
		return 0, errors.New("CreateUser failed" + r.Error.Error())
	}
	return user.ID, nil
}

func (s *UserService) UpdateUser(user model.User) (id int64, err error) {
	db := global.DefaultDb
	r := db.Where("id = ?", user.ID)
	if r.Error != nil {
		return 0, errors.New("CreateUser failed" + r.Error.Error())
	}
	return user.ID, nil
}

func (s *UserService) CheckAuthUser(username, password string) (user *model.User, exist bool) {
	user, err := s.SelectByNameAndPass(username, password)
	if err != nil {
		return nil, false
	}
	if user.ID <= 0 {
		return nil, false
	}
	return user, true
}
func (s *UserService) SelectByNameAndPass(name string, pass string) (user *model.User, err error) {
	if name == "" || pass == "" {
		return nil, errors.New("用户名或密码不能为空")
	}
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Where("account  = ? AND password = ?", name, pass).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) SelectUserById(id int) (user *model.User, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Where("id = ", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *UserService) SelectUserListByName(name string) (user []*model.User, err error) {
	db := global.DefaultDb
	db = db.Where("del_flag = 0")
	err = db.Where("name = ?", name).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}
