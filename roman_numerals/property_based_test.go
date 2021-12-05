package main

import (
	"testing"
	"testing/quick"
)

/*
  🧪 Property based test
*/
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
