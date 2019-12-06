package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	input[1] = 1
	input[2] = 2

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			newInput := make([]int, len(input))
			copy(newInput, input)
			newInput[1] = i
			newInput[2] = j
			output, err := processInput(newInput)
			if err != nil {
				//fmt.Printf("error %s\n", err.Error())
				continue
			}
			if output == 19690720 {
				fmt.Printf("i=%d j=%d\n", i, j)
				return
			}
		}
	}
	fmt.Println("done")
}

func getInput() []int {
	var err error
	inputStr := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,5,23,2,13,23,27,1,10,27,31,2,6,31,35,1,9,35,39,2,10,39,43,1,43,9,47,1,47,9,51,2,10,51,55,1,55,9,59,1,59,5,63,1,63,6,67,2,6,67,71,2,10,71,75,1,75,5,79,1,9,79,83,2,83,10,87,1,87,6,91,1,13,91,95,2,10,95,99,1,99,6,103,2,13,103,107,1,107,2,111,1,111,9,0,99,2,14,0,0"
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
	for i := 0; i < len(input); i = i + 4 {
		if input[i] == 1 {
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		} else if input[i] == 2 {
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		} else if input[i] == 99 {
			return input[0], nil
		} else {
			return 0, fmt.Errorf("not a real opcode")
		}
	}
	return 0, fmt.Errorf("went out of bounds")
}
