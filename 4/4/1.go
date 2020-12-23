/*
Задача 1. Космонавты

Для пилотируемой миссии на Марс отобрали троих космонавтов. Беда в том, что они оказались настолько похожи, что выбрать командира экипажа оказалось затруднительно. Тогда решили проверить их IQ. Напишите программу, которая запрашивает IQ трёх космонавтов и назначает командира по наивысшему IQ.
*/

package main

import (
	"fmt"
)

func main() {
	var cosmonautIQ1,  cosmonautIQ2,  cosmonautIQ3 int

	fmt.Print("Введите IQ для первого космонавта: ")
	fmt.Scan(&cosmonautIQ1)
	fmt.Print("Введите IQ для второго космонавта: ")
	fmt.Scan(&cosmonautIQ2)
	fmt.Print("Введите IQ для третьего космонавта: ")
	fmt.Scan(&cosmonautIQ3)

	if cosmonautIQ1 > cosmonautIQ2 {
		if cosmonautIQ1 > cosmonautIQ3 {
			fmt.Println("Командир экипажа первый космонавт")
		} else {
			fmt.Println("Командир экипажа третий космонавт")
		}
	} else if cosmonautIQ2 > cosmonautIQ1 {
		if cosmonautIQ2 > cosmonautIQ3 {
			fmt.Println("Командир экипажа второй космонавт")
		} else {
			fmt.Println("Командир экипажа третий космонавт")
		}
	} else {
		fmt.Println("Командир экипажа третий космонавт")
	}
}
