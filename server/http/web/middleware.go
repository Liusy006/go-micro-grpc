package web

import (
	"github.com/gin-gonic/gin"
	"go-micro-grpc/services"
)

const (
	PRODUCTSERVICE = "productservice"
)

func InitMiddleware(productService services.ProductService)gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys[PRODUCTSERVICE] = productService
		ctx.Next()
	}
}
