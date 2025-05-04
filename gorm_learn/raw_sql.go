package main

import "fmt"

func RawSqlTest() {
	var player Player
	DB.Raw("SELECT * FROM players WHERE name = ?", "b2").Scan(&player)
	fmt.Println(player)
}
