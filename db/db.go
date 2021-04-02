package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"zenlink-integral/config"
)

var db *gorm.DB

func setup() {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=true",
		config.DbCfg.DbUser,
		config.DbCfg.DbPassword,
		config.DbCfg.DbHost,
		config.DbCfg.DbPort)

	dbCfg := mysql.New(mysql.Config{
		DSN:                       dbURI, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	conn, _ := gorm.Open(dbCfg, &gorm.Config{})
	db = conn
}

func ConnectZenlinkDb() error{
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		config.DbCfg.DbUser,
		config.DbCfg.DbPassword,
		config.DbCfg.DbHost,
		config.DbCfg.DbPort,
		config.DbCfg.DbName)

	dbCfg := mysql.New(mysql.Config{
		DSN:                       dbURI, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // Disable color
		},
	)

	conn, err := gorm.Open(dbCfg, &gorm.Config{Logger:newLogger})
	if err != nil || conn == nil {
		fmt.Println("connect db server failed.")
		return errors.New("connect db server failed")
	}
	sqlDB, err := conn.DB()
	if err != nil || sqlDB == nil{
		fmt.Println("connect db server failed.")
		return errors.New("connect db server failed")
	}
	sqlDB.SetMaxIdleConns(config.DbCfg.DbMinConn)
	sqlDB.SetMaxOpenConns(config.DbCfg.DbMinConn)
	sqlDB.SetConnMaxLifetime(time.Second * config.DbCfg.DbConnDuration)
	db = conn
	return nil
}

func GetDB() *gorm.DB {
	if db == nil{
		setup()
	}
	return db
}