/*
查询：
接受返回值的形式：
	Map:
		Map[string]interface{}
		[]Map[string]interface{}
	结构体
查询主键排序后的第一条
	db.First(&user)
查询第一条(不排序)
	db.Take
查询主键排序后的最后一条
	db.Last(&user)
条件
	主键检索
	String条件：
		where
	Struct & Map条件：
		可以传入一个有内容的结构体作为条件，
		也可以用一个key为数据库表字段的Map作为条件
	内联条件：
		省略Where， 直接在得到的结果的Find First等后面写条件
Select制定字段：
Join
*/

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func QueryTest(db *gorm.DB) {

	var dbRes *gorm.DB
	var user UserTest
	var res map[string]interface{}

	/* dbRes = db.First(&user) // 根据主键或者第一个字段排序
	// SELECT * FROM `user_tests` ORDER BY `user_tests`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(user)

	dbRes = db.Last(&user)
	// SELECT * FROM `user_tests` ORDER BY `user_tests`.`id` DESC LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(user)

	dbRes = db.Take(&user) // 不排序
	// SELECT * FROM `user_tests` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(user)


	res = map[string]interface{}{}
	dbRes = db.Model(&UserTest{}).First(&res)
	// SELECT * FROM `user_tests` ORDER BY `user_tests`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(res) */

	res = map[string]interface{}{}
	// dbRes = db.Table("user_tests").First(&res)
	// can't get primary key or first column value from map[string]interface{}
	// SELECT * FROM `user_tests` ORDER BY `user_tests`.`` LIMIT 1
	dbRes = db.Table("user_tests").Take(&res)
	// it's Ok, because it's not sorting
	// SELECT * FROM `user_tests` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(res)

	dbRes = db.Model(&UserTest{}).First(&user)
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(user)
}

func QueryTest2(db *gorm.DB) {

	var dbRes *gorm.DB
	var ut UT
	// var res map[string]interface{}

	dbRes = db.Where("id = ?", 2).First(&ut)
	// SELECT * FROM `uts` WHERE id = 2 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(ut)

	dbRes = db.First(&ut, 3)
	// SELECT * FROM `uts` WHERE `uts`.`id` = 3 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(ut)

	dbRes = db.Where(UT{Name: "b2"}).First(&ut)
	// SELECT * FROM `uts` WHERE `uts`.`name` = b2 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(ut)

	db.Where(map[string]interface{}{
		"name": "b2",
	}).First(&ut)
	// SELECT * FROM `uts` WHERE `uts`.`name` = b2 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(ut)

	dbRes = db.Where("name = ?", "b2").
		Or("name = ?", "c3").First(&ut)
	// SELECT * FROM `uts` WHERE name = b2 OR name = c3 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(ut)

	dbRes = db.Where("id > ?", "1").
		Not("name = ?", "b2").First(&ut)
	// SELECT * FROM `uts` WHERE id > 1 AND NOT(name = b2) ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	// fmt.Println(ut)
}

/* Retrieving objects with primary key */
func QueryTest3(db *gorm.DB) {
	var dbRes *gorm.DB
	// var users []UserTest

	// dbRes = db.Find(&users, []int{1, 2, 3})
	// // SELECT * FROM `uts` WHERE `uts`.`id` IN (1,2,3)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(users)

	// var player Player
	// 主键不是int类型,需要指定主键字段
	// dbRes = db.First(&player, "uid = ?", "y-001")
	// // SELECT * FROM `players` WHERE uid = y-001 ORDER BY `players`.`uid` LIMIT 1
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(player)

	// 会找出匹配cu已有属性(必须是主键)的记录，并进行补全
	cu := UserTest{ID: 2}
	// dbRes = db.First(&cu)
	// SELECT * FROM `user_tests` WHERE `user_tests`.`id` = 2 ORDER BY `user_tests`.`id` LIMIT 1
	// cu := UserTest{Name: "c3"}  // 不行，不是主键的值
	dbRes = db.First(&cu)
	// SELECT * FROM `user_tests` WHERE `user_tests`.`name` = a1 ORDER BY `user_tests`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(cu)

	// dbRes = db.Model(UserTest{Name: "b2"}).First(&cu)
	// dbRes = db.Model(UserTest{ID: 2}).First(&cu)
	// // SELECT * FROM `user_tests` WHERE `user_tests`.`name` = b2 ORDER BY `user_tests`.`id` LIMIT 1
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(cu)

	// var cp Player
	// // 下面这个不成功
	// dbRes = db.Model(Player{UID: "x-002"}).First(&cp)
	// // SELECT * FROM `players` ORDER BY `players`.`uid` LIMIT 1

	// cp = Player{UID: "x-002"}
	// dbRes = db.First(&cp)
	// // SELECT * FROM `players` WHERE uid = x-002 ORDER BY `players`.`uid` LIMIT 1
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(cp)

	var ut UT
	ut = UT{Model: gorm.Model{ID: 2}}
	dbRes = db.First(&ut)
	// SELECT * FROM `uts` WHERE `uts`.`deleted_at` IS NULL AND `uts`.`id` = 2 ORDER BY `uts`.`id` LIMIT 1
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(ut)
}

