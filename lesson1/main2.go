package main

import "fmt"

func main2() {
	var result int
	result = sum(3)
	fmt.Println(result)
}

func sum(num int) (result int) {
	for i := 0; i <= num; i++ {
		result += i
	}
	return
}
