package main

import (
	"fmt"
)

func main() {
	var dayShortName string
	fmt.Print("Введите будний день в короткой форме (пн, вт, ср, чт, пт): ")

	var dayNumer int

	switch fmt.Scan(&dayShortName); dayShortName {
	case "пн":
		dayNumer = 0
	case "вт":
		dayNumer = 1
	case "ср":
		dayNumer = 2
	case "чт":
		dayNumer = 3
	case "пт":
		dayNumer = 4
	default:
		fmt.Println("Неверный день недели")
		return
	}

	switch {
	case dayNumer <= 0:
		fmt.Println("Понедельник")
		fallthrough
	case dayNumer <= 1:
		fmt.Println("Вторник")
		fallthrough
	case dayNumer <= 2:
		fmt.Println("Среда")
		fallthrough
	case dayNumer <= 3:
		fmt.Println("Четверг")
		fallthrough
	case dayNumer <= 4:
		fmt.Println("Тяпница")
	}
}
