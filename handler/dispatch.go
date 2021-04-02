package handler

import (
	"github.com/gin-gonic/gin"
	"zenlink-integral/db/service/mysqlsrv"
	"zenlink-integral/models"
)

var dbService = mysqlsrv.MysqlService{}

var MapHandlers = map[string]func([]string, *gin.Context){
	"register":     	register,	///注册
	"getUrl":    		getUrl,		///获取专属推广url
	"getTestcoin": 		getTestcoin,///获取测试币
}

func RpcDispatch(c *gin.Context) {
	rpcParams := new(models.RpcQueryParams).Parse(c)
	handlerFunc,ok := MapHandlers[rpcParams.Method]
	if !ok {
		errorMethodNotExist(c)
		return
	}
	handlerFunc(rpcParams.Params, c)
}

