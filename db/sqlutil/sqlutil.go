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

func (s SqlUtil) Describe(tableName string) ([]ColumnInfo, error) {
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
