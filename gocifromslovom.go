// Package gocifromslovom implements numbers to text in Bulgarian language
package gocifromslovom

const (
	male    = iota
	female  = iota
	neutral = iota
)

var oneInGender = []string{
	"един", "една", "едно",
}

var twoInGender = []string{
	"два", "две", "две",
}

var digitNames = []string{
	"", "еди", "два", "три", "четири",
	"пет", "шест", "седем", "осем", "девет",
}
var digitGroupNames = []string{
	"", "хиляд", "милион", "милиард", "трилион", "квадрилион", "квинтилион",
}

// ConvertMale converts numbers to text in Bulgarian language, for quantities with male gender
func ConvertMale(number int64) string {
	return convert(number, male)
}

// ConvertFemale converts numbers to text in Bulgarian language, for quantities with female gender
func ConvertFemale(number int64) string {
	return convert(number, female)
}

// ConvertNeutral converts numbers to text in Bulgarian language, for quantities with neutral gender
func ConvertNeutral(number int64) string {
	return convert(number, neutral)
}

func digitText(number int64, gender int) string {
	switch number {
	case 1:
		return oneInGender[gender]
	case 2:
		return twoInGender[gender]
	default:
		return digitNames[number]
	}
}

func convert(number int64, gender int) string {

	if number == 0 {
		return "нула"
	}

	sign := ""
	if number < 0 {
		sign = "минус "
		number = -number
	}

	if number < 20 {
		return sign + digitGroupToText(-1, number, "", "", "", gender)
	}

	ret := ""
	for i := 0; i < len(digitGroupNames); i++ {
		var groupText string
		switch i {
		case 0:
			groupText = digitGroupToText(i, number%1000, digitGroupNames[i], "", "", gender)
		case 1:
			groupText = digitGroupToText(i, number%1000, digitGroupNames[i], "а", "и", female)
		default:
			groupText = digitGroupToText(i, number%1000, digitGroupNames[i], "", "а", male)
		}
		if groupText != "" {
			if ret == "" {
				ret = groupText
			} else {
				ret = groupText + " " + ret
			}
		}
		number /= 1000
		if number == 0 {
			break
		}
	}

	return sign + ret
}

func digitGroupToText(groupNo int, groupValue int64, groupName string, single string, plural string, gender int) (ret string) {
	if groupValue == 0 {
		return
	}

	nX99 := groupValue / 100
	n9XX := groupValue % 100
	nX9 := n9XX / 10
	n9X := n9XX % 10

	if nX99 != 0 {
		switch nX99 {
		case 1:
			ret += "сто"
		case 2:
			ret += "двеста"
		case 3:
			ret += "триста"
		default:
			ret += digitNames[nX99] + "стотин"
		}
		if n9XX > 0 {
			ret += " "
			if (n9XX < 20) || (n9X == 0) {
				ret += "и "
			}
		}
	} else if (groupNo == 0) && (n9XX < 20) {
		ret += "и "
	}

	switch nX9 {
	case 0:
		if (groupNo != 1) || (n9X != 1) {
			ret += digitText(n9X, gender)
		}
	case 1:
		if n9X == 0 {
			ret += "десет"
		} else {
			ret += digitNames[n9X] + "надесет"
		}
	default:
		ret += digitText(nX9, male) + "десет"
		if n9X != 0 {
			ret += " и " + digitText(n9X, gender)
		}
	}

	if groupNo > 0 {
		if ret != "" {
			ret += " "
		}
		ret += groupName

		if groupValue > 1 {
			ret += plural
		} else {
			ret += single
		}
	}

	return
}
