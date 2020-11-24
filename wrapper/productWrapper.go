package wrapper

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"go-micro-grpc/services"
)

func defaultProductRsp(rsp interface{}) {
	data := rsp.(*services.ProductResponse)
	data.Data = make([]*services.ProductModel, 0)
	data.Data = append(data.Data, &services.ProductModel{
		Id:   123,
		Name: "123",
	})
}

type ProductWrapper struct{
	client.Client
}

func NewProductWrapper(client2 client.Client)client.Client{
	return &ProductWrapper{client2}
}

func( w *ProductWrapper)Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	cmdName := req.Service()+"."+req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout: 1000,
	}

	hystrix.ConfigureCommand(cmdName, config)
	return  hystrix.Do(cmdName, func() error {
		return  w.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		defaultProductRsp(rsp)
		return nil
	})
}
