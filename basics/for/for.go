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

}
