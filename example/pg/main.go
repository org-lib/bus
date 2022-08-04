package main

import (
	"fmt"
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
	//db = db.Raw("select * from t_store_template_task limit 10").Find(&users)
	for i, user := range users {
		fmt.Println("第", i, "个 User：", user.CDate)
	}
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
