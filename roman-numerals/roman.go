package roman_numerals

import (
	"strings"
)

// func ConvertToRoman(arabic int) string {
// 	var result strings.Builder

// 	for i := arabic; i > 0; i-- {
// 		if i == 5 {
// 			result.WriteString("V")
// 			break
// 		}

// 		if i == 4 {
// 			result.WriteString("IV")
// 			break
// 		}
// 		result.WriteString("I")
// 	}

// 	return result.String()
// }

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}
