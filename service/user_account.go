package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"zl_mark/models"
)

type UserAccountService struct {
	Id          int
	Name        string
	Pwd         string
	UserInfoId  int
	Token       string
	Status      int
	Description string
	Role        int
}

func (s *UserAccountService) CreateUserAccount() (int,error) {
	var xOrm orm.Ormer = orm.NewOrm()
	var user *models.UserAccount = &models.UserAccount{}
	user.Name = s.Name
	user.Pwd = s.Pwd
	user.Role = s.Role
	id,err:=user.CreateUserAccount(xOrm)
	if err != nil {
		fmt.Println("err: ",err)
		return -1,err
	}
	return id,nil
}
