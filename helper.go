package dbhelper

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	DefaultPageSize = 100
	MaxPageSize     = 10000
)

func countPage(total, pageSize int) int {
	if total < 1 {
		return 0
	}
	if pageSize < 1 {
		pageSize = 1
	}
	if total%pageSize == 0 {
		return total / pageSize
	}
	return total/pageSize + 1
}

func buildSQL(fields, table, where, orderby string, pageSize, pageIndex int) string {
	sql := ""
	if strings.TrimSpace(where) == "" {
		where = " 1=1 "
	}
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageSize < 1 || pageSize > MaxPageSize {
		pageSize = DefaultPageSize
	}
	if pageIndex == 1 {
		sql = fmt.Sprintf("SELECT TOP %d %s FROM %s WHERE %s ORDER BY %s DESC", pageSize, fields, table, where, orderby)
	} else {
		sql = fmt.Sprintf("SELECT %s FROM %s WHERE %s IN (SELECT TOP %d %s FROM (SELECT TOP %d %s FROM %s WHERE %s ORDER BY %s DESC) t1 ORDER BY %s ASC) ORDER BY %s DESC", fields, table, orderby, pageSize, orderby, pageSize*pageIndex, orderby, table, where, orderby, orderby, orderby)
	}
	return sql
}

// GetPage ...
func GetPage(db *sqlx.DB, data interface{}, fields, table, where, orderby string, pageSize, pageIndex int) (err error) {
	if strings.TrimSpace(where) == "" {
		where = " 1=1 "
	}
	sql := "SELECT COUNT(0) FROM " + table + " WHERE " + where
	count := 0
	err = db.Get(&count, sql)
	if err != nil {
		return err
	}
	pages := countPage(count, pageSize)
	if pageIndex < 1 {
		pageIndex = 1
	}
	if pageIndex > pages {
		pageIndex = pages
	}
	sql = buildSQL(fields, table, where, orderby, pageSize, pageIndex)
	err = db.Select(data, sql)
	return
}
