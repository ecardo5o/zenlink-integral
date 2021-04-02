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
func register(params []string, ctx *gin.Context) {
	if len(params) < 2{
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
	///2.获取第二个参数，
	///如果该url不为空，说明当前用户是被人邀请的
	///检查该url是否存在于数据库
	///如果该url为空，说明当前用户不是被人邀请的，为当前用户进行注册（如果已存在则不会进行注册）
	exclusiveUrl := params[1]

	if exclusiveUrl != "" {
		exist,err := dbService.IsExistUrl(exclusiveUrl)
		if err != nil || exist == false {
			///邀请url不存在，则视为用户不经过邀请注册
			exclusiveUrl = ""
		}
	}
	///注册时判断当前用户是否已经是注册用户
	registered:= dbService.Register(userAddress,exclusiveUrl)
	if !registered {
		errorDatabaseRetrieve("register failed",ctx)
		return
	}
	response := new(models.UserInfoResponse).Generate(registered)
	ctx.JSON(http.StatusOK, response)
}
