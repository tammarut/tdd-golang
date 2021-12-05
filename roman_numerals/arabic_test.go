package main

import (
	"fmt"
	"testing"
)

/*
  👍 My Solution
*/
type windowedRoman string

func (windowR windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(windowR); i++ {
		currentSymbol := windowR[i]
		hasNextValue := i+1 < len(windowR)
		nextValue := func() byte {
			if !hasNextValue {
				return 0
			}
			return windowR[i+1]
		}()

		if hasNextValue && isSubtractive(currentSymbol) && AllRomanNumerals.Exists(currentSymbol, nextValue) {
			symbols = append(symbols, []byte{currentSymbol, nextValue})
			i++
		} else {
			symbols = append(symbols, []byte{currentSymbol})
		}
	}

	return
}

func isSubtractive(symbol uint8) bool {
	return string(symbol) == "I" || string(symbol) == "X" || string(symbol) == "C"
}

func ConvertToArabic(roman string) int {
	total := 0
	symbols := windowedRoman(roman).Symbols()

	for _, s := range symbols {
		total += AllRomanNumerals.ValueOf(s...)
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
