package main

import "fmt"

func main() {
	//  переменные
	gb := 5000
	one_file := 256
	
    // посчитать общую сумму, которую нужно выделить 
	razmestit := gb / one_file // файлов поместиться
        svobodnoe_mesto := gb % one_file // сколько останется
	fmt.Println("Сколько можно разместить:", razmestit)
	fmt.Println("Сколько останется места:", svobodnoe_mesto, "гб")
}
