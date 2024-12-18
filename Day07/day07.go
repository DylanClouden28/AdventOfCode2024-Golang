package main

import (
	"AOC/util"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Operator int

const (
	add Operator = iota
	multiply
	concat
)

type Equation struct {
	result int
	operands []int
}

func generateCombinations(length int, maxNum int) [][]Operator {
	combinations := [][]Operator{}

	var generate func(current []Operator)
	generate = func(current []Operator) {
		if len(current) == length{
			combination := make([]Operator, len(current))
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}

		for i := 0; i <= maxNum; i++{
			generate(append(current, Operator(i)))
		}
	}

	generate([]Operator{})
	return combinations
}

// Make function the evaluates combination of operators following PEMDAS

func filterOperators(o []Operator, toBe Operator) []Operator{
	newSlice := []Operator{}
	for _, value := range o{
		if value == toBe{
			newSlice = append(newSlice, value)
		} else{
			// -1 represents blank space for current operator
			newSlice = append(newSlice, -1)
		}

	}
	return newSlice
}

func computeOperation(o Operator, a int, b int) (int, error){
	if (o == add){
		return a + b, nil
	}
	if (o == multiply){
		return a * b, nil
	}
	if (o == concat){
		concatString := strconv.Itoa(a) + strconv.Itoa(b)
		number, err := strconv.Atoi(concatString)
		return number, err
	}
	return -1, errors.New("no operation found")
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func (equation Equation) findOperators() ([]Operator, bool){
	//possibleOperators := []Operator{add, multiply}
	combinations := generateCombinations(len(equation.operands) - 1, 1)
	for _, combination := range combinations{
		tempOperands := slices.Clone(equation.operands)
		for _, operator := range combination {
			result, err := computeOperation(operator, tempOperands[0], tempOperands[1])
			util.CheckError(err)
			tempOperands[0] = result
			tempOperands = remove(tempOperands,  1)
		}
		if tempOperands[0] == equation.result{
			return combination, true
		}
		}
	return []Operator{}, false
}

func (equation Equation) findOperators_p2() ([]Operator, bool){
	//possibleOperators := []Operator{add, multiply}
	combinations := generateCombinations(len(equation.operands) - 1, 2)
	for _, combination := range combinations{
		tempOperands := slices.Clone(equation.operands)
		for _, operator := range combination {
			result, err := computeOperation(operator, tempOperands[0], tempOperands[1])
			util.CheckError(err)
			tempOperands[0] = result
			tempOperands = remove(tempOperands,  1)
		}
		if tempOperands[0] == equation.result{
			return combination, true
		}
		}
	return []Operator{}, false
}

func getEquations(s []string) []Equation{
	equations := []Equation{}
	for _, line := range s{
		parts := strings.Split(line, ": ")
		result, err := strconv.Atoi(parts[0])
		util.CheckError(err)
		operands := []int{}
		rawOperands := strings.Split(parts[1], " ")
		for _, rawNum := range rawOperands{
			operand, err := strconv.Atoi(rawNum)
			util.CheckError(err)
			operands = append(operands, operand)
		}
		newEquation := Equation{result: result, operands: operands}
		equations = append(equations, newEquation)
	}
	return equations
}

func main(){
	lines := util.GetLines("./input.txt")
	equations := getEquations(lines)
	total := 0
	for _, equation := range equations{
		_, ok := equation.findOperators()
		if !ok{
			continue
		}
		total += equation.result
		// fmt.Println("Equation Result: ", equation.result)
		// fmt.Println("Found valid Operators: ", validOperator);
	}
	fmt.Printf("Total of values is: %d\n", total)

	totalp2 := 0
	for _, equation := range equations{
		_, ok := equation.findOperators_p2()
		if !ok{
			continue
		}
		totalp2 += equation.result
		// fmt.Println("Equation Result: ", equation.result)
		// fmt.Println("Found valid Operators: ", validOperator);
	}
	fmt.Printf("Total of values is: %d\n", totalp2)
}