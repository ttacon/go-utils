package sqlutil

import (
	"database/sql"
	"fmt"
)

type SqlUtil struct {
	db *sql.DB
}

func New(dbConn *sql.DB) SqlUtil {
	return SqlUtil{
		db: dbConn,
	}
}

func (s SqlUtil) DescribeTable(tableName string) ([]ColumnInfo, error) {
	rows, err := s.db.Query(fmt.Sprintf("describe %s", tableName))
	if err != nil {
		return nil, err
	}

	var columns []ColumnInfo
	for rows.Next() {
		var field, typ, null, key, defaul, extra sql.NullString
		err = rows.Scan(&field, &typ, &null, &key, &defaul, &extra)
		if err != nil {
			return nil, err
		}
		columns = append(columns, ColumnInfo{
			Field:   denilify(field),
			Type:    denilify(typ),
			Null:    denilify(null),
			Key:     denilify(key),
			Default: denilify(defaul),
			Extra:   denilify(extra),
		})
	}
	return columns, nil
}

func denilify(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}

type ColumnInfo struct {
	Field,
	Type,
	Null,
	Key,
	Default,
	Extra string
}

func (s SqlUtil) ShowTables(database string) ([]string, error) {
	var query = "show tables"
	if len(database) > 0 {
		query = fmt.Sprintf("show tables in %s", database)
	}

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var (
		tables    []string
		tableName string
	)

	for rows.Next() {
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}

func (s SqlUtil) ShowDatabases() ([]string, error) {
	rows, err := s.db.Query("show databases")
	if err != nil {
		return nil, err
	}

	var (
		databases    []string
		databaseName string
	)

	for rows.Next() {
		err = rows.Scan(&databaseName)
		if err != nil {
			return nil, err
		}
		databases = append(databases, databaseName)
	}
	return databases, nil
}
