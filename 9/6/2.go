package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int16

	fmt.Print("Введите число 1: ")
	fmt.Scan(&a)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&b)

	result := int32(a) * int32(b)
	fmt.Printf("%d * %d = %d\n", a, b, result)

	/*
		MinInt32        -2147483648
		MinInt16        -32768
		MinInt8         -128
		MaxUint8        255
		MaxUint16       65535
		MaxUint32       4294967295
	*/

	switch {
	case result > math.MaxUint16:
		fmt.Printf("%T\n", uint32(result))
	case result > math.MaxUint8:
		fmt.Printf("%T\n", uint16(result))
	case result > 0:
		fmt.Printf("%T\n", uint8(result))
	case result > math.MinInt8:
		fmt.Printf("%T\n", int8(result))
	case result > math.MinInt16:
		fmt.Printf("%T\n", int16(result))
	case result > math.MinInt32:
		fmt.Printf("%T\n", int32(result))
	default:
		fmt.Printf("Что-то пошло не так...")
		return
	}

}
