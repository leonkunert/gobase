package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var (
    DBC *sql.DB
    err error
)

/*
Initialize MySQL Connection
Starts the connection to the Database in a global variable
 */
func InitDBConnection() {
    DBC, err = sql.Open("mysql", CFG.DBUser+":"+CFG.DBPass+"@tcp("+CFG.DBHost+")/"+CFG.DBTable+CFG.DBParams)
    if err != nil {
        log("InitDBConnection", err.Error())
    }
}