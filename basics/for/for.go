package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
		time.Sleep(1 * time.Second)
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
