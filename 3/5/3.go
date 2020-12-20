package main

import (
  "fmt"
)

func main() {
	fmt.Println("Перестановка значений переменных")
	v1()
	v2()
	v3()
}

func v1(){
	fmt.Println("Вариант 1")
	a := 42
	b := 153

	tmpVar := a
	a = b
	b = tmpVar

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func v2(){
	fmt.Println("Вариант 2")
	a := 42
	b := 153

	a, b = b, a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func v3(){
	// Мне кажется, что плохой варинт, т.к. тут есть калькуляция, может работать только с определёнными типами данных, со значениями с плавающими точками могут наверно потеряться данные
	fmt.Println("Вариант 3")
	a := 42
	b := 153

	a += b
	b = a - b
	a -= b

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
