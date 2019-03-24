package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id          int    `orm:(id)`
	Name        string `orm:(name)`
	Age         int    `orm:(age)`
	IdCard      string `orm:(string)`
	Gender      string `orm:(gender)`
	Token       string `orm:(token)`
	Status      int    `orm:(status)`
	Description string `orm:(description)`
	Tel         string `orm:(tel)`
}

func (m *UserInfo) CreateUserInfo(xOrm orm.Ormer) (int, error) {
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
