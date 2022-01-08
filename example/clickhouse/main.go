package main

import (
	"fmt"
	"github.com/openark/golib/sqlutils"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/clickhouse"
	"github.com/org-lib/bus/logger"
	"go.uber.org/zap"
)

func main() {

	//初始换数据库连接信息

	//定义 cfg 对象
	var cfg *clickhouse.Info

	cfg = &clickhouse.Info{
		Host:     config.Config.V.GetString("clickhouse.host"),
		Port:     config.Config.V.GetString("clickhouse.port"),
		Database: config.Config.V.GetString("clickhouse.database"),
		Username: config.Config.V.GetString("clickhouse.username"),
		Password: config.Config.V.GetString("clickhouse.password"),
		Debug:    true,
	}

	//获取数据库实例连接
	db, err := clickhouse.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("clickhouse", "8123"))
		panic(err)
	}
	defer db.Close()

	//unix/linux 可用 sqlutils
	//sqlutils.QueryRowsMap 不支持在Windows 上运行，log 包异常
	err = sqlutils.QueryRowsMap(db, `select * from tablename where key = ?`, func(m sqlutils.RowMap) error {
		logger.Log.Info("数据库返回信息：", zap.String("_time", m.GetString("_time")))
		logger.Log.Info("数据库返回信息：", zap.String("instanceid", m.GetString("instanceid")))
		logger.Log.Info("数据库返回信息：", zap.String("projectName", m.GetString("projectName")))
		return nil
	}, "key-value")

	if err != nil {
		logger.Log.Error(fmt.Sprintf("数据库查询，失败：%v", err), zap.String("mysql", "运行失败！"))
	}
	//日志打印
	logger.Log.Info("server start", zap.String("mysql", "运行结束！"))
	//其它操作方式不断迭代中...

}