/* retrieving all objects */
func QueryTest4(db *gorm.DB) {
	var dbRes *gorm.DB
	var players []Player

	// SELECT * FROM `players`
	dbRes = db.Find(&players)
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println("RowsAffected:", dbRes.RowsAffected)
	fmt.Println(players)

}

/* string condition */
func QueryTest5(db *gorm.DB) {
	var dbRes *gorm.DB
	// var player Player
	var players []Player

	// SELECT * FROM `players` WHERE name = 'a1' ORDER BY `players`.`uid` LIMIT 1
	// dbRes = db.Where("name = ?", "a1").First(&player)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(player)

	// // SELECT * FROM `players` WHERE name <> 'a1'
	// dbRes = db.Where("name <> ?", "a1").Find(&players)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(players)

	// // SELECT * FROM `players` WHERE name IN ('a1','c3')
	// dbRes = db.Where("name IN ?", []string{"a1", "c3"}).Find(&players)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(players)

	// // SELECT * FROM `players` WHERE name LIKE 'd%'
	// dbRes = db.Where("name LIKE ?", "d%").Find(&players)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(players)

	// SELECT * FROM `players` WHERE name = 'a1' AND age >= 10
	dbRes = db.Where("name = ? AND age >= ?", "a1", 10).Find(&players)
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(players)

	var uts []UT
	// SELECT * FROM `uts` WHERE updated_at > '2025-01-22 02:05:52.519' AND `uts`.`deleted_at` IS NULL
	dbRes = db.Where("updated_at > ?", time.Now().Add(-time.Hour*1000)).Find(&uts)
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(uts)

	// SELECT * FROM `players` WHERE age BETWEEN 10 AND 20
	dbRes = db.Where("age BETWEEN ? AND ?", 10, 20).Find(&players)
	if dbRes.Error != nil {
		log.Println(dbRes.Error)
	}
	fmt.Println(players)

	// // SELECT * FROM `players` WHERE uid = 'x-001' AND `players`.`uid` = 'x-000' ORDER BY `players`.`uid` LIMIT 1
	// Will give record not found Error
	// player = Player{UID: "x-000"}
	// dbRes = db.Where("uid = ?", "x-001").First(&player)
	// if dbRes.Error != nil {
	// 	log.Println(dbRes.Error)
	// }
	// fmt.Println(player)
}

