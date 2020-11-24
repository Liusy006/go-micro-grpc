package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	web2 "go-micro-grpc/server/http/web"
	"go-micro-grpc/services"
	"go-micro-grpc/wrapper"
)

func main(){
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	productClient := micro.NewService(
		micro.Name("product.client"),
		micro.WrapClient(wrapper.NewProductWrapper))
	productService := services.NewProductService("product_service", productClient.Client())
	httpService := web.NewService(
		web.Name("http_produce_service"),
		web.Handler(web2.NewGinRouter(productService)),
		web.Registry(etcdReg),
		web.Address(":8801"),
		web.Metadata(map[string]string{"protocol" : "http"}),
	)

	httpService.Init()
	httpService.Run()
}
