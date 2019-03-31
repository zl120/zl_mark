package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"zl_mark/errs"
)

type UserAccount struct {
	Id          int    `orm:(id)`
	Name        string `orm:(name)`
	Pwd         string `orm:(pwd)`
	UserInfoId  int    `orm:(user_info_id)`
	Token       string `orm:(token)`
	Status      int    `orm:(status)`
	Description string `orm:(description)`
	Role        int    `orm:(role)`
}

func (m *UserAccount) TableName() string {
	return "user_account"
}

func (m *UserAccount) CreateUserAccount(xOrm orm.Ormer) (int, error) {
	if xOrm == nil {
		xOrm = orm.NewOrm()
	}
	id, err := xOrm.Insert(m)
	if err != nil {
		fmt.Println("err: ", err)
		return -1, err
	}
	return int(id), nil
}

func (m *UserAccount) Login(xOrm orm.Ormer) error {
	if xOrm == nil {
		xOrm = orm.NewOrm()
	}
	var user *[]orm.Params = &[]orm.Params{}
	num,err:=xOrm.Raw(`select * from user_account where name=? and pwd=?`,m.Name,m.Pwd).Values(user)
	if err != nil {
		fmt.Println("err: ",err)
		return errs.NewComplexErrs("failed",err.Error())
	}
	if num != 1 {
		return errs.NewComplexErrs("failed","username non-exist or wrong password! ")
	}
	return nil
}

func (m *UserAccount) IsRepeatName(xOrm orm.Ormer) (bool,error) {
	if xOrm == nil {
		xOrm = orm.NewOrm()
	}
	var user *[]orm.Params = &[]orm.Params{}
	fmt.Println(m.Name)
	num,err:=xOrm.Raw(`select * from user_account where name=? `,m.Name).Values(user)
	if err != nil {
		return true,errs.NewComplexErrs("failed",err.Error())
	}
	fmt.Println(num)
	if num != 0 {
		return true,nil
	}
	return false,nil
}