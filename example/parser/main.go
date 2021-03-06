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
	roleName   map[string]struct{}
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
func GetSelectColumnElement(sql string, sqlType string) []string {
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
	return ml.GetSelectColumnElement()
}
func GetSelectExpressionElement(sql string, sqlType string) []string {
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
	return ml.GetSelectExpressionElement()
}
func GetFromClause(sql string, sqlType string) []string {
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
	return ml.GetFromClause()
}

func main() {
	//FromClauseContext
	sql := "select name from tablea where 1!2"
	//logs.Info(GetTableNames(sql, "dml"))
	fmt.Println(GetTableNames(sql, "dql"))
	//fmt.Println(GetSelectColumnElement(sql, "dql"))
	//fmt.Println(GetSelectExpressionElement(sql, "dql"))
	fmt.Println(GetFromClause(sql, "dql"))
}
func (m *Ml) EnterSelectColumnElement(ctx *parser.SelectColumnElementContext) {
	if m.roleName == nil {
		m.roleName = make(map[string]struct{})
	}
	m.roleName[ctx.GetText()] = struct{}{}
}
func (m *Ml) GetSelectColumnElement() []string {
	arr := make([]string, 0)
	if m.roleName != nil {
		for k := range m.roleName {
			arr = append(arr, k)
		}
	}
	return arr
}
func (m *Ml) EnterFromClause(ctx *parser.FromClauseContext) {
	if m.roleName == nil {
		m.roleName = make(map[string]struct{})
	}
	m.roleName[ctx.GetWhereExpr().GetText()] = struct{}{}
}
func (m *Ml) GetFromClause() []string {
	arr := make([]string, 0)
	if m.roleName != nil {
		for k := range m.roleName {
			arr = append(arr, k)
		}
	}
	return arr
}
func (m *Ml) EnterSelectExpressionElement(ctx *parser.SelectExpressionElementContext) {
	if m.roleName == nil {
		m.roleName = make(map[string]struct{})
	}
	m.roleName[ctx.GetText()] = struct{}{}
}
func (m *Ml) GetSelectExpressionElement() []string {
	arr := make([]string, 0)
	if m.roleName != nil {
		for k := range m.roleName {
			arr = append(arr, k)
		}
	}
	return arr
}
