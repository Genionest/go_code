package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

func worker(ports chan int, wg *sync.WaitGroup, results chan int) {
	for p := range ports { // 会一直循环知道channel关闭
		// fmt.Println(p)
		fmt.Println("results<-p")
		results <- p
		// wg.Done()
	}
}

func WithProcessPool() {
	ports := make(chan int, 100)
	results := make(chan int)
	var even []int
	var odd []int
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		// 每个worker都在等待着执行任务，并不是完成一个就算了，
		// 而是一直获取channel中的任务，直到关闭了channel后才退出。
		go worker(ports, &wg, results)
	}

	go func() {
		for i := 0; i < 10; i++ {
			// wg.Add(1)
			fmt.Println("ports<-i", i)
			ports <- i
		}
	}()
	// 注意channel的阻塞机制
	for i := 0; i < 10; i++ {
		fmt.Println("port<-results")
		port := <-results
		if port%2 == 0 {
			even = append(even, port)
		} else {
			odd = append(odd, port)
		}
	}
	// wg.Wait()
	close(ports) // 关闭channel
	close(results)
	sort.Ints(even)
	sort.Ints(odd)

	for _, v := range even {
		fmt.Printf("even %d\n", v)
	}
	for _, v := range odd {
		fmt.Printf("odd %d\n", v)
	}
}

func SingleProcess() {
	for i := 21; i < 120; i++ {
		address := fmt.Sprintf("20.194.168.28:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d关闭了\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了\n", address)
	}
}

func MultiProcess() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i < 120; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了\n", address)
		}(i)
	}
	wg.Wait()
	elasped := time.Since(start) / 1e9
	fmt.Printf("耗时%v秒\n", elasped)
}

func main() {
	// SingleProcess()
	// MultiProcess()
	WithProcessPool()
}
