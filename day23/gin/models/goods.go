package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	ShopId     int    `json:"shop_id"`
	CategoryId int    `json:"category_id"`
	GoodName   string `json:"good_name"`
}
