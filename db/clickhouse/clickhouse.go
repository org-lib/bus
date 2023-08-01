package clickhouse

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mailru/go-clickhouse"
	"strings"
)

// Clickhouse 初始化连接信息对象
type Info struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
}

// 获取一个数据库连接
func Open(cfg Info) (*sql.DB, error) {
	if strings.Trim(cfg.Database, "") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	path := fmt.Sprintf("http://%v:%v@%v:%v/%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	return sql.Open("clickhouse", path)
}
