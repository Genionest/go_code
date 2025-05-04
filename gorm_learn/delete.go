package main

import "fmt"

func DeleteTest() {
	var player Player

	// DB.Where("name = ?", "xy").Delete(&player)
	// player = Player{UID: "a-001"}
	// 对于有deleted_at字段的表，会执行软删除，会将deleted_at字段置为当前时间
	// DB.Delete(&player) // 根据主键删除，其他字段不起作用

	// DB.Unscoped().Delete(&player) // 硬删除,

	fmt.Println(player)
}
