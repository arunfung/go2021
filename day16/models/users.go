package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id       int
	Username string `orm:"column(name)"`
	Password string
	Email    string
	Role     *UserRole `orm:"rel(fk)"` // 表示的是 users与role ; 一对一 => rel(fk) 定义的时候 -》查找Users表中的 role_id => 根据结构体的属性 跟上 id  正向 -》 从自身查找
}

type UserRole struct {
	Id   int    `orm:"pk"`
	Name string `orm:"column(name)"`
	Desc string `orm:"column(desc)"`
	// 反向一对多 beego-orm 根据 Users结构体查找 含有 `orm:"rel(fk)"`的属性，确定关联关系
	Users []*Users `orm:"reverse(many)"`
}

func init() {
	// register model
	orm.RegisterModel(new(Users), new(UserRole))
}

func FindUserByName(name string) (*Users, error) {
	o := orm.NewOrm()

	user := &Users{}
	o.QueryTable("Users").Filter("name", name).RelatedSel().One(user)
	return user, nil
}

func FindRoleById(id int) (*UserRole, error) {
	o := orm.NewOrm()

	role := &UserRole{}
	o.QueryTable("UserRole").Filter("id", id).One(role)
	o.LoadRelated(role, "Users")
	return role, nil
}
