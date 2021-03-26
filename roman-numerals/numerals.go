package numerals

import (
	"errors"
	"strings"
)

var ErrOutOfRange = errors.New("value is out of valid range [1; 3999]")

// ConvertToRoman converts an Arabic number to a Roman numeral
func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", ErrOutOfRange
	}

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

// ConvertToArabic converts a Roman numeral to an Arabic number
func ConvertToArabic(roman string) uint16 {
	var total uint16
	for i := len(roman) - 1; i >= 0; {
		currNumber := allRomanNumerals.ValueOf(roman[i])

		if i > 0 {
			prevNumber := allRomanNumerals.ValueOf(roman[i-1])
			if prevNumber < currNumber {
				// IV, IX, XL etc.
				total += currNumber - prevNumber
				i -= 2
				continue
			}
		}
		total += currNumber
		i--
	}

	return total
}

type romanNumeral struct {
	Value  uint16
	Symbol string
}

type romanNumerals []romanNumeral

var allRomanNumerals = romanNumerals{
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

func (r romanNumerals) ValueOf(symbol ...byte) uint16 {
	roman := string(symbol)
	for _, s := range r {
		if s.Symbol == roman {
			return s.Value
		}
	}

	return 0
}
