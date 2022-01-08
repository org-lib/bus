package main

import (
	"fmt"
	"github.com/openark/golib/sqlutils"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/mysql"
	"github.com/org-lib/bus/logger"
	"go.uber.org/zap"
)

func main() {

	//初始换数据库连接信息

	//定义 cfg 对象
	var cfg *mysql.Info

	cfg = &mysql.Info{
		Host:         config.Config.V.GetString("mysql.host"),
		Port:         config.Config.V.GetString("mysql.port"),
		Database:     "mysql",
		Username:     config.Config.V.GetString("mysql.username"),
		Password:     config.Config.V.GetString("mysql.password"),
		Timeout:      0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		Charset:      "",
	}
	cfg.Database = "mysql"

	//获取数据库实例连接
	db, err := mysql.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("mysql", "3306"))
		panic(err)
	}
	defer db.Close()

	//unix/linux 可用 sqlutils
	//sqlutils.QueryRowsMap 不支持在Windows 上运行，log 包异常
	err = sqlutils.QueryRowsMap(db, `select * from help_topic where name = ?`, func(m sqlutils.RowMap) error {
		logger.Log.Info("数据库返回信息：", zap.String("name", m.GetString("name")))
		logger.Log.Info("数据库返回信息：", zap.String("help_topic_id", m.GetString("help_topic_id")))
		logger.Log.Info("数据库返回信息：", zap.String("description", m.GetString("description")))
		return nil
	}, "HEX")
	if err != nil {
		logger.Log.Error(fmt.Sprintf("数据库查询，失败：%v", err), zap.String("mysql", "运行失败！"))
	}

	//show master status   ;  show slave status;
	err = sqlutils.QueryRowsMap(db, `SHOW MASTER STATUS;`, func(m sqlutils.RowMap) error {
		logger.Log.Info("数据库返回信息：", zap.String("File", m.GetString("File")))
		logger.Log.Info("数据库返回信息：", zap.String("Position", m.GetString("Position")))
		return nil
	})
	if err != nil {
		logger.Log.Error(fmt.Sprintf("数据库查询，失败：%v", err), zap.String("mysql", "运行失败！"))
	}

	//日志打印
	logger.Log.Info("server start", zap.String("mysql", "运行结束！"))
	//其它操作方式不断迭代中...

}
