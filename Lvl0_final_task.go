package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var f string
	fmt.Scanln(&f)
	date, err := time.Parse("02.01.2006", f)
	var name1, name2, name3 string
	fmt.Scanln(&name1)
	fmt.Scanln(&name2)
	fmt.Scanln(&name3)
	var price1, price2, price3 float64
	fmt.Scanln(&price1)
	fmt.Scanln(&price2)
	fmt.Scanln(&price3)

	if err != nil {
		fmt.Printf("Ошибка парсинга: %v\n", err)
		return
	}

	var new_date = date.Add(15 * 24 * time.Hour)
	var total_costs = price1 + price2 + price3
	var total_costs_rub = math.Trunc(total_costs)
	var total_costs_kop = math.Round((total_costs - total_costs_rub) * 100)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы. \nДата подписания договора: %v. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %.0f руб. %.0f коп.\n\nС уважением,\nГл. бух. Иванов А.Е.\n", name2, name1, name3, new_date.Format("02.01.2006"), total_costs_rub, total_costs_kop)
}
