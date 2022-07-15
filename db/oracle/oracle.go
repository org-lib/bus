package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/godror/godror"
	"strings"
)

// Info oracle 初始化连接信息对象
type Info struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

// Open 获取一个数据库连接
func Open(cnf *Info) (*sql.DB, error) {
	if strings.Trim(cnf.Database, "") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	// 用户名/密码@IP:端口/实例名
	tmpInfo := fmt.Sprintf("%s/%s@%s:%d/%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.Database)
	db, err := sql.Open("godror", tmpInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
