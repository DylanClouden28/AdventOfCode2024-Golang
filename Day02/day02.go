package main

import (
	"AOC/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func isValidLevel(currentNum int, lastNum int, isIncreasing bool) bool{
	diffInc := (currentNum - lastNum)
	if (isIncreasing && (diffInc > 3)) || (isIncreasing && !(diffInc > 0)){
		return false
	} 
	diffDec := (lastNum - currentNum)
	if (!isIncreasing && (diffDec > 3)) || (!isIncreasing &&!(diffDec > 0)){
		return false
	} 
	return true
}


func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

func createCombinations(numbers []string) [][]string{
	var combinations [][]string
	for index, _ := range numbers{
		newArr := slices.Clone(numbers)
		slicedArr := RemoveIndex(newArr, index)
		combinations = append(combinations, slicedArr)
	}
	return combinations
}

func runThroughNumbers(numbers []string) bool{
	isValid := true
	isIncreasing := true
	firstNum, err := strconv.Atoi(numbers[0])
	util.CheckError(err)
	secondNum, err := strconv.Atoi(numbers[1])
	util.CheckError(err)
	lastNum := firstNum - 1
	if firstNum > secondNum{
		isIncreasing = false
		lastNum = firstNum + 1
	}

	
	for _, number := range numbers{
		currentNum, err := strconv.Atoi(number)
		util.CheckError(err)
		isValidLvl := isValidLevel(currentNum, lastNum, isIncreasing)
		if (!isValidLvl){
			isValid = isValidLvl
			break
		}
		lastNum = currentNum
	}

	return isValid
}

func main(){
	lines := util.GetLines("./input.txt")

	validLevelP1 := 0
	validLevelP2 := 0

	for _, line := range lines{
		numbers := strings.Split(line, " ")
		isValid := runThroughNumbers(numbers)

		if(isValid){
			validLevelP1 += 1
		}

		combinations := createCombinations(numbers)
		for _, combination := range combinations{
			isValid := runThroughNumbers(combination)

			if(isValid){
				validLevelP2 += 1
				break
			}
		}

	}

	fmt.Println("Part 1 Current Num of Valid Level: ", validLevelP1)
	fmt.Println("Part 2 Current Num of Valid Level: ", validLevelP2)
}
