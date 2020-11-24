package web

import (
	"github.com/gin-gonic/gin"
	"go-micro-grpc/services"
)

func NewGinRouter(productService services.ProductService)*gin.Engine{
	r := gin.Default()
	r.Use(InitMiddleware(productService))
	vg := r.Group("v1")
	vg.Handle("POST", "/product", HandlerProductList)
	return  r
}
