package main

import (
	"fmt"
	"net"
)

func worker(ports chan int, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("20.194.168.28:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// fmt.Printf("Addr %s is closed\n", address)
			result <- 0
			continue
		}
		conn.Close()
		result <- p
		// fmt.Printf("Addr %s is open\n", address)
		// fmt.Println(p)
		// wg.Done()
	}
}

func main() {
	ports := make(chan int, 20)
	result := make(chan int)
	var openports []int
	var closeports []int
	// var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
		// go worker(ports, &wg, result)
	}

	go func() {
		for i := 1; i < 50; i++ {
			port := <-result
			if port != 0 {
				openports = append(openports, port)
			} else {
				closeports = append(closeports, port)
			}
		}
	}()

	for i := 1; i < 50; i++ {
		// wg.Add(1)
		ports <- i
	}
	// wg.Wait()
	close(ports)
	close(result)

	for _, port := range closeports {
		fmt.Printf("%d closed\n", port)
	}

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

// mult proccessor
// func main() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	for i := 21; i < 50; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("20.194.168.28:%d", j)
// 			conn, err := net.Dial("tcp", address)
// 			if err != nil {
// 				fmt.Printf("Addr %s is closed\n", address)
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("Addr %s is open\n", address)
// 		}(i)
// 	}
// 	wg.Wait() // counter is 0, then it will not stop
// 	elapsed := time.Since(start) / 1e9
// 	fmt.Printf("Time taken: %d seconds\n", elapsed)
// }

// single processor
// func main() {
// 	for i := 21; i < 50; i++ {
// 		address := fmt.Sprintf("20.194.168.28:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			fmt.Printf("Addr %s is closed\n", address)
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("Addr %s is open\n", address)
// 	}
// }
