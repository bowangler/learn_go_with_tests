package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    //
    db, _ := sql.Open("mysql", "root:Password001@(172.18.100.26:3306)/business")
    defer db.Close()
    err := db.Ping()
    if err != nil {
        fmt.Println("connect false")
        fmt.Println(err)
        return
    }

    rows, _ := db.Query("select * from accountvalue where policyNO = '0100034225'")
    var policyNo, branch, accountName, accountSum, accountValue string
    for rows.Next() {
        rows.Scan(&policyNo, &branch, &accountName, &accountSum, &accountValue)
        fmt.Println(policyNo, "--", branch, "--", accountName, "--", accountSum, "--", accountValue)
    }
}
