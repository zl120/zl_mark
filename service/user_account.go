package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"zl_mark/errs"
	"zl_mark/models"
	"zl_mark/utils"
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
	var err error
	//check data, repeat name,valid name and password
	if err = utils.IsUserNameValid(s.Name);err != nil {
		return -1, errs.NewComplexErrs("failed", err.Error())
	}
	if err = utils.IsPwdValid(s.Pwd);err!=nil{
		return -1, errs.NewComplexErrs("failed", err.Error())
	}
	var isRepeat bool
	if isRepeat,err = s.IsRepeatName();err!=nil {
		return -1, errs.NewComplexErrs("failed", err.Error())
	}
	if isRepeat {
		return -1, errs.NewComplexErrs("failed","user name already exist")
	}
	var xOrm orm.Ormer = orm.NewOrm()
	var user *models.UserAccount = &models.UserAccount{}
	user.Name = s.Name
	user.Pwd = s.Pwd
	user.Role = s.Role
	id,err:=user.CreateUserAccount(xOrm)
	if err != nil {
		fmt.Println("err: ",err)
		return -1,errs.NewComplexErrs("failed",err.Error())
	}
	return id,nil
}

func (s *UserAccountService) IsRepeatName() (bool,error) {
	var user *models.UserAccount = &models.UserAccount{Name:s.Name,Pwd:s.Pwd}
	return user.IsRepeatName(nil)
}

func (s *UserAccountService) Login() error {
	var user *models.UserAccount = &models.UserAccount{Name:s.Name,Pwd:s.Pwd}
	if err := user.Login(nil);err != nil {
		return err
	}
	return nil
}
