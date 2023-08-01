package main

import (
	"fmt"
	"github.com/openark/golib/sqlutils"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/mysql"
	"github.com/org-lib/bus/logger"
	"go.uber.org/zap"
	"strings"
)

func main() {

	//初始换数据库连接信息

	//定义 cfg 对象
	var cfg mysql.Info

	cfg = mysql.Info{
		Host:         config.Config.V.GetString("mysql.host"),
		Port:         config.Config.V.GetString("mysql.port"),
		Database:     config.Config.V.GetString("mysql.database"),
		Username:     config.Config.V.GetString("mysql.username"),
		Password:     config.Config.V.GetString("mysql.password"),
		Timeout:      3000,
		ReadTimeout:  3000,
		WriteTimeout: 3000,
		Charset:      "utf8",
	}

	//获取数据库实例连接
	db, err := mysql.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("mysql", "3306"))
		panic(err)
	}
	defer db.Close()

	//unix/linux 可用 sqlutils
	//sqlutils.QueryRowsMap 不支持在Windows 上运行，log 包异常
	err = sqlutils.QueryRowsMap(db, `select * from sys_config where id = ?`, func(m sqlutils.RowMap) error {
		logger.Log.Info("数据库返回信息：", zap.String("config_value", m.GetString("config_value")))
		logger.Log.Info("数据库返回信息：", zap.String("description", m.GetString("description")))
		logger.Log.Info("数据库返回信息：", zap.String("saas_tenant_code", m.GetString("saas_tenant_code")))
		return nil
	}, "1")
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
	tmpResults, err := sqlutils.QueryNamedResultData(db, fmt.Sprintf("select id from trade_task limit 10"))
	// insert all
	// 根据 主键 批次查询 并发插入

	// insert 头部
	tmp := "insert into tempTable("
	for _, n := range tmpResults.Columns {
		tmp = tmp + n + ","
	}
	tmp = strings.TrimRight(tmp, ",")
	tmp = tmp + ")values"
	//insert 数据部分
	tmpv := ""
	logger.Log.Info(fmt.Sprintf("v-1=%v", len(tmpResults.Data)))
	for _, datum := range tmpResults.Data {
		tmpx := ""
		for _, data := range datum {
			if !data.Valid {
				tmpx = tmpx + "," + "NULL"

			} else {

				tmpx = tmpx + "," + "'" + data.String + "'"
			}
		}
		logger.Log.Info(fmt.Sprintf("v0=(%v)", strings.TrimLeft(tmpx, ",")))
		tmpv = tmpv + "," + "(" + strings.TrimLeft(tmpx, ",") + ")"
	}
	logger.Log.Info(fmt.Sprintf("v1=%v%v", tmp, strings.TrimLeft(tmpv, ",")))

	//其它操作方式不断迭代中...

}
