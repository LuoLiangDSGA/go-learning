package model

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *sql.DB

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
	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db init error: ", err)
	}
	db, err := sqlDB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)
}

func Close() {
	defer db.Close()
}