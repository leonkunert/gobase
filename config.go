package main

type Configuration struct {
    // External Port
    HttpPort   string
    // MySQL Host and Port
    DBHost     string
    // MySQL User
    DBUser     string
    // MySQL Password
    DBPass     string
    // MySQL Table
    DBTable    string
    // Additional Parameters
    // For further info https://github.com/Go-SQL-Driver/MySQL/#parameters
    DBParams   string
}

/*
Initialize Configuration
Defines some configuration options in a global variable
 */
var CFG = Configuration {
    // Http Port
    HttpPort: "80",
    // Database Host and Port
    DBHost: "10.0.0.1:1234",
    // Database User
    DBUser: "user",
    // Password for Database User
    DBPass: "secret",
    // Database Table
    DBTable: "table",
    // Additional Parameters
    DBParams: "?charset=utf8&collation=utf8_general_ci&parseTime=true",
}