package main

import (
	"context"
	"fmt"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/mongo"
	"github.com/org-lib/bus/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"log"
	"net/url"
)

func main() {

	//定义 cfg 对象
	var cfg *mongo.Info
	cfg = &mongo.Info{
		Host:          config.Config.V.GetString("mongo.host"),
		Port:          config.Config.V.GetString("mongo.port"),
		Username:      config.Config.V.GetString("mongo.username"),
		Password:      url.QueryEscape(config.Config.V.GetString("mongo.password")),
		DefaultAuthDB: config.Config.V.GetString("mongo.defaultAuthDB"),
		Options:       config.Config.V.GetString("mongo.replicaSet"),
	}

	//获取数据库实例连接
	db, err := mongo.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("mongo", config.Config.V.GetString("mysql.port")))
		panic(err)
	}
	defer db.Disconnect(context.TODO())
	databases, err := db.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	//日志打印
	logger.Log.Info("server start", zap.String("mongo", "运行结束！"))
	//其它操作方式不断迭代中...

}
