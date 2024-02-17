package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lease/config"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	var err error

	host := config.Cfg.Section("MYSQL").Key("host").String()
	port := config.Cfg.Section("MYSQL").Key("port").String()
	user := config.Cfg.Section("MYSQL").Key("user").String()
	password := config.Cfg.Section("MYSQL").Key("password").String()
	database := config.Cfg.Section("MYSQL").Key("database").String()

	//  连接MySQL
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL初始化失败", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("MySQL配置初始化失败", err)
	}

	maxOpenConn, err := config.Cfg.Section("DATASOURCE").Key("max_open_conn").Int()
	if err != nil {
		log.Fatal("DataSource配置文件类型转换失败失败", err)
	}
	maxIdleConn, err := config.Cfg.Section("DATASOURCE").Key("max_idle_conn").Int()
	if err != nil {
		log.Fatal("DataSource配置文件类型转换失败失败", err)
	}

	//  设置最大连接数
	sqlDB.SetMaxOpenConns(maxOpenConn)
	//  设置最大空闲连接数
	sqlDB.SetMaxIdleConns(maxIdleConn)
}
