package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ttacon/go-utils/db/sqlutil"
)

var (
	dsn   = flag.String("dsn", "example:example@/example", "the dsn to use to connect to the desired db with")
	table = flag.String("t", "", "the desired table to describe")
)

func main() {
	flag.Parse()

	if len(*dsn) == 0 {
		fmt.Println("must provide dsn info to connect to db")
		return
	}

	if len(*table) == 0 {
		fmt.Println("must provide a table to inspect")
		return
	}

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		fmt.Println("failed to connect to db, err: ", err)
		return
	}

	sqlUtil := sqlutil.New(db)
	cols, err := sqlUtil.DescribeTable(*table)
	if err != nil {
		fmt.Printf("failed to describe table %q, err: %v\n", *table, err)
		return
	}

	fmt.Printf("%13s%13s%13s%13s%13s%13s\n",
		"Field", "Type", "Null", "Key", "Default", "Extra")
	for _, col := range cols {
		fmt.Printf("%13s%13s%13s%13s%13s%13s\n",
			col.Field, col.Type, col.Null, col.Key, col.Default, col.Extra)
	}
}
