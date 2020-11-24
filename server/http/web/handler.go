package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-grpc/services"
)

func HandlerProductList(ctx *gin.Context){
	productService := ctx.Keys[PRODUCTSERVICE].(services.ProductService)
	req := services.ProductRequest{}
	err := ctx.Bind(&req)
	if err != nil{
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
		return
	}
	rsp, err := productService.GetProductList(context.Background(), &req)

	if err != nil{
		ctx.JSON(500, gin.H{
			"status": err.Error(),
		})
		return
	}else{
		ctx.JSON(200, gin.H{
			"data": rsp.Data,
		})
	}



}
