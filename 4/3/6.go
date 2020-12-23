/*
Задача 4. Барберы* (сложная задача)

В небольшом городке проживает 18 000 взрослых человек: 9 000 мужчин и 9 000 женщин. Согласно традиции этого городка, все мужчины обязаны носить бороду. Один предприимчивый житель открыл сеть барбершопов, чтобы помочь мужчинам ухаживать за их бородами. Однако, он никак не может посчитать, сколько всего специалистов-барберов для этого нужно.

Напишите программу, которая будет рассчитывать: сколько мужчин можно постричь во всех барбершопах, исходя из следующих данных:

Каждый из 9000 мужчин посещает барбершоп раз в месяц (30 дней)
Один барбер способен обслужить одного клиента за 1 час
Смена барбера — 8 часов
Программа должна спросить, сколько мужчин проживает в городе и сколько всего барберов уже работает во всех барбершопах, и посчитать сколько барберов нужно, и если их недостаточно — выдать сообщение об этом. Если барберов достаточно — сказать и об этом.

Подсказка: в конструкции условия можно сравнивать не только переменную с числом, но и две переменные! А еще — целые математические выражения. Например: 

if a * b > c {

    // ...

}
*/

package main

import (
	"fmt"
)

func main() {

	var men, hairdresserCount int
	shiftDuration := 8
	hairdresserSpeed := 1
	daysInMonth := 30

	fmt.Print("Сколько мужчин сейчас в городе ? ")
	fmt.Scan(&men)

	fmt.Print("Сколько бродобреев уже сейчас в городе ? ")
	fmt.Scan(&hairdresserCount)
	// Сколько человек могут обработь за месяц
	hairdresserNeedHours := hairdresserCount * hairdresserSpeed * shiftDuration * daysInMonth
	// Обрабатывает за месяц
	hairdresserSpeedPerEmployee := hairdresserSpeed * shiftDuration * daysInMonth

	fmt.Println("DEBUG доступно человеко-часов", hairdresserNeedHours)
	fmt.Println("DEBUG один парихмахер обрабатывает ", hairdresserSpeedPerEmployee)

	hairdresserNeed := (men - hairdresserNeedHours) / hairdresserSpeedPerEmployee

	fmt.Println("DEBUG hairdresserNeed", hairdresserNeed)

	if (men - hairdresserNeedHours) % hairdresserSpeedPerEmployee > 0 {
		hairdresserNeed++
	}

	if men > hairdresserNeedHours {
		fmt.Println("Не хватает бродобреев:", hairdresserNeed)
	} else {
		fmt.Println("Хватает бродобреев")
	}

}
