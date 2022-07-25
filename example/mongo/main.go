package main

import (
	"context"
	"fmt"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/mongodb"
	"github.com/org-lib/bus/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
	"net/url"
	"reflect"
)

func main() {

	//定义 cfg 对象
	var cfg *mongodb.Info
	cfg = &mongodb.Info{
		Host:          config.Config.V.GetString("mongo.host"),
		Port:          config.Config.V.GetString("mongo.port"),
		Username:      config.Config.V.GetString("mongo.username"),
		Password:      url.QueryEscape(config.Config.V.GetString("mongo.password")),
		DefaultAuthDB: config.Config.V.GetString("mongo.defaultAuthDB"),
		Options:       config.Config.V.GetString("mongo.replicaSet"),
	}
	var cmd Command
	//获取数据库实例连接
	db, err := mongodb.Open(cfg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("获取数据库实例连接，失败：%v", err), zap.String("mongodb", config.Config.V.GetString("mongodb.port")))
		panic(err)
	}
	defer db.Disconnect(context.TODO())
	_, err = db.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(databases)
	var end bson.Raw
	err = db.Database("db_pim2_dev").RunCommand(context.TODO(), bson.M{"listIndexes": "product_rule"}).Decode(&end)
	logger.Log.Info(fmt.Sprintf("数据翎%v", end))
	cur := end.Lookup("cursor")
	_ = cur.Unmarshal(&cmd.Cursor)

	//日志打印
	logger.Log.Info("数据一", zap.String("mongodb", fmt.Sprintf("%v", cmd.Cursor.FirstBatch[0].Key[0].Key)))

	// objectID 操作方式
	//方式一
	var oid1 []ObjectIdStr
	tmp_cursor1, _ := db.Database("db_pim2_dev").Collection("kk_test").Find(context.TODO(), bson.D{})
	_ = tmp_cursor1.All(context.TODO(), &oid1)
	logger.Log.Info(fmt.Sprintf("方式一%v", oid1))

	//方式二
	var oid []ObjectId
	tmp_cursor, _ := db.Database("db_pim2_dev").Collection("kk_test").Find(context.TODO(), bson.D{})
	_ = tmp_cursor.All(context.TODO(), &oid)
	logger.Log.Info(fmt.Sprintf("方式二%v", oid))

	//方式三
	var tmpIdStr []primitive.ObjectID
	for _, str := range oid1 {
		tx, _ := primitive.ObjectIDFromHex(str.ID)
		tmpIdStr = append(tmpIdStr, tx)
	}
	filter := bson.M{"_id": bson.M{"$in": tmpIdStr}}
	logger.Log.Info(fmt.Sprintf("filter333:=%v", filter))
	inres, err := db.Database("db_pim2_dev").Collection("kk_test").Find(context.TODO(), filter)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("err:=%v", err.Error()))
	}
	var xw []ObjectId
	_ = inres.All(context.TODO(), &xw)
	for _, id := range xw {
		logger.Log.Info(fmt.Sprintf("数据三：%v", id))
	}
	//方式四
	//x1, _ := primitive.ObjectIDFromHex("621f0c12f3a959014cbe9fc0")
	//x2, _ := primitive.ObjectIDFromHex("621f0c16f3a959014cbe9fc1")
	//x3, _ := primitive.ObjectIDFromHex("621f0c1af3a959014cbe9fc2")
	//filter2 := bson.M{"_id": bson.M{"$in": []primitive.ObjectID{x1, x2, x3}}}
	//inres2, err := db.Database("db_pim2_dev").Collection("kk_test").Find(context.TODO(), filter2)
	//if err != nil {
	//	logger.Log.Info(fmt.Sprintf("err:=%v", err.Error()))
	//}
	//var xw2 []bson.M
	//_ = inres2.All(context.TODO(), &xw2)
	////for _, id := range xw2 {
	//logger.Log.Info(fmt.Sprintf("数据四：%v", xw2))
	//}
	//方式五
	tmpIdStr2 := ArrayLengthTes(oid1)
	if len(tmpIdStr2) == 0 {
		return
	}
	filter3 := bson.M{"_id": bson.M{"$in": tmpIdStr2}}
	inres3, err := db.Database("db_pim2_dev").Collection("kk_test").Find(context.TODO(), filter3, options.Find().SetProjection(bson.D{{"noticeEffect", 0}, {"noticeInvalid", 0}}))
	if err != nil {
		logger.Log.Info(fmt.Sprintf("err:=%v", err.Error()))
	}
	var xw5 []bson.D
	var docs []interface{}
	_ = inres3.All(context.TODO(), &xw5)
	logger.Log.Info(fmt.Sprintf("xw5:=%v", xw5))
	for _, d := range xw5 {
		docs = append(docs, d)
		for _, e := range d {

			logger.Log.Info(fmt.Sprintf("column_key:=%v", e.Key))
			logger.Log.Info(fmt.Sprintf("column_value:=%v", e.Value))
		}
		logger.Log.Info(fmt.Sprintf("v:=%v", d))
	}
	xy, err := db.Database("db_pim2_dev").Collection("kk_test_back_test1").InsertMany(context.TODO(), docs)
	if err != nil {
		logger.Log.Info(fmt.Sprintf("err:=%v", err.Error()))
	}
	logger.Log.Info(fmt.Sprintf("res:=%v", xy.InsertedIDs))
}
func ArrayLengthTes(v interface{}) []primitive.ObjectID {
	var tmpIdStr2 []primitive.ObjectID
	switch v.(type) {
	case []ObjectIdStr:
		value := v.([]ObjectIdStr)
		for _, str := range value {
			tx, _ := primitive.ObjectIDFromHex(str.ID)
			tmpIdStr2 = append(tmpIdStr2, tx)
		}
	default:
		logger.Log.Error(fmt.Sprintf("类型不匹配:%v", reflect.TypeOf(v)))
	}
	return tmpIdStr2
}

type Command struct {
	Cursor Cursor `json:"cursor"`
	Ok     int64  `json:"ok"`
}
type Cursor struct {
	FirstBatch []Indexes `json:"firstBatch"`
	Id         int64     `json:"id"`
	Ns         string    `json:"ns"`
}
type Indexes struct {
	Key                bson.D      `json:"key"`
	Name               string      `json:"name"`
	Ns                 string      `json:"ns"`
	V                  int64       `json:"v"`
	Unique             bool        `json:"unique"`
	ExpireAfterSeconds interface{} `json:"ExpireAfterSeconds"`
}
type ObjectId struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}
type KKTest struct {
	ID              primitive.ObjectID `json:"_id,omitempty"`
	BID             string             `json:"bid,omitempty"`
	Name            string             `json:"name,omitempty"`
	EffectStartTime int                `json:"effectStartTime,omitempty"`
	EffectEndTime   string             `json:"effectEndTime,omitempty"`
	Type            string             `json:"type,omitempty"`
	OriginId        string             `json:"originId,omitempty"`
	OriginCode      string             `json:"originCode,omitempty"`
	Version         int                `json:"version,omitempty"`
	CreateTime      int                `json:"createTime,omitempty"`
	CreateUser      string             `json:"createUser,omitempty"`
	Catalog         string             `json:"catalog,omitempty"`
	SaasTenantCode  string             `json:"saasTenantCode,omitempty"`
}
type ObjectIdStr struct {
	ID string `bson:"_id,omitempty"`
}
