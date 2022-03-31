package services

import "frame/models"

func GetGoodsDetail(where map[string]interface{}) *models.Goods {
	var goods models.Goods
	models.DB().Debug().Where(where).Unscoped().Find(&goods)
	return &goods
}
