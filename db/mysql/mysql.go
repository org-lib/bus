package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//MySQL 初始化连接信息对象
type Info struct {
	Host			string
	Port			string
	Database		string
	Username		string
	Password		string
	Timeout			int
	ReadTimeout		int
	WriteTimeout	int
	Charset			string
}

//获取一个数据库连接
func Open(cnf *Info) (*sql.DB, error) {
	if strings.Trim(cnf.Database ,"") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	path := strings.Join([]string{
		cnf.Username, ":",
		cnf.Password,
		"@tcp(", cnf.Host, ":", fmt.Sprintf("%v", cnf.Port), ")/",
		cnf.Database, "?timeout=3000ms&readTimeout=3000ms&writeTimeout=3000ms&charset=utf8"}, "")
	return sql.Open("mysql", path)
}