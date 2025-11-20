package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// TODO: прочитать аргумент как float64
	// TODO: посчитать по формуле и вывести с 1 знаком после запятой
	C, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Print("Ошибка\n")
	}
	fmt.Printf("%.1f °F\n", C*9/5+32)

}
