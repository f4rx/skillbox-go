package main

import (
	"fmt"
)

func main() {
	var username, password string
	var age int

	fmt.Println("Добро пожаловать в программу регистрации")
	fmt.Println("Введите имя пользователя:")
	fmt.Scan(&username)
	fmt.Println("Введите пароль пользователя:")
	fmt.Scan(&password)
	fmt.Println("Введите возраст пользователя:")
	fmt.Scan(&age)

	fmt.Printf(`Поздравляю, %s, теперь вы зарегистрированы!
Ваш пароль: %s
Ваш возраст: %d
`, username, password, age)

}
