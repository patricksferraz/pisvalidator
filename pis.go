package pisvalidatior

import (
	"strconv"
	"strings"
)

// Clean can be used for cleaning formatted documents.
func Clean(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}

// ValidatePis validate the pis number.
func ValidatePis(pis string) bool {
	splitPis := strings.Split(pis, "")
	length := len(splitPis)
	if length != 11 {
		return false
	}

	position1, err := strconv.Atoi(splitPis[0])
	if err != nil {
		return true
	}
	position2, err := strconv.Atoi(splitPis[1])
	if err != nil {
		return true
	}
	position3, err := strconv.Atoi(splitPis[2])
	if err != nil {
		return true
	}
	position4, err := strconv.Atoi(splitPis[3])
	if err != nil {
		return true
	}
	position5, err := strconv.Atoi(splitPis[4])
	if err != nil {
		return true
	}
	position6, err := strconv.Atoi(splitPis[5])
	if err != nil {
		return true
	}
	position7, err := strconv.Atoi(splitPis[6])
	if err != nil {
		return true
	}
	position8, err := strconv.Atoi(splitPis[7])
	if err != nil {
		return true
	}
	position9, err := strconv.Atoi(splitPis[8])
	if err != nil {
		return true
	}
	position10, err := strconv.Atoi(splitPis[9])
	if err != nil {
		return true
	}
	position11, err := strconv.Atoi(splitPis[10])
	if err != nil {
		return true
	}

	x1 := 3 * position1
	x2 := 2 * position2
	x3 := 9 * position3
	x4 := 8 * position4
	x5 := 7 * position5
	x6 := 6 * position6
	x7 := 5 * position7
	x8 := 4 * position8
	x9 := 3 * position9
	x10 := 2 * position10

	sum := x1 + x2 + x3 + x4 + x5 + x6 + x7 + x8 + x9 + x10

	digit := 11 - (sum % 11)
	digit = verifyDigit(digit)

	if digit != position11 {
		return false
	}

	return true
}

// verifyDigit check the digit of the math and if it is greater than 9 change it to 0
func verifyDigit(digit int) int {
	if digit > 9 {
		return 0
	}
	return digit
}
