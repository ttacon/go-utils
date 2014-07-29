// Package sqlutil is a package that contains several useful utility
// functions for interacting with sql databases.
//
// When connecting to a database with a specific dsn
// ([user]:[password]@/[database]), information such as what databases
// on the current host are available to the current user, what tables
// are availble for a given database and how a given table was defined
// is not easily available. This collection of functions drastically
// simplifies the retreival of this information, as the following code
// demonstrates.
//
//     // assuming we have a *sql.DB from db, err := sql.Open(...)
//     // we create a SqlUtil
//     sqlUtil := sqlutil.New(db)
//
//     // We can then retrieve the databases on the host with
//     dbs, err := sqlUtil.ShowDatabases()
//
//     // We can show the tables for a given database with
//     tables, err := sqlUtil.ShowTables("myDB")
//
//     // And we can see how a table is defined with
//     columns, err := sqlUtil.DescribeTable("myTable")
//
package sqlutil
