package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"zenlink-integral/models"
	"zenlink-integral/utils"
)

func getTestcoin(params []string, ctx *gin.Context) {
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

	///获取测试币时判断当前用户是否已经是注册用户
	url,_ := dbService.FindUrl(userAddress)
	if url == "" {
		///表示该用户未注册
		errorParams("this user cannot register", ctx)
		return
	}

	utils.FaucetTestcoin(userAddress)
	//if !result {
	//	errorDatabaseRetrieve("get testcoin failed,please retry!",ctx)
	//	return
	//}
	response := new(models.UserInfoResponse).Generate(true)
	ctx.JSON(http.StatusOK, response)
}