package main

import (
	"AOC/util"
	"fmt"
	"regexp"
	"strconv"
)


func main(){
	lines := util.GetLines("./input.txt")

	re := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)`)
	re_do := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)|do\(\)|don't\(\)`)
	total := 0
	total_p2 := 0
	for _, line := range lines{
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches{
			firstNum, err := strconv.Atoi(match[1])
			util.CheckError(err)
			secondNum, err := strconv.Atoi(match[2])
			util.CheckError(err)
			// fmt.Printf("First Num: %d * %d\n", firstNum, secondNum)
			total += firstNum * secondNum
		}

		matches_do := re_do.FindAllStringSubmatch(line, -1)
		enabled := true
		for _, match := range matches_do{
			value := match[0]
			if (value == "do()"){
				enabled = true
			}else if (value == "don't()"){
				enabled = false
			} else if (enabled) {
				firstNum, err := strconv.Atoi(match[1])
				util.CheckError(err)
				secondNum, err := strconv.Atoi(match[2])
				util.CheckError(err)
				total_p2 += firstNum * secondNum
			}

		}
	}
	fmt.Println("Total P1: ", total)
	fmt.Println("Total P2: ", total_p2)
}