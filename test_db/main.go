package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	server   = "localhost"
	port     = 3306
	user     = "root"
	password = "root1234"
	database = "testdb"
)

func main() {
	// connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	server, user, password, port, database)
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connected!")

	d, err := getOne(3)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(d)

	d.name = "Charlie"

	err = d.Update()
	if err != nil {
		log.Fatalln(err.Error())
	}

	newd, err := getOne(3)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(newd)

	d2 := data{
		name:  "Cat",
		email: "cat@example.com",
	}

	err = d2.Insert()
	if err != nil {
		log.Fatalln(err.Error())
	}
	one, err := getOne(d2.ID)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(one)

	// datas, err := getMany(2)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// fmt.Println(datas)
}
