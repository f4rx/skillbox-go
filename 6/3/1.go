package main

import (
	"fmt"
)

const password = "pass"

func main(){

	var input string
	for  {
		fmt.Println("Введите пароль")
		fmt.Scan(&input)
		if input == password {
			fmt.Println("Пароль верный")
			break
		}
		fmt.Println("Пароль неверный, попробуйте еще раз")
	}

}
