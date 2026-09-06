package main

import (
"fmt"
"math"
)

func main() {
	var r float64
    fmt.Print("Введите радиус клумбы: ")
    fmt.Scan(&r)

    length := 2 * math.Pi * r
    area := math.Pi * r * r

    fmt.Println("Длина окружности:", length)
    fmt.Println("Площадь круга:", area)
    }
