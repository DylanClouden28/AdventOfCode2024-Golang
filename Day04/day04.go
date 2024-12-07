package main

import (
	"AOC/util"
	"fmt"
)

func safeCheckCol(lines []string, idx int, idy int, totalCount *int, checkString string){
	maxIdx := len(lines[0])
	maxIdy := len(lines)

	if(idx + 4 <= maxIdx && lines[idx][idy] == checkString[0] && lines[idx + 1][idy] == checkString[1] && lines[idx + 2][idy] == checkString[2] && lines[idx + 3][idy] == checkString[3]){
		fmt.Printf("%d: Found %s col | %d, %d\n", *totalCount, checkString, idx, idy)
		*totalCount += 1
	}

	if(idy + 4 <= maxIdy && idx + 4 <= maxIdx && lines[idx][idy] == checkString[0] && lines[idx + 1][idy + 1] == checkString[1] && lines[idx + 2][idy + 2] == checkString[2] && lines[idx + 3][idy + 3] == checkString[3]){
		fmt.Printf("%d: Found %s dag+ | %d, %d\n", *totalCount, checkString, idx, idy)
		*totalCount += 1
	}
	
	if(idy >= 3 && idx + 4 <= maxIdx && lines[idx][idy] == checkString[0] && lines[idx + 1][idy - 1] == checkString[1] && lines[idx + 2][idy - 2] == checkString[2] && lines[idx + 3][idy - 3] == checkString[3]){
		fmt.Printf("%d: Found %s dag- | %d, %d\n", *totalCount, checkString, idx, idy)
		*totalCount += 1
	}
}

func checkMAS(lines []string, idx int, idy int) bool{
	maxIdx := len(lines[0])
	maxIdy := len(lines)
	if (idy + 2 >= maxIdy){
		return false
	}
	if (idx + 2 >= maxIdx){
		return false
	}
	
	//Check for center A
	if (lines[idx + 1][idy + 1] != 'A'){
		return false
	}

	// Check dag +
	dagPos1 := (lines[idx][idy] == 'M' && lines[idx + 2][idy + 2] == 'S')
	dagPos2 := (lines[idx][idy] == 'S' && lines[idx + 2][idy + 2] == 'M')
	if (!dagPos1 && !dagPos2){
		return false
	}

	// Check dag -
	dagNeg1 := (lines[idx][idy + 2] == 'M' && lines[idx + 2][idy] == 'S')
	dagNeg2 := (lines[idx][idy + 2] == 'S' && lines[idx + 2][idy] == 'M')
	if (!dagNeg1 && !dagNeg2){
		return false
	}

	return true

}

func main(){
	lines := util.GetLines("./input.txt")
	totalCount := 0
	totalCountp2 := 0

	for idx, line := range lines{
		for idy, _ := range line{
			maxIdy := len(line)
			curSlice := ""
			if (idy + 4 <= maxIdy){
				curSlice = line[idy:idy+4]
			}
			if (curSlice == "XMAS"){
				fmt.Printf("%d: Found XMAS row | %d, %d\n", totalCount, idx, idy)
				totalCount += 1
			}
			if (curSlice == "SAMX"){
				fmt.Printf("%d: Found SAMX row | %d, %d\n", totalCount, idx, idy)
				totalCount += 1
			}
			safeCheckCol(lines, idx, idy, &totalCount, "XMAS")
			safeCheckCol(lines, idx, idy, &totalCount, "SAMX")
		}
	}

	for idx, line := range lines{
		for idy, _ := range line{
			isMas := checkMAS(lines, idx, idy)
			if (isMas){
				fmt.Printf("Found MAS | %d, %d\n", idx, idy)
				totalCountp2 += 1
			}
		}
	}

	println("Total XMAS: ", totalCount)
	println("Total MAS: ", totalCountp2)
}	