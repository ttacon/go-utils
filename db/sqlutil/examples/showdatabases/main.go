package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ttacon/go-utils/db/sqlutil"
)

var (
	dsn = flag.String("dsn", "example:example@/example", "the dsn to use to connect to the desired db with")
)

func main() {
	flag.Parse()

	if len(*dsn) == 0 {
		fmt.Println("must provide dsn info to connect to db")
		return
	}

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		fmt.Println("failed to connect to db, err: ", err)
		return
	}

	sqlUtil := sqlutil.New(db)
	databases, err := sqlUtil.ShowDatabases()
	if err != nil {
		fmt.Printf("failed to show databases for host, err: %v\n", err)
		return
	}

	fmt.Println("Databases:")
	for _, database := range databases {
		fmt.Println(database)
	}
}
