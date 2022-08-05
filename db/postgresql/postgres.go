package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	// 原始支持
	_ "github.com/lib/pq"
	//gorm
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

// Info 初始化连接信息
type Info struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	SslMode  string
	TimeZone string
}

// Open 获取一个数据库连接
func Open(cnf *Info) (*gorm.DB, error) {
	if strings.Trim(cnf.Database, " ") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	if strings.Trim(cnf.Host, " ") == "" {
		cnf.Host = "localhost"
	}
	if strings.Trim(cnf.Host, " ") == "" {
		cnf.Port = "5432"
	}
	if strings.Trim(cnf.TimeZone, " ") == "" {
		cnf.TimeZone = "Asia/Shanghai"
	}
	if strings.Trim(cnf.SslMode, " ") == "" || strings.Trim(strings.ToLower(cnf.SslMode), " ") != "enable" {
		cnf.SslMode = "disable"
	}
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=%v TimeZone=%v", cnf.Host,
		cnf.Port,
		cnf.Database,
		cnf.Username,
		cnf.Password,
		cnf.SslMode,
		cnf.TimeZone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// OpenPlus 获取一个数据库连接
func OpenPlus(cnf *Info) (*sql.DB, error) {
	if strings.Trim(cnf.Database, " ") == "" {
		return nil, errors.New("*** 请至指定一个数据库名称")
	}
	if strings.Trim(cnf.Host, " ") == "" {
		cnf.Host = "localhost"
	}
	if strings.Trim(cnf.Host, " ") == "" {
		cnf.Port = "5432"
	}
	if strings.Trim(cnf.TimeZone, " ") == "" {
		cnf.TimeZone = "Asia/Shanghai"
	}
	if strings.Trim(cnf.SslMode, " ") == "" || strings.Trim(strings.ToLower(cnf.SslMode), " ") != "enable" {
		cnf.SslMode = "disable"
	}
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=%v TimeZone=%v", cnf.Host,
		cnf.Port,
		cnf.Database,
		cnf.Username,
		cnf.Password,
		cnf.SslMode,
		cnf.TimeZone)
	return sql.Open("postgres", dsn)
}
