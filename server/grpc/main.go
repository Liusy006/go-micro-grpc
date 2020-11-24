package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"go-micro-grpc/serviceImpl"
	"go-micro-grpc/services"
)

func main(){
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	productService := micro.NewService(
		micro.Name("api.jmming.com.test"),
		micro.Registry(etcdReg),
		)
	fmt.Println(productService)
	services.RegisterProductServiceHandler(productService.Server(), new(serviceImpl.ProductServiceImpl))
	productService.Init()
	productService.Run()
}
