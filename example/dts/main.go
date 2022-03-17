package main

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/org-lib/bus/config"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}
type C_Sharp_Arg struct {
	Addr     string `toml:"addr"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	// Databases []string `toml:"databases"`
	Database  string   `toml:"database"`
	Tables    []string `toml:"tables"`
	TAddr     string   `toml:"taddr"`
	TUser     string   `toml:"tuser"`
	TPassword string   `toml:"tpassword"`
	TDatabase string   `toml:"tdatabase"`
}

//监听数据记录
func (h *MyEventHandler) OnRow(ev *canal.RowsEvent) error {
	//record := fmt.Sprintf("%s %v %v %v %s\n",e.Action,e.Rows,e.Header,e.Table,e.String())

	//库名，表名，行为，数据记录
	record := fmt.Sprintf("%v %v %s %v\n", ev.Table.Schema, ev.Table.Name, ev.Action, ev.Rows)
	fmt.Println(record)

	//此处是参考 https://github.com/gitstliu/MysqlToAll 里面的获取字段和值的方法
	for columnIndex, currColumn := range ev.Table.Columns {
		//字段名，字段的索引顺序，字段对应的值
		row := fmt.Sprintf("%v, %v, %v 。\n", currColumn.Name, columnIndex, ev.Rows[len(ev.Rows)-1][columnIndex])
		fmt.Println("row info:", row)
	}
	return nil
}

//创建、更改、重命名或删除表时触发，通常会需要清除与表相关的数据，如缓存。It will be called before OnDDL.
func (h *MyEventHandler) OnTableChanged(schema string, table string) error {
	//库，表
	record := fmt.Sprintf("%s %s \n", schema, table)
	fmt.Println(record)
	return nil
}

//监听binlog日志的变化文件与记录的位置
func (h *MyEventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	//源码：当force为true，立即同步位置
	record := fmt.Sprintf("%v %v \n", pos.Name, pos.Pos)
	fmt.Println("OnPosSynced", record)
	return nil
}

//当产生新的binlog日志后触发(在达到内存的使用限制后（默认为 1GB），会开启另一个文件，每个新文件的名称后都会有一个增量。)
func (h *MyEventHandler) OnRotate(r *replication.RotateEvent) error {
	//record := fmt.Sprintf("On Rotate: %v \n",&mysql.Position{Name: string(r.NextLogName), Pos: uint32(r.Position)})
	//binlog的记录位置，新binlog的文件名
	record := fmt.Sprintf("On Rotate %v %v \n", r.Position, r.NextLogName)
	fmt.Println(record)
	return nil

}

// create alter drop truncate(删除当前表再新建一个一模一样的表结构)
func (h *MyEventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
	//binlog日志的变化文件与记录的位置
	record := fmt.Sprintf("%v %v\n", nextPos.Name, nextPos.Pos)
	query_event := fmt.Sprintf("%v\n %v\n %v\n %v\n %v\n",
		queryEvent.ExecutionTime,         //猜是执行时间，但测试显示0
		string(queryEvent.Schema),        //库名
		string(queryEvent.Query),         //变更的sql语句
		string(queryEvent.StatusVars[:]), //测试显示乱码
		queryEvent.SlaveProxyID) //从库代理ID？
	fmt.Println("OnDDL:", record, query_event)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func DefaultExport(csa *C_Sharp_Arg) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = csa.Addr
	cfg.User = csa.User
	cfg.Password = csa.Password

	cfg.Dump.TableDB = csa.Database
	// cfg.Dump.Tables = []string{in_table1, in_table2}
	cfg.Dump.Tables = csa.Tables

	c, err := canal.NewCanal(cfg)
	if err != nil {
		fmt.Println("error", err)
	}

	c.SetEventHandler(&MyEventHandler{})
	//mysql-bin.000004, 1027
	startPos := mysql.Position{Name: "mysql-bin.000001", Pos: 0}

	fmt.Println("Go run")
	//从头开始监听
	//c.Run()
	//根据位置监听
	c.RunFrom(startPos)
}
func main() {
	csa_in := C_Sharp_Arg{
		Addr:     fmt.Sprintf("%v:%v", config.Config.V.GetString("mysql.host"), config.Config.V.GetString("mysql.port")),
		User:     config.Config.V.GetString("mysql.username"),
		Password: config.Config.V.GetString("mysql.password"),
		Database: config.Config.V.GetString("mysql.database"),
		Tables:   []string{config.Config.V.GetString("mysql.tables")},
	}
	DefaultExport(&csa_in)

	// go build -buildmode=c-shared -o mysqlslave.dll mysql57slave/slave.go
	// [DllImport("mysqlslave.dll", EntryPoint="DefaultExport")]
	// static extern void DefaultExport();

	// [DllImport("exportgo.dll", EntryPoint="Sum")]
	// static extern int Sum(int a, int b);

	// static void Main()
	// {
	// 	Console.WriteLine("Hello World!");
	// 	PrintBye();

	// 	Console.WriteLine(Sum(33, 22));
	// }

}
