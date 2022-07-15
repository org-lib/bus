package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

// Info MongoDB 初始化连接信息对象
type Info struct {
	Host          string
	Port          string
	Username      string
	Password      string
	DefaultAuthDB string
	Options       string
}

// Open 获取一个数据库连接
func Open(cnf *Info) (*mongo.Client, error) {
	var url string
	if strings.Trim(cnf.Username, " ") == "" || strings.Trim(cnf.Password, " ") == "" {
		url = fmt.Sprintf("mongodb://%v:%v", cnf.Host, cnf.Port)
	} else {
		if strings.Trim(cnf.DefaultAuthDB, " ") == "" {
			if strings.Trim(cnf.Options, " ") == "" {
				url = fmt.Sprintf("mongodb://%v:%v@%v:%v", cnf.Username, cnf.Password, cnf.Host, cnf.Port)
			} else {
				url = fmt.Sprintf("mongodb://%v:%v@%v:%v/?replicaSet=%v", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.Options)
			}
		} else {
			if strings.Trim(cnf.Options, " ") == "" {
				url = fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.DefaultAuthDB)
			} else {
				url = fmt.Sprintf("mongodb://%v:%v@%v:%v/%v?replicaSet=%v", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.DefaultAuthDB, cnf.Options)
			}

		}
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(url)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
