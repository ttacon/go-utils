package sqlutil

import (
	"database/sql"
	"fmt"
)

// SqlUtil is useful for extra information about a database, it has
// wrapper functions to show what databases exist on the host
// specified by dsn (or at least those that the user specified by
// the dsn can see), for showing what tables exist on a database
// and for describing a database table.
type SqlUtil struct {
	db *sql.DB
}

// ColumnInfo describes the six fields of information
// for a column in a SQL table.
type ColumnInfo struct {
	Field,
	Type,
	Null,
	Key,
	Default,
	Extra string
}

// New returns a SqlUtil which uses the given db connection to
// communicate with the desired database host.
func New(dbConn *sql.DB) SqlUtil {
	return SqlUtil{
		db: dbConn,
	}
}

// DescribeTable returns the columns in the given table.
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

// ShowTables returns a slice of the names of all the tables
// in the given database. The given database can be the empty
// string, in which case the tables for the currently selected
// database are returned.
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

// ShowDatabases shows all the databases that the current user
// can see on the current database host.
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

func denilify(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
