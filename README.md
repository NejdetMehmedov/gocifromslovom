# gocifromslovom
Конвертиране на число в думи на български (Цифром -> Словом) на Go (Golang)

## Пример за реална употреба
Конвертиране цифром -> словом, на суми, тегло и др.
```go
package main

import (
	"fmt"
	"math"

	"github.com/dekiland/gocifromslovom"
)

// Помощна функция за отделяне на значещите цифри преди и след десетичният знак
// Параметри: Първия е числото с плаваща запетая
//            Втория е броя на цифрите след десетичната запетая
// Резултати: Първия е цялото число преди десетичният знак
//            Втория е цялото число след десетичния знак закръглено съгласно втория параметър
func separateFloat64ByDecimalPoint(value float64, precision int) (leftPart int64, rightPart int64) {
	intpart, divpart := math.Modf(value)
	leftPart = int64(intpart)
	rightPart = int64(math.Round(divpart * math.Pow10(precision)))
	return
}

func main() {

	//*** Примери за конвертиране цифром в словом на цели числа

	fmt.Println(21, gocifromslovom.ConvertMale(21))
	// Изход: 21 двадесет и един"

	fmt.Println(-31, gocifromslovom.ConvertFemale(-31))
	// Изход: -31 минус тридесет и една"

	fmt.Println(1000001, gocifromslovom.ConvertMale(1000001))
	// Изход: 1000001 един милион и едно"

	//*** Примери за конвертиране цифром в словом на числа с плаваща запетая, суми, тегло и др.

	fmt.Println(101.51,
		gocifromslovom.ConvertMale(101)+" лева и "+
			gocifromslovom.ConvertFemale(51)+" стотинки") // лев е мъжки, а стотинка е женски род
	// Изход: 101.51 сто и един лева и петдесет и една стотинки"

	// Пример за конвертиране на суми
	suma := 58.8
	lv, st := separateFloat64ByDecimalPoint(suma, 2) // закръгляване 2 цифри след десетичния знак

	fmt.Println(suma,
		gocifromslovom.ConvertMale(lv)+" лева и "+
			gocifromslovom.ConvertFemale(st)+" стотинки") // лев е мъжки, а стотинка е женски род
	// Изход: 58.8 петдесет и осем лева и осемдесет стотинки

	// Пример за конвертиране на тегло
	teglo := 124.56089
	kg, gr := separateFloat64ByDecimalPoint(teglo, 3) // закръгляване 3 цифри след десетичния знак

	fmt.Println(teglo,
		gocifromslovom.ConvertMale(kg)+" килограма, "+
			gocifromslovom.ConvertMale(gr)+" грама") // килограм и грам са в мъжки род
	// Изход: 124.56089 сто двадесет и четири килограма, петстотин шестдесет и един грама
}
```

