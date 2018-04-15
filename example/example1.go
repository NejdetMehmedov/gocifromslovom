package main

import (
	"fmt"
	"math"

	"github.com/dekiland/gocifromslovom"
)

func separateFloat64ByDecimalPoint(value float64, precision int) (leftPart int64, rightPart int64) {
	intpart, divpart := math.Modf(value)
	leftPart = int64(intpart)
	rightPart = int64(math.Round(divpart * math.Pow10(precision)))
	return
}

func main() {
	fmt.Println("* Конвертиране цифром в словом на цели числа:")
	fmt.Println(21, gocifromslovom.ConvertMale(21))           // 21 двадесет и един"
	fmt.Println(-31, gocifromslovom.ConvertFemale(-31))       // -31 минус тридесет и една"
	fmt.Println(1000001, gocifromslovom.ConvertMale(1000001)) // 1000001 един милион и едно"
	fmt.Println("* За конвертиране на числа с плаваща запетая:")
	// 101.51 сто и един лева и петдесет и една стотинки"
	fmt.Println(101.51,
		gocifromslovom.ConvertMale(101)+" лева и "+
			gocifromslovom.ConvertFemale(51)+" стотинки") // лев е мъжки, а стотитнка е женски род

	// Пример за конвертиране на суми
	suma := 58.8
	lv, st := separateFloat64ByDecimalPoint(suma, 2) // закръгляване 2 цифри след десетичния знак
	// 58.8 петдесет и осем лева и осемдесет стотинки
	fmt.Println(suma,
		gocifromslovom.ConvertMale(lv)+" лева и "+
			gocifromslovom.ConvertFemale(st)+" стотинки") // лев е мъжки, а стотитнка е женски род

	// Пример за конвертиране на тегло
	teglo := 124.56
	kg, gr := separateFloat64ByDecimalPoint(teglo, 3) // закръгляване 3 цифри след десетичния знак
	// 124.56 сто двадесет и четири килограма петстотин и шестдесет грама
	fmt.Println(teglo,
		gocifromslovom.ConvertMale(kg)+" килограма, "+
			gocifromslovom.ConvertMale(gr)+" грама") // килограм и грам са в мъжки род
}
