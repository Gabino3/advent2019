package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	inputMin := 171309
	inputMax := 643603

	count := 0

	for i := inputMin; i <= inputMax; i++ {
		if numberMeetsCriteria(i) == true {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}

func numberMeetsCriteria(input int) bool {

	if hasDoubleDigits(input) == true && alwaysIncrease(input) == true {
		return true
	}

	return false

}

func hasDoubleDigits(input int) bool {
	digits := len(strconv.Itoa(input))
	lastDigit := 0
	isDouble := false
	isString := false

	for i := digits; i >= 1; i-- {
		digit := getDigitAtPosition(input, i)

		//state machine ftw
		if digit == lastDigit {
			isDouble = isDouble == false && isString == false
			isString = isDouble == false
		} else {
			if isDouble == true {
				break
			}
			isDouble = false
			isString = false
		}

		lastDigit = digit
	}

	return isDouble
}

func alwaysIncrease(input int) bool {

	digits := len(strconv.Itoa(input))
	lastDigit := 0

	for i := digits; i >= 1; i-- {
		digit := getDigitAtPosition(input, i)
		if digit < lastDigit {
			return false
		}
		lastDigit = digit
	}

	return true
}

//return digits with position being
// 654321
func getDigitAtPosition(input, position int) int {
	return (input / int(math.Pow10(position-1))) % 10
}
