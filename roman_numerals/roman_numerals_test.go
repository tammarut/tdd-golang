package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
  👍 My Solution
*/
// Old soulution
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

// New soulution
type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

// func (r RomanNumerals) ValueOf(symbol string) int {
func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var AllRomanNumerals = RomanNumerals{
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

	for _, numeral := range AllRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

/*
  🧪 My Test
*/
type Roman struct {
	Arabic int
	Roman  string
}

func TestRomanNumerals(t *testing.T) {
	cases := []Roman{
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
		testCaseName := fmt.Sprintf("%d gets converted to %q", testCase.Arabic, testCase.Roman)
		t.Run(testCaseName, func(t *testing.T) {
			got := ConvertToRoman(testCase.Arabic)

			if got != testCase.Roman {
				t.Errorf("got %q, want %q", got, testCase.Roman)
			}
		})
	}
}
