package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ttacon/go-utils/db/sqlutil"
)

var (
	dsn    = flag.String("dsn", "example:example@/example", "the dsn to use to connect to the desired db with")
	dbName = flag.String("db", "", "the desired database to describe")
)

func main() {
	flag.Parse()

	if len(*dsn) == 0 {
		fmt.Println("must provide dsn info to connect to db")
		return
	}

	if len(*dbName) == 0 {
		fmt.Println("must provide a database to inspect")
		return
	}

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		fmt.Println("failed to connect to db, err: ", err)
		return
	}

	sqlUtil := sqlutil.New(db)
	tables, err := sqlUtil.ShowTables(*dbName)
	if err != nil {
		fmt.Printf("failed to show tables for db %q, err: %v\n", *dbName, err)
		return
	}

	fmt.Println("Tables:")
	for _, table := range tables {
		fmt.Println(table)
	}
}
