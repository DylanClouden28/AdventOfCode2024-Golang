package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkError(e error){
	if e != nil{
		log.Fatal(e)
	}
}

func absInt(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main(){
	file, err := os.Open("./input.txt")
	checkError(err)
	data, err := io.ReadAll(file)
	checkError(err)

	lines := strings.Split(string(data), "\n")

	total := 0
	var firstCol []int
	var secondCol []int
	similarScore := make(map[int]int)
	for _, line := range lines{
		firstNum, err := strconv.Atoi(strings.Split(line, " ")[0])
		checkError(err)
		secondSplit := strings.Split(line, " ")
		SecondNum, err := strconv.Atoi(secondSplit[3])
		checkError(err)
		firstCol = append(firstCol, firstNum)
		secondCol = append(secondCol, SecondNum)
		similarScore[SecondNum] += 1
	}

	sort.Slice(firstCol, func(i, j int) bool {
		return firstCol[i] < firstCol[j]
	})

	sort.Slice(secondCol, func(i, j int) bool {
		return secondCol[i] < secondCol[j]
	})

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