package main

import (
	"AOC/util"
	"fmt"
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

func main(){
	lines := util.GetLines("./test.txt")

	validLevel := 0

	for _, line := range lines{
		isValid := true
		numbers := strings.Split(line, " ")
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

		canSkip := true
		for index, number := range numbers{
			currentNum, err := strconv.Atoi(number)
			util.CheckError(err)
			isValidLvl := isValidLevel(currentNum, lastNum, isIncreasing)
			if (canSkip && !isValidLvl){
				canSkip = false
				numbers = RemoveIndex(numbers, index)
				currentNum, err = strconv.Atoi(numbers[index])
				util.CheckError(err)
				isValidLvl = isValidLevel(currentNum, lastNum, isIncreasing)
			}
			if (!isValidLvl){
				isValid = isValidLvl
				break
			}
			lastNum = currentNum
		}

		if(!isValid){
			continue
		}
		validLevel += 1

	}

	fmt.Println("Current Num of Valid Level: ", validLevel)
}
