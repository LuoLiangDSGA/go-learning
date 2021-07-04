package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var sqlDB *gorm.DB

// 定义基础属性
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreateTime string `gorm:"create_time"`
	ModifyTime string `gorm:"modify_time"`
}

func Init() {
	//dbType := viper.GetString("db.dbType")
	user := viper.GetString("db.username")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	dbName := viper.GetString("db.name")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	var err error
	sqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db init error: ", err)
	}
	dbCfg, err := sqlDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbCfg.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbCfg.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbCfg.SetConnMaxLifetime(time.Hour)
}
