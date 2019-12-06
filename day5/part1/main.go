package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	firstParamPosition  = 2
	secondParamPosition = 3
	thirdParamPosition  = 4
)

func main() {

	output, err := processInput(getInput())
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return
	}
	fmt.Printf("Output: %d\n", output)

	fmt.Println("done")
}

func getInput() []int {
	var err error
	inputStr := "3,225,1,225,6,6,1100,1,238,225,104,0,1101,90,64,225,1101,15,56,225,1,14,153,224,101,-147,224,224,4,224,1002,223,8,223,1001,224,3,224,1,224,223,223,2,162,188,224,101,-2014,224,224,4,224,1002,223,8,223,101,6,224,224,1,223,224,223,1001,18,81,224,1001,224,-137,224,4,224,1002,223,8,223,1001,224,3,224,1,223,224,223,1102,16,16,224,101,-256,224,224,4,224,1002,223,8,223,1001,224,6,224,1,223,224,223,101,48,217,224,1001,224,-125,224,4,224,1002,223,8,223,1001,224,3,224,1,224,223,223,1002,158,22,224,1001,224,-1540,224,4,224,1002,223,8,223,101,2,224,224,1,223,224,223,1101,83,31,225,1101,56,70,225,1101,13,38,225,102,36,192,224,1001,224,-3312,224,4,224,1002,223,8,223,1001,224,4,224,1,224,223,223,1102,75,53,225,1101,14,92,225,1101,7,66,224,101,-73,224,224,4,224,102,8,223,223,101,3,224,224,1,224,223,223,1101,77,60,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,7,226,677,224,1002,223,2,223,1005,224,329,1001,223,1,223,1007,226,677,224,1002,223,2,223,1005,224,344,101,1,223,223,108,226,226,224,1002,223,2,223,1006,224,359,101,1,223,223,7,226,226,224,102,2,223,223,1005,224,374,101,1,223,223,8,677,677,224,1002,223,2,223,1005,224,389,1001,223,1,223,107,677,677,224,102,2,223,223,1006,224,404,101,1,223,223,1107,677,226,224,102,2,223,223,1006,224,419,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,434,1001,223,1,223,7,677,226,224,102,2,223,223,1006,224,449,1001,223,1,223,1107,226,226,224,1002,223,2,223,1005,224,464,101,1,223,223,1108,226,677,224,102,2,223,223,1005,224,479,101,1,223,223,1007,677,677,224,102,2,223,223,1006,224,494,1001,223,1,223,1107,226,677,224,1002,223,2,223,1005,224,509,101,1,223,223,1007,226,226,224,1002,223,2,223,1006,224,524,101,1,223,223,107,226,226,224,1002,223,2,223,1005,224,539,1001,223,1,223,1108,677,677,224,1002,223,2,223,1005,224,554,101,1,223,223,1008,677,226,224,102,2,223,223,1006,224,569,1001,223,1,223,8,226,677,224,102,2,223,223,1005,224,584,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,599,1001,223,1,223,108,677,677,224,102,2,223,223,1006,224,614,1001,223,1,223,108,226,677,224,102,2,223,223,1005,224,629,101,1,223,223,8,677,226,224,102,2,223,223,1005,224,644,101,1,223,223,107,677,226,224,1002,223,2,223,1005,224,659,101,1,223,223,1108,677,226,224,102,2,223,223,1005,224,674,1001,223,1,223,4,223,99,226"
	inputStrArr := strings.Split(inputStr, ",")
	input := make([]int, len(inputStrArr))
	for i, val := range inputStrArr {
		input[i], err = strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error converting index %d val %s with err %s", i, val, err.Error())
		}
	}
	return input
}
func processInput(input []int) (int, error) {

	instructionStep := 4
	for i := 0; i < len(input); {
		log("full OPCode: %d", input[i])
		log("OPCode: %d", getOPCode(input[i]))
		switch getOPCode(input[i]) {
		case 1:
			log("first mode: %d, second mode %d, first val: %d, second val: %d", getDigit(input[i], firstParamPosition), getDigit(input[i], secondParamPosition), getParamValue(input, i+1, getDigit(input[i], firstParamPosition)), getParamValue(input, i+2, getDigit(input[i], secondParamPosition)))
			input[input[i+3]] = getParamValue(input, i+1, getDigit(input[i], firstParamPosition)) + getParamValue(input, i+2, getDigit(input[i], secondParamPosition))
			instructionStep = 4
		case 2:
			input[input[i+3]] = getParamValue(input, i+1, getDigit(input[i], firstParamPosition)) * getParamValue(input, i+2, getDigit(input[i], secondParamPosition))
			instructionStep = 4
		case 3:
			num, err := getUserInput()
			if err != nil {
				return 0, err
			}
			input[input[i+1]] = num
			instructionStep = 2
		case 4:
			fmt.Printf("%d\n", getParamValue(input, i+1, getDigit(input[i], firstParamPosition)))
			instructionStep = 2
		case 99:
			return input[0], nil
		default:
			return 0, fmt.Errorf("not a real opcode")
		}
		i = i + instructionStep
	}
	return 0, fmt.Errorf("went out of bounds")
}

func log(message string, args ...interface{}) {
	return
	fmt.Printf(message+"\n", args...)
}

func getUserInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Gimme some input")
	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(strings.Replace(text, "\n", "", -1))
	if err != nil {
		return 0, err
	}

	return num, nil
}

func getParamValue(input []int, paramPosition, mode int) int {
	//fmt.Printf("getParamVal - inputLen: %d, paramPosition: %d, mode %d\n", len(input), paramPosition, mode)
	if mode == 0 {
		return input[input[paramPosition]]
	} else if mode == 1 {
		return input[paramPosition]
	}

	return -1
}

func getDigit(num, position int) int {
	return (num / int(math.Pow10(position))) % 10
}

func getOPCode(num int) int {
	return num % 100
}
