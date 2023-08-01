package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// MySQL 初始化连接信息对象
type Info struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	Timeout      int
	ReadTimeout  int
	WriteTimeout int
	Charset      string
}

//?timeout=3000ms&readTimeout=3000ms&writeTimeout=3000ms&charset=utf8

// Open 获取一个数据库连接
func Open(cnf Info) (*sql.DB, error) {
	if strings.Trim(cnf.Database, "") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	if strings.Trim(cnf.Charset, "") == "" {
		return nil, errors.New("*** 请至指定一个数据库字符集")
	}

	// 初始化超时参数值

	if cnf.Timeout <= 0 {
		cnf.Timeout = 3000
	}
	if cnf.ReadTimeout <= 0 {
		cnf.ReadTimeout = 3000
	}
	if cnf.WriteTimeout <= 0 {
		cnf.WriteTimeout = 3000
	}

	//

	path := strings.Join([]string{
		cnf.Username, ":",
		cnf.Password,
		"@tcp(", cnf.Host, ":", fmt.Sprintf("%v", cnf.Port), ")/",
		cnf.Database, fmt.Sprintf("?timeout=%dms&readTimeout=%dms&writeTimeout=%dms&charset=%v", cnf.Timeout, cnf.ReadTimeout, cnf.WriteTimeout, cnf.Charset)}, "")

	return sql.Open("mysql", path)
}
