package main

import (
	"AOC/util"
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

func (equation Equation) findOperators() ([]Operator, bool){
	//possibleOperators := []Operator{add, multiply}
	combinations := generateCombinations(len(equation.operands) - 1, 1)
	for index, combination := range combinations{
		total := 0
		for _, operator := range combination{
			// Run evulation function that uses pemdas
			// if operator == add{
			// 	total += 
			// }
			// if operator == multiply{

			// }
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
		equation.findOperators()
	}
}