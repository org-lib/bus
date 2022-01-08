package main

import (
	"fmt"
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
	//windows
	rows, err := db.Query("select name,help_topic_id,description from help_topic where name = ?", "HEX")
	if err != nil {
		logger.Log.Error(fmt.Sprintf("数据获取，失败：%v", err), zap.String("mysql", "3306"))
		panic(err)
	}

	for rows.Next() {
		var s1 string
		var s2 int
		var s3 string
		err = rows.Scan(&s1, &s2, &s3)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("数据库查询，失败：%v", err), zap.String("mysql", "运行失败!"))
			panic(err)
		}
		logger.Log.Info("数据库返回信息：", zap.String("name", s1))
		logger.Log.Info("数据库返回信息：", zap.Int("help_topic_id", s2))
		logger.Log.Info("数据库返回信息：", zap.String("description", s3))
	}
	rows.Close()

	//日志打印
	//logger.Log.Info("server start", zap.String("mysql", "运行结束！"))

	//其它操作方式不断迭代中...

	//修改、删除以及ddl
	//rs, err := db.Exec("INSERT INTO test.hello(world) VALUES ('hello world')")
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//rowCount, err := rs.RowsAffected()
	//if err != nil{
	//	log.Fatalln(err)
	//}
	//logger.Log.Info("inserted rows：", zap.Int("rowCount", rowCount))

}
