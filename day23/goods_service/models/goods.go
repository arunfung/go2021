package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	ShopId     int64  `json:"shop_id"`
	CategoryId int64  `json:"category_id"`
	GoodName   string `json:"good_name"`
}
