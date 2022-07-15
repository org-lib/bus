package main

import (
	"context"
	"fmt"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/mongodb"
	"github.com/org-lib/bus/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"log"
	"net/url"
)

func main() {

	//定义 cfg 对象
	var cfg *mongodb.Info
	cfg = &mongodb.Info{
		Host:          config.Config.V.GetString("mongodb.host"),
		Port:          config.Config.V.GetString("mongodb.port"),
		Username:      config.Config.V.GetString("mongodb.username"),
		Password:      url.QueryEscape(config.Config.V.GetString("mongodb.password")),
		DefaultAuthDB: config.Config.V.GetString("mongodb.defaultAuthDB"),
		Options:       config.Config.V.GetString("mongodb.replicaSet"),
	}

	//获取数据库实例连接
	db, err := mongodb.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("mongodb", config.Config.V.GetString("mongodb.port")))
		panic(err)
	}
	defer db.Disconnect(context.TODO())
	databases, err := db.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	//日志打印
	logger.Log.Info("server start", zap.String("mongodb", "运行结束！"))
	//其它操作方式不断迭代中...

}
