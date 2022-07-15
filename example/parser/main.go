package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/org-lib/bus/parser"
	"strings"
)

type Ml struct {
	*parser.BaseMySqlParserListener
	tableNames map[string]struct{}
}

func (m *Ml) EnterTableName(ctx *parser.TableNameContext) {
	if m.tableNames == nil {
		m.tableNames = make(map[string]struct{})
	}
	m.tableNames[ctx.GetText()] = struct{}{}
}
func (m *Ml) GetTableNames() []string {
	arr := make([]string, 0)
	if m.tableNames != nil {
		for k := range m.tableNames {
			arr = append(arr, k)
		}
	}
	return arr
}
func GetTableNames(sql string, sqlType string) []string {
	tokenStream := antlr.NewCommonTokenStream(parser.NewMySqlLexer(antlr.NewInputStream(strings.ToUpper(sql))), antlr.TokenDefaultChannel)
	sqlParser := parser.NewMySqlParser(tokenStream)
	ml := Ml{}
	switch sqlType {
	case "dml":
		antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DmlStatement())
	case "ddl":
		antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DdlStatement())
	case "dql":
		antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.SelectStatement())
	}
	return ml.GetTableNames()
}

func main() {
	sql := "alter table `t_bi_user_shop_middle` add UNIQUE KEY `uniq_t_user_shop_middle_002` (`shop_id`,`user_id`,`role_type`,`is_delete`,`app_code`,`add_from_role_id`) USING BTREE;"
	//logs.Info(GetTableNames(sql, "dml"))
	fmt.Println(GetTableNames(sql, "ddl"))
}
