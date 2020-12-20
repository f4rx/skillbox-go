package main

import (
	"fmt"
)

func main() {
	orderPrice := 6400
	deliveryPrice := 350
	discoutAmount := 700
	totalPrice := orderPrice + deliveryPrice - discoutAmount

	fmt.Printf("Стоимость товара: %d\n", orderPrice)
	fmt.Printf("Стоимость доставки: %d\n", deliveryPrice)
	fmt.Printf("Размер скидки: %d\n", discoutAmount)
	fmt.Print("---------\n")
	fmt.Printf("Итого: %d\n", totalPrice)

}
