package main

import (
	"fmt"

	"example.com/gorm_learn/v2/customize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type UserTest struct {
	ID   int
	Name string
}

type UT struct {
	gorm.Model
	Name string
}

type Player struct {
	UID  string `gorm:"primaryKey"`
	Name string
	Age  int
}

type Role struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

var DB *gorm.DB

func connectDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:root1234@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{ // 表名命名策略
			// TablePrefix:   "t_",  // 表名前缀，`User` 对应的表名是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 表名将变成 `t_user`, 不启用为 `t_users`
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 是否禁用外键(option)
		// 现在主张逻辑外键（代码里体现外键关系），不建议使用数据库外键，会影响性能
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})
	if err != nil {
		fmt.Println("连接数据库失败, error=", err)
		return nil
	}
	return db
}

func main() {
	db := connectDB()
	DB = db
	// 自动创建表

	// CreateTest(db)
	// CreateTest2(db)
	// CreateTest3(db)
	// CreateTest4(db)
	// QueryTest(db)
	// QueryTest2(db)
	// QueryTest3(db)
	// QueryTest4(db)
	// QueryTest5(db)
	// QueryTest6(db)
	// QueryTest7(db)
	// QueryTest8(db)
	// QueryTest9(db)
	// QueryTest10(db)
	// UpdateTest()
	// DeleteTest()
	// RawSqlTest()
	// one2one.Test(db)
	// one2one.Test2(db)
	// one2many.Test(db)
	// one2many.Test2(db)
	// many2many.Test(db)
	// many2many.Test2(db)
	// polymorphic.Test(db)
	// tags.Test2(db)
	// tags.Test3(db)
	// transaction.Test(db)
	customize.Test(db)
}
