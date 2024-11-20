//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func hello(i int) {
//	println("hello goroutine : " + fmt.Sprint(i))
//}
//
//func main() {
//	for i := 0; i < 5; i++ {
//		go func(j int) {
//			hello(j)
//		}(i)
//	}
//	time.Sleep(time.Second)
//}

package Bytedance

func Judge(a int) bool {
	if a > 50 {
		return true
	}
	return false
}

func foo3() {
	src := make(chan int)
	dest := make(chan int, 3)

	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()
	for i := range dest {
		println(i)
	}
}
