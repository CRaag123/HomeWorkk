package main

import "fmt"

func main() {
	sad(4)
}

func sad(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * num)
	}

}
