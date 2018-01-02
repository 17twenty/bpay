package bpay

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// ValidateLuhn returns a boolean indicating if the argument was valid according to the Luhn algorithm.
func ValidateLuhn(luhnString string) bool {
	checksumMod := calculateChecksum(luhnString, false) % 10

	return checksumMod == 0
}

func calculateChecksum(luhnString string, double bool) int {
	source := strings.Split(luhnString, "")
	checksum := 0

	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}

func GenerateMOD10V1(input string) string {
	reverseString := reverse(input)
	subtotal := 0
	total := 0
	for i := 0; i < len(reverseString); i++ {
		multiplier := 1
		if i%2 == 0 {
			multiplier = 2
		}
		ival, _ := strconv.Atoi(string(reverseString[i]))
		subtotal = ival * multiplier
		if subtotal >= 10 {
			temp := strconv.Itoa(subtotal)
			temp1, _ := strconv.Atoi(string(temp[0]))
			temp2, _ := strconv.Atoi(string(temp[1]))
			subtotal = temp1 + temp2
		}
		total += subtotal
	}

	checkDigit := (10 - (total % 10)) % 10
	return fmt.Sprintf("%s%d", strings.TrimSpace(input), checkDigit)
}

func GenerateMOD10V5(input string) string {
	strLen := len(input)
	total := 0

	for i := 0; i < strLen; i++ {
		ascInt, _ := strconv.Atoi(string(input[i]))
		total += ascInt * (i + 1)
	}
	checkDigit := total % 10

	return fmt.Sprintf("%s%d", input, checkDigit)
}
