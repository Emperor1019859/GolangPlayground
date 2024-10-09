package main

import (
	"fmt"
)

func main() {
	for i := range 11 {
		even := i%2 == 0

		if even {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is old")
		}
	}
}
