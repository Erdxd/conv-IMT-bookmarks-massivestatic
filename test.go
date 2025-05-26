package main

import (
	"fmt"
	"sort"
)

type massive = map[string]float64
type mmap = map[string]string

func main() {

	for {
		fmt.Println("Выберите вариант")
		fmt.Println("1-конвертор валют")
		fmt.Println("2-создание закладок")
		fmt.Println("3-найти медиану/среднее арифметическое/одноступенчатый калькулятор чисел")
		fmt.Println("4-выход")
		user := ""
		fmt.Scan(&user)

		if user == "1" {
			conv()
		} else if user == "4" {
			break

		} else if user == "2" {
			boockmarks()
		} else if user == "3" {
			calc2()
		}
	}

}
func conv() {
	for {

		fmt.Println("Конвертор валют(доллар, евро)")
		currencyyouhave := ""
		quantity := 0.0
		currencyyouwant := ""
		valutedollar := map[string]float64{"euro": 0.89, "ruble": 80.50}
		valuteruble := map[string]float64{"euro": 0.011, "dollar": 0.012}
		valuteeuro := map[string]float64{"dollar": 1.12, "ruble": 90.23}
		fmt.Println("Хотите начать?(yes/not)")
		yesornot := ""
		fmt.Scan(&yesornot)
		if yesornot == "yes" {

			fmt.Println("Введите валюту которая у вас есть")
			fmt.Scan(&currencyyouhave)

			if currencyyouhave == "dollar" {
				fmt.Println("current rate:")
				for key, value := range valutedollar {
					fmt.Println(key, value)
				}

				fmt.Println("how much", currencyyouhave, "do you have")
				fmt.Scan(&quantity)
				fmt.Println("Введите какую валюту хотите")
				fmt.Scan(&currencyyouwant)
				if currencyyouwant == "ruble" {
					res := quantity * valutedollar["ruble"]
					fmt.Println("you have", res, currencyyouwant)

				} else if currencyyouwant == "euro" {

					res := quantity * valutedollar["euro"]
					fmt.Println("you have", res, currencyyouwant)
				}
				fmt.Println("how much", currencyyouhave, "do you have")
				fmt.Scan(&quantity)
				fmt.Println("Введите какую валюту хотите")
				fmt.Scan(&currencyyouwant)

			} else if currencyyouhave == "euro" {
				fmt.Println("current rate:")
				for key, value := range valuteeuro {

					fmt.Println(key, value)

				}
				fmt.Println("how much", currencyyouhave, "do you have")
				fmt.Scan(&quantity)
				fmt.Println("Введите какую валюту хотите")
				fmt.Scan(&currencyyouwant)

				if currencyyouwant == "dollar" {
					res := quantity * valuteeuro["dollar"]
					fmt.Println("you have", res, currencyyouwant)
				} else if currencyyouwant == "ruble" {
					res := quantity * valuteeuro["ruble"]
					fmt.Println("you have", res, currencyyouwant)
				}

			} else if currencyyouhave == "ruble" {
				fmt.Println("current rate:")
				for key, value := range valuteruble {

					fmt.Println(key, value)

				}
				fmt.Println("how much", currencyyouhave, "do you have")
				fmt.Scan(&quantity)
				fmt.Println("Введите какую валюту хотите")
				fmt.Scan(&currencyyouwant)
				if currencyyouwant == "euro" {
					res := quantity * valuteruble["euro"]
					fmt.Println("you have", res, currencyyouwant)

				} else if currencyyouwant == "dollar" {
					res := quantity * valuteruble["dollar"]
					fmt.Println("you have", res, currencyyouwant)
				}
			}
		} else {
			break

		}
	}

}

func boockmarks() {

	fmt.Println("Хотите начать? yes/not")
	user1 := ""
	fmt.Scan(&user1)

	if user1 == "yes" {

		m := map[string]string{}
	menu:
		for {

			user := ""
			fmt.Println("Что хотите сделать?")
			fmt.Println("1-Вывести все существующие закладки")
			fmt.Println("2-Добавить закладку")
			fmt.Println("3-удалить закладку")
			fmt.Println("4-Выйти")

			fmt.Scan(&user)
			if user == "1" {
				Printm(m)

			} else if user == "2" {

				m = addbookmark(m)

			} else if user == "3" {
				c := ""
				fmt.Println("Название закладки которую хотите удалить")
				fmt.Scan(&c)
				delete(m, c)

			} else if user == "4" {
				break menu
			}

		}

	}

}

func addbookmark(m mmap) mmap {
	var a string
	var b string

	fmt.Println("Название закладки")
	fmt.Scan(&a)
	fmt.Println("Адрес закладки")
	fmt.Scan(&b)

	for {
		if len(b) > 7 {

			if b[:8] == "https://" || b[:7] == "http://" {
				m[a] = b
				return m

			} else if b == "4" || a == "4" {
				break

			} else {
				fmt.Println("Введите URL (начинается с http:// или с https://). Попробуйте снова ")
				break

			}
		} else {
			fmt.Println("Введите URL (начинается с http:// или с https://). Попробуйте снова ")
			break
		}
	}

	return m

}
func Printm(m mmap) {
	if len(m) == 0 {
		fmt.Println("У вас еще нету закладок")
	} else {
		for key, value := range m {
			fmt.Println(key, "=", value)

		}
	}
}

func calc2() {

	user1 := ""
	for {
		fmt.Println("Готовы начать? (yes/not)")
		fmt.Scan(&user1)

		if user1 == "yes" {

			fmt.Println("что хотите найти? (avg or med or sum )")
			a := ""

			d := 0.0
			fmt.Scan(&a)
			transaction := []float64{}
			for {
				fmt.Println("что хотите найти?")
				fmt.Println("Введите числа(0 для выхода)")

				fmt.Scan(&d)
				fmt.Println("Введите числа(0 для выхода)")
				if d == 0 {
					break
				} else {
					transaction = append(transaction, d)
				}

			}
			fmt.Println(transaction)

			if a == "avg" {

				avg := 0.0
				for _, value := range transaction {
					avg += value

				}
				avg1 := 0.0
				for index := range transaction {
					avg1 += float64(index)

				}
				fmt.Println("среднее значение =", avg/(avg1))

			} else if a == "sum" {
				sum := 0.0
				for _, value := range transaction {
					sum += value

				}
				fmt.Println("сумма =", sum)

			} else if a == "med" {
				sort.Float64s(transaction)
				fmt.Println(transaction)

				if len(transaction)%2 == 0 { // Кол-во чисел четное
					midle2 := 0
					midle3 := 0
					midle2 = len(transaction) / 2
					midle3 = midle2 - 1

					fmt.Println("медиана чисел =", (transaction[midle2]+transaction[midle3])/2)

				} else {
					midle1 := 0
					midle1 = len(transaction) / 2
					fmt.Println("медиана чисел =", transaction[midle1])

				}

			} else {

				error := "Вы ввели что-то не то, попробуйте снова"
				fmt.Println(error)

			}
			//time.Sleep(time.Minute)
		} else if user1 == "not" {
			break
		}
	}
	fmt.Println("dawd")
}
