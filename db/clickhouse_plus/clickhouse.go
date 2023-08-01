package clickhouse_plus

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

// Info Clickhouse 初始化连接信息对象
type Info struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`
}

// Open 获取数据库连接
func Open(cfg Info) (*sqlx.DB, error) {
	path := fmt.Sprintf("tcp://%v:%v?debug=true&username=%v&password=%v&database=%v",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
	connect, err := sqlx.Connect("clickhouse", path)
	if err != nil {
		return nil, err
	}
	//defer func() {
	//	_ = connect.Close()
	//}()

	return connect, nil
}