/* Struct & Map Condition */
func QueryTest6(db *gorm.DB) {
	// var dbRes *gorm.DB
	var player Player

	// // SELECT * FROM `players` WHERE `players`.`name` = 'b2' AND `players`.`age` = 19 ORDER BY `players`.`uid` LIMIT 1
	// db.Where(&Player{Name: "b2", Age: 19}).First(&player)
	// fmt.Println(player)

	// player = Player{}
	// // SELECT * FROM `players` WHERE `players`.`name` = 'b2' AND `players`.`age` = 19 ORDER BY `players`.`uid` LIMIT 1
	// db.Where(map[string]interface{}{"name": "b2", "age": 19}).Find(&player)
	// fmt.Println(player)

	// var uts []UT
	// // SELECT * FROM `uts` WHERE `uts`.`id` IN (1,2,3) AND `uts`.`deleted_at` IS NULL
	// db.Where([]int64{1, 2, 3}).Find(&uts)
	// fmt.Println(uts)

	// player = Player{}
	// // struct里的0值不会被查询,包括其他默认值
	// // SELECT * FROM `players` WHERE `players`.`name` = 'c3'
	// db.Where(&Player{Name: "c3", Age: 0}).Find(&player)
	// fmt.Println(player)

	// player = Player{}
	// // map里的0值会被查询
	// // SELECT * FROM `players` WHERE `Age` = 0 AND `Name` = 'c3'
	// db.Where(map[string]interface{}{"Name": "c3", "Age": 0}).Find(&player)
	// fmt.Println(player)

	player = Player{}
	// // SELECT * FROM `players` WHERE `players`.`name` = 'c3' AND `players`.`age` = 0
	// db.Where(&Player{Name: "c3"}, "name", "Age").Find(&player)
	// SELECT * FROM `players` WHERE `players`.`age` = 0
	db.Where(&Player{Name: "c3"}, "Age").Find(&player)
	fmt.Println(player)
}

/* inline Condition */
func QueryTest7(db *gorm.DB) {
	var player Player
	// // directly use First, don't need to use Where
	// db.First(&player, "name = ?", "a1")
	// fmt.Println(player)

	// player = Player{}
	// db.First(&player, "name <> ? AND age > ?", "a1", 10)
	// fmt.Println(player)

	// Not
	player = Player{}
	db.Not("name = ?", "a1").First(&player)
	fmt.Println(player)

	// Or
	var players []Player
	db.Where("name = ?", "a1").Or("name = ?", "c3").Find(&players)
	fmt.Println(players)
}

/* Select Specify Fields */
func QueryTest8(db *gorm.DB) {
	var player Player

	db.Select("name", "age").First(&player)
	fmt.Println(player)

	player = Player{}
	db.Select([]string{"name", "age"}).First(&player)
	fmt.Println(player)

	player = Player{}
	// SELECT COALESCE(age, 19) FROM `players
	res, err := db.Table("players").Select("COALESCE(age, ?)", 19).Rows()
	if err != nil {
		log.Println(err)
	}
	for res.Next() {
		db.ScanRows(res, &player)
		fmt.Println(player)
	}

}

/* Order */
func QueryTest9(db *gorm.DB) {
	var players []Player

	// SELECT * FROM `players` ORDER BY age desc, name
	db.Order("age desc, name").Find(&players)
	fmt.Println(players)

	// SELECT * FROM `players` ORDER BY age desc, name
	db.Order("age desc").Order("name").Find(&players)
	fmt.Println(players)

	var player Player
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "FIELD(name,?)",
			Vars:               []interface{}{[]string{"b2", "c3", "a1"}},
			WithoutParentheses: true,
		},
	}).Find(&player)
	fmt.Println(player)
}

/* Limit & Offset */
func QueryTest10(db *gorm.DB) {
	var players []Player
	var players2 []Player

	// // SELECT * FROM `players` LIMIT 2
	// db.Limit(2).Find(&players)
	// fmt.Println(players)

	// // Cancel limit condition with -1
	// players = []Player{}
	// // SELECT * FROM `players` LIMIT 2
	// // SELECT * FROM `players`
	// db.Limit(2).Find(&players).Limit(-1).Find(&players2)
	// fmt.Println(players)
	// fmt.Println(players2)

	players = []Player{}
	db.Offset(2).Find(&players)
	fmt.Println(players)

	// players = []Player{}
	// // SELECT * FROM `players` LIMIT 2 OFFSET 1
	// db.Limit(2).Offset(1).Find(&players)
	// fmt.Println(players)

	players = []Player{}
	players2 = []Player{}
	// Cancel offset condition with -1
	db.Offset(2).Find(&players).Offset(-1).Find(&players2)
	fmt.Println(players)
	fmt.Println(players2)
}

/* Group By & Having */
func QueryTest11(db *gorm.DB) {

}
