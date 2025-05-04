package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func fn() {
	// str := "5Yqg5YWlR1ZB5Lqk5rWB576k"
	// decodeBytes, err := base64.StdEncoding.DecodeString(str)
	// fmt.Println(decodeBytes, err)
}

func fn2() {
	var a []interface{}
	for _, v := range a {
		fmt.Println(v)
	}
}

func fn3() {
	// type A struct {
	// 	n map[string]int
	// 	m sync.Mutex
	// }
	// a := A{n: map[string]int{"a": 1}}
	type A struct {
		n int
		m sync.Mutex
	}
	a := A{n: 1}
	go func() {
		a.m.Lock()
		time.Sleep(300 * time.Millisecond)
		defer a.m.Unlock()
	}()
	go func() {
		a.m.Lock()
		defer a.m.Unlock()
		for i := 0; i < 10; i++ {
			// fmt.Println(a.n["a"])
			fmt.Println(a.n)
		}
	}()
	time.Sleep(1000 * time.Millisecond)
}

func fn4() {
	fmt.Println(strconv.ParseUint("200", 10, 0))
}

func fn5() {
	type A struct {
		n int
	}
	fmt.Println(reflect.TypeOf(A{}))
	fmt.Println(reflect.TypeOf(A{}).Kind())
}

func fn6() {
	lcnt, rcnt := 0, 0
	for i := 0; i < 20; i++ {
		target := rand.Float64() * 100
		left, right := 1, 100
		mid := 0
		// for left <= right {
		// 	mid = (left + right) / 2
		// 	if float64(mid)-target < 0 {
		// 		left = mid + 1
		// 	} else {
		// 		right = mid - 1
		// 	}
		// }
		for left < right {
			mid = (left + right) / 2
			if float64(mid)-target < 0 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// fmt.Println(left, right, mid, int(target))
		if left == mid {
			fmt.Printf("left=mid=%d,\ttarget=%.2f\n", left, target)
			lcnt++
		}
		if right == mid {
			fmt.Printf("right=mid=%d,\ttarget=%.2f\n", right, target)
			rcnt++
		}
		if right != mid && left != mid {
			fmt.Println("error", left, right, mid)
		}

	}
	fmt.Println(lcnt, rcnt)
}

type A struct {
	data int
}

func (a *A) fn() {
	fmt.Println(a.data)
}

func fn7() {
	a := A{data: 1}
	a.fn()
}

func main() {
	fn7()
}
