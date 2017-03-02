package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep
	}
	end := time.Since(start)
	fmt.Printf("Execution Time loop: %v\n", end)

	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	end = time.Since(start)
	fmt.Printf("Execution Time Join: %v\n", end)

}
