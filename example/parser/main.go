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
	sql := `
SELECT
  t1.id AS id,
  t1.version AS version,
  t1.create_time AS createTime,
  t1.channel_code AS channelCode,
  t1.customer_code AS customerCode,
  t1.owner_code AS ownerCode,
  t1.bin_code AS binCode,
  t1.sku_code AS skuCode,
  t1.location_code AS locationCode,
  t1.share_ratio AS shareRatio,
  t1.is_availability AS isAvailability,
  t1.share_thresholds AS shareThresholds,
  t1.is_specify_sku AS isSpecifySku,
  t1.operation_status AS operationStatus
FROM
  inv_sku_balance_63 t1 FORCE INDEX (idx1_inv_sku_balance)
  JOIN inv_sku_balance_63 t2 FORCE INDEX (idx1_inv_sku_balance) ON t1.customer_code = t2.customer_code
  AND t1.bin_code = t2.bin_code
  AND t1.sku_code = t2.sku_code
  AND t1.location_code = t2.location_code
  AND t1.channel_code = t2.channel_code
  AND t2.customer_code = 'I.TKG'
  AND t2.owner_code = 'ITMall'
  AND t2.bin_code = 'PT090'
  AND t2.sku_code = '5587461-YEX-00S'
  AND t2.location_code = '999'
  AND t2.saas_tenant_code = 'baozun'
  AND t1.saas_tenant_code = 'baozun'
WHERE
  t1.saas_tenant_code = 'baozun'`
	//logs.Info(GetTableNames(sql, "dml"))
	//fmt.Println(GetTableNames(sql, "ddl"))
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
