package main

import (
	"fmt"
)

func main() {
	busStop1 := "Улица смешное название 1"
	busStop2 := "Улица смешное название 2"
	busStop3 := "Улица смешное название 3"
	busStop4 := "Улица смешное название 4"
	fare := 20
	passengersTotal := 0
	passengersCount := 0
	passengersExitedCount := 0
	passengersEnteredCount := 0

	fmt.Println("*Имитация маршртутки*")
	// Step 1
	fmt.Printf("Прибываем на остановку `%s`. В салоне пассажиров: %d\n", busStop1, passengersCount)
	/*
		понятно что на первой остановке не может никто выйти, т.к. 0 в салоне
		fmt.Print("Сколько пассажиров вышло на остановке? ")
		fmt.Scan(&passengersExitedCount)
		passengersCount -= passengersExitedCount
	*/
	fmt.Print("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnteredCount)
	passengersCount += passengersEnteredCount
	passengersTotal += passengersEnteredCount
	fmt.Printf("Отправляемся с остановки `%s`. В салоне пассажиров: %d\n", busStop1, passengersCount)
	fmt.Println("-----------Едем---------")

	// Step 2
	fmt.Printf("Прибываем на остановку `%s`. В салоне пассажиров: %d\n", busStop2, passengersCount)
	fmt.Print("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersExitedCount)
	passengersCount -= passengersExitedCount
	fmt.Print("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnteredCount)
	passengersCount += passengersEnteredCount
	passengersTotal += passengersEnteredCount
	fmt.Printf("Отправляемся с остановки `%s`. В салоне пассажиров: %d\n", busStop2, passengersCount)
	fmt.Println("-----------Едем---------")

	// Step 3
	fmt.Printf("Прибываем на остановку `%s`. В салоне пассажиров: %d\n", busStop3, passengersCount)
	fmt.Print("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersExitedCount)
	passengersCount -= passengersExitedCount
	fmt.Print("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnteredCount)
	passengersCount += passengersEnteredCount
	passengersTotal += passengersEnteredCount
	fmt.Printf("Отправляемся с остановки `%s`. В салоне пассажиров: %d\n", busStop3, passengersCount)
	fmt.Println("-----------Едем---------")

	// Step 4
	fmt.Printf("Прибываем на остановку `%s`. В салоне пассажиров: %d\n", busStop4, passengersCount)
	fmt.Print("Сколько пассажиров вышло на остановке? ")
	fmt.Scan(&passengersExitedCount)
	passengersCount -= passengersExitedCount
	fmt.Print("Сколько пассажиров зашло на остановке? ")
	fmt.Scan(&passengersEnteredCount)
	passengersCount += passengersEnteredCount
	passengersTotal += passengersEnteredCount
	fmt.Printf("Отправляемся с остановки `%s`. В салоне пассажиров: %d\n", busStop4, passengersCount)
	fmt.Println("-----------Едем---------")

	fmt.Println("--Приехали на конечную--")
	fmt.Println()

	// Calculation
	totalIncome := passengersTotal * fare
	salaryExpenditure := float64(totalIncome) * 0.25
	tmpExpenditure20 := float64(totalIncome) * 0.2
	fuelExpenditure, taxExpenditure, carRepairExpenditure := tmpExpenditure20, tmpExpenditure20, tmpExpenditure20
	profit := float64(totalIncome) * 0.15

	fmt.Println("Расчёт прибыли")
	fmt.Printf("Всего заработали: %.2f руб.\n", float64(totalIncome))
	fmt.Printf("Зарплата водителя: %.2f руб.\n", salaryExpenditure)
	fmt.Printf("Расходы на топливо: %.2f руб.\n", fuelExpenditure)
	fmt.Printf("Налоги: %.2f руб.\n", taxExpenditure)
	fmt.Printf("Расходы на ремонт машины: %.2f руб.\n", carRepairExpenditure)
	fmt.Printf("Прибыль: %.2f руб.\n", profit)

}
