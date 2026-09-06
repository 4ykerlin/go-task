package main

import 
"fmt"


func main() {
	var total float64
    fmt.Print("Введите сумму покупки: ")
    fmt.Scan(&total)

    discount := total * 0.20
    finalPrice := total - discount

    fmt.Println("Скидка: ", discount)
    fmt.Println("Итог: ", finalPrice)
    }
