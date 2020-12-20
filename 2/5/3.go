package main

import (
	"fmt"
)

func main() {
	shiftDuration := 480
	customerOrderTime := 2
	cashierPerOrderTime := 4

	fmt.Println("Эта программа рассчитает, сколько клиентов успеет обслужить кассир за смену.")
	fmt.Printf("Введите длительность смены в минутах: %d\n", shiftDuration)
	fmt.Printf("Сколько минут клиент делает заказ? %d\n", customerOrderTime)
	fmt.Printf("Сколько минут кассир собирает заказ? %d\n", cashierPerOrderTime)
	fmt.Println("-----Считаем-----")

	customersPerShift := shiftDuration / (customerOrderTime + cashierPerOrderTime)

	fmt.Printf("За смену длиной %d минут кассир успеет обслужить %d клиентов.\n", shiftDuration, customersPerShift)

}
