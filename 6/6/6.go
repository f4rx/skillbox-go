/*
Задание 6*. Движение лифта
В доме 24 этажа. Лифт должен ходить по кругу, пока не доставит всех пассажиров на первый этаж. Три пассажира ждут на четвёртом, седьмом и десятом этажах. При движении вверх лифт не должен останавливаться, при движении вниз — собирать всех, но не более двух человек в лифте. При этом лифт каждый раз доезжает до самого верхнего этажа и только после этого начинает движение вниз. Напишите программу, которая доставит всех пассажиров на первый этаж.
*/

package main

import (
	"fmt"
)

const floorsCount = 24
const directionUp = 1
const directionDown = 2
const directionFloorDown = 3
const directionStop = 4

func main() {

	var floors [floorsCount]int
	floors[3] = 1
	floors[6] = 1
	floors[9] = 3

	direction := directionUp
	passengersInLift := 0

	fmt.Printf("%v\n", getHighFloor(floors))

	for {
		fmt.Printf("%v\n", floors)

		if direction == directionUp {
			goToFloor := getHighFloor(floors)
			fmt.Printf("Лифт поднимается на этаж %d за пассажирами\n", goToFloor+1)
			enterToLift(&passengersInLift, &floors[goToFloor])
			if passengersInLift == 2 || !isNeedRunLift(floors) {
				direction = directionDown
			} else {
				direction = directionFloorDown
			}
		} else if direction == directionDown {
			fmt.Printf("Лифт едет прямиком на 1й этаж и высаживает %d пассажиров\n", passengersInLift)
			passengersInLift = 0
			if isNeedRunLift(floors) {
				direction = directionUp
			} else {
				direction = directionStop
			}
		} else if direction == directionFloorDown {
			goToFloor := getHighFloor(floors)
			fmt.Printf("Лифт опускается на этаж %d за пассажирами\n", goToFloor+1)
			enterToLift(&passengersInLift, &floors[goToFloor])
			if passengersInLift == 2 || !isNeedRunLift(floors) {
				direction = directionDown
			} else {
				direction = directionFloorDown
			}
		} else if direction == directionStop {
			break
		}
	}

	fmt.Println("Всех развезли, всех доставили")

}

func isNeedRunLift(l [floorsCount]int) bool {
	isNeed := false
	for _, passenger := range l {
		if passenger > 0 {
			isNeed = true
			break
		}
	}
	return isNeed
}

func getHighFloor(l [floorsCount]int) int {
	for i := len(l) - 1; i >= 0; i-- {
		if l[i] > 0 {
			return i
		}
	}
	return 0
}

func enterToLift(inLiftCount, inFloor *int) {
	if *inLiftCount == 2 {
		return
	} else {
		canEnter := 2 - *inLiftCount
		if canEnter < *inFloor {
			*inLiftCount += canEnter
			*inFloor -= canEnter
		} else {
			*inLiftCount += *inFloor
			*inFloor = 0
		}
	}
}
