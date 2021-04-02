package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zenlink-integral/db"
	"zenlink-integral/handler"
	"zenlink-integral/middleware"
)

func main() {
	err := db.ConnectZenlinkDb()
	if err != nil{
		return
	}
	route := gin.Default()
	route.Use(middleware.Cors())

	route.POST("/rpc", handler.RpcDispatch)
	err = route.Run(":10004")
	if err != nil{
		fmt.Println(err)
	}
}