// Mutex

package main

import "sync"

func main() {
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()
}
