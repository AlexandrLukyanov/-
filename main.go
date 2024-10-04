package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	fmt.Println("Калькулятор массы тела")
	for {
		userKg, userHeigh := getUserInput()
		IMT, err := calculIMT(userKg, userHeigh)
		if err != nil {
			fmt.Println("Не заданны параметры для расчета ")
			continue
		}
		outputResult(IMT)
		isChekRepit := nextIndex()
		if !isChekRepit {
			break
		}
	}

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дифицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")

	}
}

func calculIMT(userKg float64, userHeigh float64) (float64, error) {
	if userKg <= 0 || userHeigh <= 0 {
		return 0, errors.New("No_PARAMS_ERORR")
	}
	IMT := userKg / math.Pow(userHeigh/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userKg float64
	var userHeigh float64
	fmt.Println("Какой у вас рост :")
	fmt.Scan(&userHeigh)
	fmt.Println("Какой у вас вес :")
	fmt.Scan(&userKg)
	return userKg, userHeigh
}

func nextIndex() bool {
	var userChek string
	fmt.Print("Вы желаете продолжить да/нет: ")
	fmt.Scan(&userChek)
	if userChek == "да" {
		return true
	}
	return false

}
