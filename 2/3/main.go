package main

import (
	"fmt"
)

func main() {

	var lapNumber int = 4
	var leaderDriverName string = "Шумахер"
	var leaderDriverSpeed int = 358
	var carEngine int = 254
	var carWheels int = 93
	var carSteeringWheel int = 49
	var wind int = -21
	var rain int = -17

	fmt.Print("==================\n")
	fmt.Print("Супер гонки. Круг ", lapNumber, "\n")
	fmt.Print("==================\n")
	fmt.Print(leaderDriverName, " (", leaderDriverSpeed, ")\n")
	fmt.Print("==================\n")
	fmt.Print("Водитель: ", leaderDriverName, "\n")
	fmt.Print("Скорость: ", leaderDriverSpeed, "\n")
	fmt.Print("-------------------\n")
	fmt.Print("Оснащение\n")
	fmt.Print("Двигатель: +", carEngine, "\n")
	fmt.Print("Колеса: +", carWheels, "\n")
	fmt.Print("Руль: +", carSteeringWheel, "\n")
	fmt.Print("-------------------\n")
	fmt.Print("Действия плохой погоды\n")
	fmt.Print("Ветер: ", wind, "\n")
	fmt.Print("Дождь: ", rain, "\n")

}
