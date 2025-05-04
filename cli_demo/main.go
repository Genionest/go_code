package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf("Your name is %s", text)
}
