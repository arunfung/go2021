package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id int
	Username string `orm:"column(name)"`
	Password string
	Email string
}

type UserRole struct {
	Id int    `orm:"pk"`
	Name  string `orm:"column(name)"`
	Desc  string `orm:"column(desc)"`
	// 反向一对多 beego-orm 根据 Users结构体查找 含有 `orm:"rel(fk)"`的属性，确定关联关系
	//Users []*Users `orm:"reverse(many)"`
}

func init() {
	// register model
	orm.RegisterModel(new(Users), new(UserRole))
}

func FindUserByName(name string) (*Users,error) {
	o := orm.NewOrm()

	user := &Users{}
	o.QueryTable("Users").Filter("name",name).One(user)
	return user,nil
}
