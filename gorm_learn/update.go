package main

import "fmt"

func UpdateTest() {
	// update 更新选择的field
	// updates 更新所有field, 传入Map或Strcut,里面的零值不参与更新
	// save  类似updates，但零值也会参与更新, 根据主键进行更新，没有对应的主键则会插入
	// DB.Create(&[]Player{
	// 	{UID: "a-001", Name: "mn", Age: 22},
	// })

	// UPDATE `players` SET `name`='mm' WHERE name like '%m%'
	// DB.Model(&Player{}).Where("name like ?", "%m%").Update("name", "mm")

	// DB.Save(&Player{Name: "mm", Age: 22})

	// var players []Player

	// dbRes := DB.Where("age > ?", "20").Find(&players)

	// for k, _ := range players {
	// 	players[k].Age += 1
	// }
	// dbRes.Save(&players)

	var player Player
	// Struct
	// DB.First(&player).Updates(Player{Name: "xy", Age: 25})

	// Map
	DB.First(&player).Updates(map[string]interface{}{"name": "xy", "age": 0})

	DB.Find(&player)
	fmt.Println(player)
}
