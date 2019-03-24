package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
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
