package services

import (
	"context"
	"encoding/json"
	"fmt"
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
	jsonBytes, _ := json.Marshal(goods)
	fmt.Println("GetGoodsDetail")

	res.Code = 200
	res.Msg = "成功"
	res.Data = jsonBytes
	return nil
}
