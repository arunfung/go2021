package services

import (
	"context"
	"frame/models"
	proto "frame/proto"
)

type RequestGoods struct {
}

func (r *RequestGoods) GetGoodsDetail(ctx context.Context, req *proto.RequestGoods, res *proto.ResponseGoods) error {
	where := map[string]interface{}{
		"id": req.GoodsId,
	}
	var goods models.Goods
	models.DB().Debug().Where(where).Unscoped().Find(&goods)
	data := &proto.GoodsInfo{
		ShopId:     goods.ShopId,
		CategoryId: goods.CategoryId,
		GoodName:   goods.GoodName,
	}

	res.Code = 200
	res.Msg = "成功"
	res.Data = data
	return nil
}
