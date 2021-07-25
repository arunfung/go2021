package model

import "strconv"

type User struct {
	username string
	password string
	age      int
	sex      string
}

var userDatas map[string]Model

func NewUser() *User {
	return &User{}
}

func (u *User) SetUserName(username string)  {
	u.username = username
}

func (u *User) SetPassword(password string)  {
	u.password = password
}
func (u *User) SetAge(age int)  {
	u.age = age
}
func (u *User) SetSex(sex string)  {
	u.sex = sex
}

func (u *User) GetUserName() string  {
	return u.username
}
func (u *User) GetPassword() string  {
	return u.password
}
func (u *User) GetAge() int  {
	return u.age
}
func (u *User) GetSex() string  {
	return u.sex
}

func (u *User) ToString() string {
	return u.username + "," + strconv.Itoa(u.age) + "," + u.sex
}