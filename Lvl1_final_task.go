package main

import (
	"fmt"
)

func main() {
	a := "-"
	b := "-"
	c := "-"
	d := "-"
	e := "-"
	cnt := 0
	for {
		var s string
		fmt.Scan(&s)
		if s == "очередь" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", a, b, c, d, e)
			continue
		}
		if s == "конец" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", a, b, c, d, e)
			return
		}
		if s == "количество" {
			fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", 5-cnt, cnt)
			continue
		}
		var x int
		fmt.Scan(&x)
		if x > 5 || x < 1 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", x)
			continue
		}
		if cnt == 5 {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", x)
			continue
		}
		switch x {
		case 1:
			if a == "-" {
				a = s
				cnt++
			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", x)
			}

		case 2:
			if b == "-" {
				b = s
				cnt++

			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", x)
			}
		case 3:
			if c == "-" {
				c = s
				cnt++

			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", x)
			}
		case 4:
			if d == "-" {
				d = s
				cnt++

			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", x)
			}
		case 5:
			if e == "-" {
				e = s
				cnt++

			} else {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", x)
			}
		default:
			continue
		}
	}
}
