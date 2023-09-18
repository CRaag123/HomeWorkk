package main

import "fmt"

func main3() {
	set(4, 10)
}

func set(a int, b int) {
	for i := a; i < b; i += 2 {
		fmt.Println(i)
	}
}
