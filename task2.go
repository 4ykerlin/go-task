package main

import "fmt"

func main() {
	//  переменные
	var monitor float64 = 21830
	var mishka float64 = 890
	var noutbuk float64 = 55480
	var klava float64 = 1560
    // посчитать общую сумму, которую нужно выделить 
	var summa float64 = (monitor * 3) + (noutbuk * 6) + (mishka * 11) + (klava * 5)

	fmt.Println("Нужно выделить:", summa)
}
