package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 使用容器名作为主机名
	db, err := sql.Open("mysql", "root:mysecret@tcp(mysql-container:3306)/mydb")
	// mysql-container是容器名，不同于docker-compose.yml里的是service名
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 等待数据库启动（最多尝试10次）
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("连接失败，重试中... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("最终连接失败:", err)
	}

	fmt.Println("成功连接到MySQL容器！")
}
