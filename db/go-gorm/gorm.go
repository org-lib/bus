package go_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//gorm.io/gorm 是gorm的官方

type Info struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Charset  string
}

// Open 获取一个数据库连接
func Open(cnf Info) (*gorm.DB, error) {
	// 可以在api包里设置成init函数
	if cnf.Charset == "" {
		cnf.Charset = "utf8mb4"
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%v&parseTime=True&loc=Local",
		cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.Database, cnf.Charset)
	return gorm.Open(mysql.Open(conn), &gorm.Config{})
}
