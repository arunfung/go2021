package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dsn = "root:root123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4"
)

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	fmt.Println("err : ", err)
	return db
}
