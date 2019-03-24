package service

import (
	"github.com/astaxie/beego/orm"
	"zl_mark/models"
)

type UserInfoService struct {
	Id          int
	Name        string
	Age         int
	IdCard      string
	Gender      string
	Token       string
	Status      int
	Description string
	Tel         string
}

func (s *UserInfoService) IsDataValid() error {
	return nil
}

func (s *UserInfoService) CreateUserInfo() (int,error) {
	var xOrm orm.Ormer = orm.NewOrm()
	var user *models.UserInfo = &models.UserInfo{}
	if err:=s.IsDataValid();err!=nil {
		return -1,err
	}
	user.Name = s.Name
	user.Age = s.Age
	user.Gender = s.Gender
	user.IdCard = s.IdCard
	user.Tel = s.Tel
	id,err:=user.CreateUserInfo(xOrm)
	if err != nil {
		return -1,err
	}
	return id,nil
}
