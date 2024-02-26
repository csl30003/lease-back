package database

import (
	"github.com/go-redis/redis"
	"lease/config"
	"log"
)

var (
	Cache *redis.Client
)

func init() {
	var err error

	host := config.Cfg.Section("REDIS").Key("host").String()
	port := config.Cfg.Section("REDIS").Key("port").String()
	password := config.Cfg.Section("REDIS").Key("password").String()
	database, err := config.Cfg.Section("REDIS").Key("database").Int()
	if err != nil {
		log.Fatal("Redis配置文件类型转换失败失败", err)
	}

	Cache = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       database,
	})

	_, err = Cache.Ping().Result()
	if err != nil {
		log.Fatal("Redis初始化失败", err)
	}
}
