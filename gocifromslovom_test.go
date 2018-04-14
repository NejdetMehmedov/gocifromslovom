package gocifromslovom_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/dekiland/gocifromslovom"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convey("Testing conversions", t, func() {
		Convey("Testing conversions by example", func() {
			So(ConvertMale(0), ShouldEqual, "нула")
			So(ConvertMale(-1), ShouldEqual, "минус един")
			So(ConvertMale(2), ShouldEqual, "два")
			So(ConvertMale(1), ShouldEqual, "един")
			So(ConvertMale(6), ShouldEqual, "шест")
			So(ConvertMale(10), ShouldEqual, "десет")
			So(ConvertMale(11), ShouldEqual, "единадесет")
			So(ConvertMale(12), ShouldEqual, "дванадесет")
			So(ConvertMale(13), ShouldEqual, "тринадесет")
			So(ConvertMale(20), ShouldEqual, "двадесет")
			So(ConvertMale(37), ShouldEqual, "тридесет и седем")
			So(ConvertMale(42), ShouldEqual, "четиридесет и два")
			So(ConvertFemale(-42), ShouldEqual, "минус четиридесет и две")
			So(ConvertFemale(51), ShouldEqual, "петдесет и една")
			So(ConvertNeutral(61), ShouldEqual, "шестдесет и едно")
			So(ConvertNeutral(99), ShouldEqual, "деветдесет и девет")
			So(ConvertMale(100000000), ShouldEqual, "сто милиона")
			So(ConvertMale(123001), ShouldEqual, "сто двадесет и три хиляди и един")
			So(ConvertMale(-1000000), ShouldEqual, "минус един милион")
			So(ConvertMale(100), ShouldEqual, "сто")
			So(ConvertMale(1024), ShouldEqual, "хиляда двадесет и четири")
		})
	})
}

func ExampleConvert() {
	fmt.Println(ConvertMale(0))
	fmt.Println(ConvertMale(11))
	fmt.Println(ConvertMale(100))
	fmt.Println(ConvertMale(1000))
	fmt.Println(ConvertMale(123))
	fmt.Println(ConvertNeutral(10001))
	fmt.Println(ConvertMale(100014))
	fmt.Println(ConvertFemale(10142))
	fmt.Println(ConvertMale(1032140))
	fmt.Println(ConvertMale(12312140921))
	fmt.Println(ConvertMale(math.MaxInt64))
	fmt.Println(ConvertFemale(-91))
	// Output: нула
	// единадесет
	// сто
	// хиляда
	// сто двадесет и три
	// десет хиляди и едно
	// сто хиляди и четиринадесет
	// десет хиляди сто четиридесет и две
	// един милион тридесет и две хиляди сто и четиридесет
	// дванадесет милиарда триста и дванадесет милиона сто и четиридесет хиляди деветстотин двадесет и един
	// девет квинтилиона двеста двадесет и три квадрилиона триста седемдесет и два трилиона тридесет и шест милиарда осемстотин петдесет и четири милиона седемстотин седемдесет и пет хиляди осемстотин и седем
	// минус деветдесет и една
}
