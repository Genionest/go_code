package main

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("")
	log.SetOutput(os.Stderr)                     // 输出位置
	log.SetFlags(log.LstdFlags | log.Lshortfile) // 输出格式
}

func main() {
	log.Println("a")
	log.Fatalln("b")
	log.Panicln("c")
}
