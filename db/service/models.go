package service

type UserInfo struct {
	ID 						int		`gorm:"primaryKey;autoIncrement"`
	UserAddress             string 	`gorm:"uniqueIndex:address"`
	ParentAddress           string 	`gorm:"type:char(255)"`
	Integral_c_url          int
	Integral_invited_number int
	Integral_get_testcoin   int
	Get_testcoin_version    int
	Integral_dex_op         int
	Call_url                int
	Get_testcoin            int
	Dex_op                  int
	Time                    string 	`gorm:"type:char(255)"`
}
