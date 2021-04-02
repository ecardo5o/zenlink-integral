package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var DbCfg DbConfig
var LogCfg LogConfig

type ChainCfg struct {
	ChainId uint32   `json:"chain_id"`
	Type    string   `json:"type"`
	Nodes   []string `json:"nodes"`
}

type DbConfig struct {
	DbUser         string        `json:"db_user"`
	DbPassword     string        `json:"db_password"`
	DbHost         string        `json:"db_host"`
	DbPort         uint16        `json:"db_port"`
	DbName         string        `json:"db_name"`
	DbMaxConn      int           `json:"db_max_conn"`
	DbMinConn      int           `json:"db_min_conn"`
	DbConnDuration time.Duration `json:"db_conn_duration"`
	Chains         []ChainCfg    `json:"chains"`
	TokenList      []Token       `json:"token_list"`
}

type LogConfig struct {
	Logger        string `json:"logger"`
	LogLevel      int    `json:"log_level"`
	Project       string `json:"project"`
	Name          string `json:"name"`
	LogDir        string `json:"log_dir"`
	MaxDay        int    `json:"max_day"`
	RotateSeconds int    `json:"rotate_seconds"`
	Extname       string `json:"extname"`
}

type Token struct {
	ChainId     uint32 `json:"chain_id"`
	ModuleIndex uint8  `json:"module_index"`
	AssetIndex  uint32 `json:"asset_index"`
	Decimal     uint8  `json:"decimal"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
}

func init() {
	err := initDbConfigFromFile("./config/dbconfig.json")
	if err != nil {
		panic(err)
	}
	err = initLogConfigFromFile("./config/logconfig.json")
	if err != nil {
		panic(err)
	}
}

func initDbConfigFromFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(content, &DbCfg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func initLogConfigFromFile(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(content, &LogCfg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
