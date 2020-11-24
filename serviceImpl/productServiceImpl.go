package serviceImpl

import (
	"context"
	"go-micro-grpc/services"
	"strconv"
	"strings"
	"time"
)

type  ProductServiceImpl struct{

}

func NewProductServiceImpl(){

}

func (p *ProductServiceImpl) GetProductList(ctx context.Context, in *services.ProductRequest, out *services.ProductResponse) error {
	time.Sleep(time.Second* 3)
	ret := make([]*services.ProductModel, 0)
	for i := 0; i < int(in.Size); i++{
		var sb strings.Builder
		sb.WriteString("product"+strconv.Itoa(100+i))
		ret = append(ret, &services.ProductModel{
			Id:   int32(100+i),
			Name: sb.String(),
		})

		out.Data = ret
	}
	return nil
}
