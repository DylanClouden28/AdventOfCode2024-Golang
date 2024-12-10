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
		for _, opType := range []Operator{multiply, add}{
			for idxOp, operator := range filterOperators(combination, opType){
				if operator == -1{
					continue
				}
				result, err := computeOperation(operator, tempOperands[idxOp], tempOperands[idxOp + 1])
				util.CheckError(err)
				tempOperands[idxOp] = result
				tempOperands = remove(tempOperands, idxOp + 1)
			}
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
	lines := util.GetLines("./test.txt")
	equations := getEquations(lines)
	for _, equation := range equations{
		validOperator, ok := equation.findOperators()
		if !ok{
			continue
		}
		fmt.Println("Found valid Operators: ", validOperator);
	}
}