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

func isValidUpdate(update []int, pageOrdering map[int][]int) (bool, int){
	for index, pageNumber := range update{
		previousValues := update[0:index]
		pageRules := pageOrdering[pageNumber]
		for _, prevVal := range previousValues{
			if (slices.Contains(pageRules, prevVal)){
				return false, prevVal
			}
		}
	}
	return true, -1
}

// Loop through find index that is causing problems swap with one it has issues with and run again till fixed
func fixUpdate(update []int, pageOrdering map[int][]int) []int{
	newUpdate := slices.Clone(update)
	isFixed := false
	for (!isFixed){
		isValid, badVal := isValidUpdate(newUpdate, pageOrdering)
		if isValid{
			isFixed = true
			break
		}
		tempUpdate := slices.Clone(update)
		tempUpdate
	}
}

func main(){
	data := util.GetAllDataString("./input.txt")
	pageOrderingRulesRaw := util.SplitLines(strings.Split(data, "\n\n")[0])
	updateStepsRaw := util.SplitLines(strings.Split(data, "\n\n")[1])

	pageOrdering := getPageOrdering(pageOrderingRulesRaw)
	updateSteps := getUpdateSteps(updateStepsRaw)
	totalp1 := 0

	for _, update := range updateSteps{
		isValid, _ := isValidUpdate(update, pageOrdering)
		fmt.Printf("IsValid Update: %t\n", isValid)
		if isValid{
			totalp1 += update[len(update) / 2]
		}
	}
	fmt.Printf("Total: %d\n", totalp1)


	for _, update := range updateSteps{
		isValid, _ := isValidUpdate(update, pageOrdering)
		fmt.Printf("IsValid Update: %t\n", isValid)
		if isValid{
			totalp1 += update[len(update) / 2]
		}
	}
}