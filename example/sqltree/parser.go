package main

import (
	"fmt"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/parse"
)

//代码实测没问题，但是依赖1.19
func main() {
	// 定义要解析的SQL语句
	sqlStr := "SELECT inv.`id` AS id, inv.`customer_code` AS customerCode, inv.`owner_code` AS ownerCode, inv.`bin_code` AS binCode, inv.`sku_code` AS skuCode , inv.`inv_status_code` AS invStatusCode, inv.`location_code` AS locationCode, inv.`quota_interval` AS quotaInterval , CASE WHEN SUM(occ.`qty`) != '' THEN inv.`qty` + SUM(occ.`qty`) ELSE inv.`qty` END AS qty FROM inv_customer_inventory_98 inv LEFT JOIN inv_occupy_98 occ FORCE INDEX (idx_comp1) ON occ.`bin_code` = inv.`bin_code` AND occ.`sku_code` = inv.`sku_code` AND occ.`inv_status_code` = inv.`inv_status_code` AND occ.`location_code` = inv.`location_code` AND occ.`customer_code` = inv.`customer_code` AND occ.`owner_code` = inv.`owner_code` AND occ.`quota_interval` = inv.`quota_interval` AND occ.saas_tenant_code = 'YACE' AND inv.saas_tenant_code = 'YACE' WHERE inv.`customer_code` = '压测002' AND inv.`owner_code` = 'IT后端测试店铺02' AND inv.`bin_code` = 'SDTEST01' AND inv.`sku_code` = '[TEST]0BKYHy_080-y03080-9' AND inv.`inv_status_code` = '10' AND inv.`location_code` = '999' AND inv.quota_interval = 'DEFAULT_NULL' AND inv.saas_tenant_code = 'YACE' GROUP BY inv.`customer_code`, inv.`bin_code`, inv.`sku_code`, inv.`inv_status_code`, inv.`location_code`, inv.`owner_code`, inv.`quota_interval`"

	// 创建解析上下文
	ctx := sql.NewContext(nil)

	// 解析SQL语句
	stmt, err := parse.Parse(ctx, sqlStr)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	// 输出语法树和语义信息
	fmt.Println("AST:", stmt.String())
}
