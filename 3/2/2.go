package main

import (
	"fmt"
)

func main() {
	var orderPrice, deliveryPrice, discoutAmount int
	fmt.Println("Введите стоимость товара:")
	fmt.Scan(&orderPrice)
	fmt.Println("Введите стоимость доставки:")
	fmt.Scan(&deliveryPrice)
	fmt.Println("Введите размер скидки:")
	fmt.Scan(&deliveryPrice)


	totalPrice := orderPrice + deliveryPrice - discoutAmount

	fmt.Print("---------\n")
	fmt.Printf("Итого: %d\n", totalPrice)

}
