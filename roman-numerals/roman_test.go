package roman_numerals

import (
	"fmt"
	"testing"
)

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
