//проекты запускать в папке C:\Users\{c-student/student}\source\folderProject

package main

import "fmt"

func main() {
	//  переменные
	var arenda float64 = 95000
	var procent float64 = 10

	var povishenie float64 = arenda * (procent / 100)

	// Плюсуем к старой цене
	var novaya_cena float64 = arenda + povishenie

	fmt.Println("Новая цена аренды:", novaya_cena)
}
