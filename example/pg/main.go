package main

import (
	"fmt"
	"github.com/openark/golib/sqlutils"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/postgresql"
	"github.com/org-lib/bus/logger"
	"go.uber.org/zap"
	"time"
)

//  关于Gorm执行原生SQL
// **********语句字段要小写************
// ***********查询用db.Raw,其他用db.Exec
// *********** 字段大小写要对应上 **************
// *************** 注意要 defer rows.Close()

func main() {
	//定义 cfg 对象
	var cfg *postgresql.Info

	cfg = &postgresql.Info{
		Host:     config.Config.V.GetString("postgres.host"),
		Port:     config.Config.V.GetString("postgres.port"),
		Database: config.Config.V.GetString("postgres.database"),
		Username: config.Config.V.GetString("postgres.username"),
		Password: config.Config.V.GetString("postgres.password"),
		SslMode:  "disable",
		TimeZone: "Asia/Shanghai",
	}
	//获取数据库实例连接
	db, err := postgresql.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("postgres", "5432"))
		panic(err)
	}
	var users []TStoreTemplateTask
	//   查询 执行用Scan 和Find 一样
	db = db.Raw("select create_date,id,end_date,order_code,relate_id,request_data,response_date,retry_time,sku_code,status,type,execute_date from t_store_template_task limit 10").Scan(&users)
	for i, user := range users {
		fmt.Println("第", i, "个 User：", user.CDate)
	}

	db2, err := postgresql.OpenPlus(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("postgres", "5432"))
		panic(err)
	}
	defer db2.Close()

	//unix/linux 可用 sqlutils
	//sqlutils.QueryRowsMap 不支持在Windows 上运行，log 包异常

	//（sqlutils postgresql, clickhouse QueryRowsMap.arg 参数形式是$num）
	//（sqlutils mysql QueryRowsMap.arg 参数形式是?）

	err = sqlutils.QueryRowsMap(db2, `select * from t_store_template_task where id = $1 limit 1`, func(m sqlutils.RowMap) error {
		logger.Log.Info("数据库返回信息1：", zap.String("response_date", m.GetString("response_date")))
		logger.Log.Info("数据库返回信息2：", zap.String("order_code", m.GetString("order_code")))
		logger.Log.Info("数据库返回信息3：", zap.String("create_date", m.GetString("create_date")))
		return nil
	}, 11)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("数据库查询，失败：%v", err), zap.String("postgres2", "运行失败！"))
	}

	/*
		##测试null和空字符串
			CREATE TABLE testa(
			"id" bigserial NOT NULL,
			"name" char(50) NULL,
			PRIMARY KEY(id)
			);
			insert into testa values(1,'1');
			insert into testa values(2,'2');
			insert into testa values(3,'');
			insert into testa values(4,null);
			insert into testa values(5,null);
	*/
}

type TStoreTemplateTask struct {
	Id           int
	OrderCode    string
	RelateId     int8
	RequestData  string
	ResponseDate string
	RetryTime    int
	SkuCode      string
	Status       int
	Type         int
	ExecuteDate  time.Time
	CDate        time.Time `gorm:"column:create_date;type:TIMESTAMP"`
	EndDate      time.Time
}
type TestA struct {
	Id   int
	Name string `gorm:"column:name;type:char"`
}

/*
    1 | id            | int8      |      8 |        -1 | t       |
    2 | create_date   | timestamp |      8 |         6 | f       |
    3 | end_date      | timestamp |      8 |         6 | f       |
    4 | order_code    | varchar   |     -1 |       259 | f       |
    5 | relate_id     | int8      |      8 |        -1 | f       |
    6 | request_data  | varchar   |     -1 |      2052 | f       |
    7 | response_date | varchar   |     -1 |       259 | f       |
    8 | retry_time    | int4      |      4 |        -1 | f       |
    9 | sku_code      | varchar   |     -1 |       259 | f       |
   10 | status        | int4      |      4 |        -1 | f       |
   11 | type          | int4      |      4 |        -1 | f       |
   12 | execute_date  | timestamp |      8 |         6 | f       |
*/
