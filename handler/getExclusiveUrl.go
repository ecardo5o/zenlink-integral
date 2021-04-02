package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"zenlink-integral/models"
)

///需要的参数有：
///1.用户地址
///2.当前的url
func getUrl(params []string, ctx *gin.Context) {
	if len(params) < 1{
		errorParams("not enough params", ctx)
		return
	}
	///1.获取第一个参数,检查格式
	userAddress := params[0]

	regular := "^5[a-zA-Z1-9]{47}"
	reg1 := regexp.MustCompile(regular)
	resultBool := reg1.MatchString(userAddress)
	if resultBool == false {
		errorParams("userAddress invalid!",ctx)
		return
	}

	///注册时判断当前用户是否已经是注册用户
	url,_ := dbService.FindUrl(userAddress)

	response := models.GenerateUrlResponse(url)
	ctx.JSON(http.StatusOK, response)
}
