package main

import "gorm.io/gorm"

func FindTest(db *gorm.DB) {
	var user UserTest
	// 这里Model里的结构体起的作用，和map类似，就是有什么字段就查哪些字段，没有的字段不查
	db.Model(&UserTest{}).First(&user, 1)   // 根据整型主键查找
	db.First(&user, 1)                      // 根据整型主键查找
	db.First(&user, "name = ?", "test")     // 查找 code 字段值为 "test" 的记录
	db.First(&user, "name LIKE ?", "test%") // 查找 name 字段 LIKE "test%" 的记录
	db.First(&user, "name LIKE ?", "test%") // 查找 name 字段 LIKE "
}
