package main

import "fmt"

func main() {
	fmt.Println("Приветствуб тебя! Я система LOS!")
	fmt.Println("by Art'S")
	//int:
	var id uint8 = 23 //uint значит , что id > 0! int - любое
	ider := 3.4       //Сокращеный вариант верхней части! (Чаще всего применяется)
	fmt.Println(id)
	//float:
	var id_float float64 = 3.422
	fmt.Println(id_float)
	fmt.Println(ider)
	//string
	var name string = "Arrt"
	namer := "Arrt"
	fmt.Println(len(namer)) //Функция len() позваляет найти строки
	fmt.Println(name)
	fmt.Println(ider)

	// Запись(Чтение)

	var nameser string
	fmt.Println("Введите 1 - 4")
	fmt.Scan(&nameser)
	fmt.Println("Вы выбрали " + nameser + "!")
	fmt.Println(fmt.Sprint(id)) // fmt.Sprint() - из строки в число

	//типы
	var ei int8 = 2
	var sf int64 = int64(2)
	fmt.Println(ei)
	fmt.Println(sf)
	//if / else
	linker := 0
	fmt.Println("Введите 1 - 4")
	fmt.Scan(&linker)

	if linker > 2 {
		fmt.Println("Вы получили новый опыт!")
		linker = 0
	} else if linker <= 2 {
		fmt.Println("Вы НЕ получили новый опыт!")
		linker = 0
	} else if linker = 0 {
		fmt.Println("Вы !НЕ! получили новый опыт!")	
	}

}
