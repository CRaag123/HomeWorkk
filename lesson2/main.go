package main

import "fmt"

func main() {
	var info string
	info = printInfo(30, 12, 2020)
	fmt.Println(info)

}

func printInfo(day, month, year int) (info string) {

	if day%10 == 0 {
		if month < 10 {
			info = fmt.Sprintf("Сегодня - %d.0%d.%d", day, month, year)
		} else {
			info = fmt.Sprintf("Сегодня - %d.%d.%d", day, month, year)
		}
		return info
	} else {
		info = fmt.Sprintf("%d - не делится на 10", day)
		return info
	}

}
