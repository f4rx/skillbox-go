package main

import (
	"fmt"
)

func main() {
	var username, password string

	fmt.Print("Добро пожаловать в программу имитации входа")
	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&username)
	fmt.Print("Введите пароль пользователя: ")
	fmt.Scan(&password)

	fmt.Print("-----\n")
	fmt.Printf("%s, вы успешно зашли \n", username)

}
