package main

import (
	"AOC/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)


func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main(){
	lines := util.GetLines("./input.txt")

	total := 0
	var firstCol []int
	var secondCol []int
	similarScore := make(map[int]int)
	for _, line := range lines{
		firstNum, err := strconv.Atoi(strings.Split(line, " ")[0])
		util.CheckError(err)
		secondSplit := strings.Split(line, " ")
		SecondNum, err := strconv.Atoi(secondSplit[3])
		util.CheckError(err)
		firstCol = append(firstCol, firstNum)
		secondCol = append(secondCol, SecondNum)
		similarScore[SecondNum] += 1
	}

	sort.Ints(firstCol)

	sort.Ints(secondCol)

	for index, value := range firstCol{
		total += absInt(value - secondCol[index])
	}

	similarTotal := 0
	for _, value := range firstCol{
		similarTotal += value * similarScore[value]
	}

	fmt.Println("Total for sorted values are: ", total)
	fmt.Println("Total for similar values are: ", similarTotal)
}