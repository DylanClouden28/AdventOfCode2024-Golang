package main

import (
	"AOC/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getPageOrdering(s []string) map[int][]int{
	pageOrdering := make(map[int][]int)
	for _, line := range s{
		parts := strings.Split(line, "|")
		firstNum, err := strconv.Atoi(parts[0])
		util.CheckError(err)
		SecondNum, err := strconv.Atoi(parts[1])
		util.CheckError(err)
		pageOrdering[firstNum] = append(pageOrdering[firstNum], SecondNum)
	}
	return pageOrdering
}

func getUpdateSteps(s []string) [][]int{
	allSteps := [][]int{}
	for _, line := range s{
		updateSteps := []int{}
		parts := strings.Split(line, ",")
		for _, strNum := range parts{
			num, err := strconv.Atoi(strNum)
			util.CheckError(err)
			updateSteps = append(updateSteps, num)
		}
		allSteps = append(allSteps, updateSteps)
	}
	return allSteps
}

func isValidUpdate(update []int, pageOrdering map[int][]int) bool{
	for index, pageNumber := range update{
		previousValues := update[0:index]
		pageRules := pageOrdering[pageNumber]
		for _, prevVal := range previousValues{
			if (slices.Contains(pageRules, prevVal)){
				return false
			}
		}
	}
	return true
}

func main(){
	data := util.GetAllDataString("./input.txt")
	pageOrderingRulesRaw := util.SplitLines(strings.Split(data, "\n\n")[0])
	updateStepsRaw := util.SplitLines(strings.Split(data, "\n\n")[1])

	pageOrdering := getPageOrdering(pageOrderingRulesRaw)
	updateSteps := getUpdateSteps(updateStepsRaw)
	totalp1 := 0

	for _, update := range updateSteps{
		isValid := isValidUpdate(update, pageOrdering)
		fmt.Printf("IsValid Update: %t\n", isValid)
		if isValid{
			totalp1 += update[len(update) / 2]
		}
	}
	fmt.Printf("Total: %d\n", totalp1)
}