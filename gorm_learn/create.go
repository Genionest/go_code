/*
创建：
单独创建：
	db.Create(&User{Name: "jinzhu", Age: 18})
制定字段：
	db.Select("Name").Create(&User{Name: "jinzhu", Age: 18})
跳过字段：
	db.Omit("Name").Create(&User{Name: "jinzhu", Age: 18})
批量创建：
	db.Create(&[]User{})
*/

package main

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateTest(db *gorm.DB) {
	db.AutoMigrate(&UserTest{})

	// result := db.Create(&UserTest{
	// 	Name: "test",
	// })
	// fmt.Print(result.Error, result.RowsAffected)
	// db.Select("Name").Create(&UserTest{Name: "test"}) // 指定字段
	// db.Omit("Name").Create(&UserTest{Name: "test"})   // 跳过字段
	// db.Create(&[]UserTest{
	// 	{Name: "test1"},
	// 	{Name: "test2"},
	// })
	var res *gorm.DB
	res = db.Create(&[]UserTest{
		{Name: "a1"},
		{Name: "b2"},
		{Name: "c3"},
		{Name: "d4"},
	})
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(res.RowsAffected)
}

func CreateTest2(db *gorm.DB) {
	db.AutoMigrate(&UT{})

	var res *gorm.DB
	res = db.Create(&[]UT{
		{Name: "a1"},
		{Name: "b2"},
		{Name: "c3"},
		{Name: "d4"},
	})
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(res.RowsAffected)
}

func CreateTest3(db *gorm.DB) {
	db.AutoMigrate(&Player{})

	var res *gorm.DB
	res = db.Create(&[]Player{
		{UID: "x-001", Name: "a1", Age: 22},
		{UID: "x-002", Name: "b2", Age: 19},
		{UID: "x-003", Name: "c3", Age: 20},
		{UID: "y-001", Name: "d4", Age: 15},
	})
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(res.RowsAffected)
}

func CreateTest4(db *gorm.DB) {
	db.AutoMigrate(&Role{})

	var res *gorm.DB
	res = db.Create(&[]Role{
		{ID: 1, Name: "b1"},
		{ID: 2, Name: "a2"},
		{ID: 3, Name: "d3"},
		{ID: 4, Name: "c4"},
	})
	if res.Error != nil {
		panic(res.Error)
	}
	fmt.Println(res.RowsAffected)
}
