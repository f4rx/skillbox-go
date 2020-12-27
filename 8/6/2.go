package main

import (
	"fmt"
)

func main() {
	var dayShortName string
	fmt.Print("Введите будний день в короткой форме (пн, вт, ср, чт, пт): ")
	fmt.Scan(&dayShortName)

	dayNumberMap := map[string]int{
		"пн": 0,
		"вт": 1,
		"ср": 2,
		"чт": 3,
		"пт": 4,
	}

	if _, ok := dayNumberMap[dayShortName]; !ok {
		fmt.Println("Неверный день недели")
		return
	}

	switch {
	case dayNumberMap[dayShortName] <= 0:
		fmt.Println("Понедельник")
		fallthrough
	case dayNumberMap[dayShortName] <= 1:
		fmt.Println("Вторник")
		fallthrough
	case dayNumberMap[dayShortName] <= 2:
		fmt.Println("Среда")
		fallthrough
	case dayNumberMap[dayShortName] <= 3:
		fmt.Println("Четверг")
		fallthrough
	case dayNumberMap[dayShortName] <= 4:
		fmt.Println("Тяпница")
	}
}
