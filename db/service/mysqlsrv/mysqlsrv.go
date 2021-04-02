package mysqlsrv

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"zenlink-integral/config"
	"zenlink-integral/db"
	"zenlink-integral/db/service"
)

type MysqlService struct {
}
///获取自己的推广url
func (srv *MysqlService) FindUrl(userAddress string) (string,error){
	///通过address查询url
	var exclusiveUrl string
	var id int
	result := db.GetDB().Table("zenlink_user_info").Select("id").Where("user_address = ?", userAddress).Scan(&id)
	if result.Error != nil {
		return "",result.Error
	}
	if result.RowsAffected == 0 {
		exclusiveUrl = ""
	}else{
		index := 1000000 + id
		exclusiveUrl = "https://test.zenlink-dex.pro/?slkkne=" + strconv.Itoa(index)
	}
	///TODO：需要知道网址
	return exclusiveUrl, result.Error
}
///检查对应地址的url是否正确（该方法是为了防范不存在的分享url被用户访问和使用）
func (srv *MysqlService) IsExistUrl(exclusiveUrl string) (bool,error) {
	///查询指定url是否存在
	///解析出id
	index := strings.Index(exclusiveUrl,"=")
	id,_ := strconv.Atoi(exclusiveUrl[index+1:])
	id = id-1000000
	exist := false
	result := db.GetDB().Table("zenlink_user_info").Select("*").Where("id = ?", id)
	if result.Error == nil {
		exist = true
	}
	return exist, result.Error
}
///用户注册
func (srv *MysqlService) Register(userAddress string,exclusiveUrl string) bool{
	///注册用户，注册前检查用户是否已注册
	url,_ := srv.FindUrl(userAddress)
	if url != "" {
		///表示该用户已注册
		return false
	}

	index := strings.Index(exclusiveUrl,"=")
	id,_ := strconv.Atoi(exclusiveUrl[index+1:])
	id = id-1000000

	u := service.UserInfo{
		UserAddress:   userAddress,
		ParentAddress: strconv.Itoa(id),
		Time:    time.Now().String(),
	}
	///注册用户
	result := db.GetDB().Table("zenlink_user_info").Create(&u)
	isSuccess := result.Error == nil
	return isSuccess
}
///========================================================================================================================
func (srv *MysqlService) ConnectDataBase() error {
	res := db.GetDB().Exec("CREATE DATABASE " + config.DbCfg.DbName)
	if res.Error != nil {
		fmt.Println(res)
	}
	res = db.GetDB().Exec("USE " + config.DbCfg.DbName)
	if res.Error != nil {
		fmt.Println(res)
	}
	return nil
}