package main

import "fmt"

func main() {
	//  переменные
        var f float64
        fmt.Print("Введите температуру по Фаренгейту: ")
        fmt.Scan(&f)
        c := (5.0 / 9.0) * (f - 32)
        fmt.Printf("Температура по Цельсию: %.2f°C\n", c)
    }
