package dbhelper

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var (
	DefaultPageSize = 100
	MaxPageSize     = 10000
)

// PageCount ...
func PageCount(total, pagesize int) int {
	if total < 1 {
		return 0
	}
	if pagesize < 1 {
		pagesize = 1
	}
	if total%pagesize == 0 {
		return total / pagesize
	}
	return total/pagesize + 1
}

// BuildSQL ...
func BuildSQL(fields, table, where, orderby string, pageSize, pageNum int) string {
	tmpValue := ""
	if where != "" {
		where = " WHERE " + where
	}
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 || pageSize > MaxPageSize {
		pageSize = DefaultPageSize
	}
	if pageNum == 1 {
		tmpValue = fmt.Sprintf("SELECT TOP %d %s FROM %s %s ORDER BY %s DESC", pageSize, fields, table, where, orderby)
	} else {
		tmpValue = "SELECT " + fields + " FROM " + table + " WHERE " + orderby +
			" IN (SELECT TOP " + strconv.Itoa(pageSize) + " " + orderby + " FROM (SELECT TOP " + strconv.Itoa(pageSize*pageNum) + " " + orderby + " FROM " + table + " " + where + " ORDER BY " + orderby + " DESC) t1 ORDER BY " + orderby + " ASC) ORDER BY " + orderby + " DESC"
	}
	return tmpValue
}

// GetPage ...
func GetPage(db *sqlx.DB, fields, table, where, orderby string, pageSize, pageNum int) (data []interface{}, err error) {
	sql := "SELECT COUNT(0) FROM " + table + " WHERE " + where
	count := 0
	err = db.Get(&count, sql)
	if err != nil {
		return nil, err
	}
	pageCount := PageCount(count, pageSize)
	if pageNum < 1 {
		pageNum = 1
	}
	if pageNum > pageCount {
		pageNum = pageCount
	}
	sql = BuildSQL(fields, table, where, orderby, pageSize, pageNum)
	err = db.Select(&data, sql)
	return
}
