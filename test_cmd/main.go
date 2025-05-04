package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var s, sep string
	// // os.Args
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}
