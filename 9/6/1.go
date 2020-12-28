package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v\n", math.MaxUint8)
	fmt.Printf("%v\n", math.MaxUint16)
	fmt.Printf("%v\n", math.MaxUint32)

	OverFlowCountsUint8 := 0
	OverFlowCountsUint16 := 0
	var varUint8 uint8 = 0
	var varUint16 uint16 = 0

	for i := uint32(0); i < math.MaxUint32; i++ {
		varUint8++
		varUint16++
		if varUint8 == 0 {
			OverFlowCountsUint8++
		}
		if varUint16 == 0 {
			OverFlowCountsUint16++
		}
	}
	fmt.Printf("OverFlowCountsUint8: %d\n", OverFlowCountsUint8)
	fmt.Printf("OverFlowCountsUint16: %d\n", OverFlowCountsUint16)
}
