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
	tableName := strings.ToLower(ctx.GetText())
	m.tableNames[tableName] = struct{}{}
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

func ParseSQL(sql string) []string {
	input := antlr.NewInputStream(sql)
	lexer := parser.NewMySqlLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMySqlParser(stream)
	ml := Ml{tableNames: make(map[string]struct{})}
	antlr.ParseTreeWalkerDefault.Walk(&ml, p.Root())
	return ml.GetTableNames()
}

// github.com/akito0107/xsqlparser 支持with 语法
func main() {
	sql1 := "SELECT * FROM table1;"
	fmt.Println(ParseSQL(sql1))

	sql2 := "INSERT INTO table2 VALUES (1, 'test');"
	fmt.Println(ParseSQL(sql2))

	sql3 := "UPDATE table3 SET name = 'test' WHERE id = 1;"
	fmt.Println(ParseSQL(sql3))

	sql4 := "DELETE FROM table4 WHERE id = 1;"
	fmt.Println(ParseSQL(sql4))

	sql5 := "TRUNCATE TABLE table5;"
	fmt.Println(ParseSQL(sql5))

	sql6 := "DROP TABLE table6;"
	fmt.Println(ParseSQL(sql6))

	sql7 := "ALTER TABLE table7 ADD COLUMN column1 varchar(255);"
	fmt.Println(ParseSQL(sql7))
}
