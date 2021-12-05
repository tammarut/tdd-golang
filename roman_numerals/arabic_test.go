package main

import (
	"fmt"
	"testing"
)

/*
  👍 My Solution
*/

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	hasNextValue := index+1 < len(roman)
	isSubtractiveSymbol := string(currentSymbol) == "I" || string(currentSymbol) == "X" || string(currentSymbol) == "C"
	return hasNextValue && isSubtractiveSymbol
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		currentSymbol := roman[i]
		if couldBeSubtractive(i, currentSymbol, roman) {
			nextSymbol := roman[i+1]
			value := AllRomanNumerals.ValueOf(currentSymbol, nextSymbol)
			if value != 0 {
				total += value
				i++
			} else {
				total += AllRomanNumerals.ValueOf(currentSymbol)
			}
		} else {
			total += AllRomanNumerals.ValueOf(currentSymbol)
		}

	}

	return total
}

/*
  🧪 My Test
*/
type roman struct {
	Arabic int
	Roman  string
}

func TestConvertingRomanToArbic(t *testing.T) {
	cases := []roman{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}

	for _, testCase := range cases {
		testCaseName := fmt.Sprintf("%q gets converted to %d", testCase.Roman, testCase.Arabic)
		t.Run(testCaseName, func(t *testing.T) {
			got := ConvertToArabic(testCase.Roman)

			if got != testCase.Arabic {
				t.Errorf("got %d, want %d", got, testCase.Arabic)
			}
		})
	}

}
