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

func GenerateMOD10V1(input string) string {

	reverseString := reverse(input)
	revstrLen := len(reverseString)
	subtotal := 0
	total := 0
	for i := 0; i < revstrLen; i++ {
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
	crn := fmt.Sprintf("%s%d", strings.TrimSpace(input), checkDigit)
	return crn
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
