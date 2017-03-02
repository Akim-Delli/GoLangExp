package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	n := 45
	fib45 := fibN(n) // <- slow
	fmt.Printf("Fibonnacci(%d) : %d\n", n, fib45)
}

func fibN(n int) int {
	if n < 2 {
		return n
	}
	return fibN(n-2) + fibN(n-1)
}

func spinner(delay time.Duration) {
	//for {
	//	for _, r := range `-\|/` {
	//		fmt.Printf("\r%c", r)
	//		time.Sleep(delay)
	//	}
	//}
	sep := "="
	for {

		fmt.Printf("\r%s>", sep)
		sep += "="
		time.Sleep(delay)
	}
}
