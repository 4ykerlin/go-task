package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b int
    fmt.Print("Введите два целых числа (a b): ")
    fmt.Scan(&a, &b)

    if b == 0 {
        fmt.Println("Деление на ноль!")
        return
    }

    result := float64(a) / float64(b)
    rounded := math.Round(result*100) / 100 // округление до 2 знаков

    fmt.Println("Результат деления:", rounded)
    fmt.Println("Без округления:", result)
}