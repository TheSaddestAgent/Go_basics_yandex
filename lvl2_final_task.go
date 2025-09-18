package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var ErrIncorrectDayInput = errors.New("исправь свой ответ, а лучше ложись поспать")

func currentDayOfTheWeek() string {
	var now time.Time = TimeNow()
	weekDay := now.Weekday()
	switch weekDay.String() {
	case "Monday":
		return "Понедельник"
	case "Tuesday":
		return "Вторник"
	case "Wednesday":
		return "Среда"
	case "Thursday":
		return "Четверг"
	case "Friday":
		return "Пятница"
	case "Saturday":
		return "Суббота"
	case "Sunday":
		return "Воскресенье"
	default:
		return "NOT DEFINED"
	}
}
func dayOrNight() string {
	var now time.Time = TimeNow()
	hours := now.Hour()
	if 10 <= hours && hours <= 22 {
		return "День"
	}
	return "Ночь"
}
func nextFriday() int {
	weekDay := currentDayOfTheWeek()
	switch weekDay {
	case "Понедельник":
		return 4
	case "Вторник":
		return 3
	case "Среда":
		return 2
	case "Четверг":
		return 1
	case "Пятница":
		return 0
	case "Суббота":
		return 6
	case "Воскресенье":
		return 5
	default:
		return -1
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	return strings.EqualFold(strings.ToLower(answer), strings.ToLower(currentDayOfTheWeek()))
}
func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) < 4 || utf8.RuneCountInString(answer) > 4 {
		return false, ErrIncorrectDayInput
	}
	return strings.EqualFold(strings.ToLower(answer), strings.ToLower(dayOrNight())), nil
}
