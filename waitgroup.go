package Bytedance

import (
	"fmt"
	"sync"
)

func hello(i int) {
	fmt.Println(i)
}

func foo1() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
