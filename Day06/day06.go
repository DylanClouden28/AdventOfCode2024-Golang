package main

import (
	"AOC/util"
	"fmt"
)

type Coordinates struct {
    X, Y int
}

func (c Coordinates) Add(other Coordinates) Coordinates{
	return Coordinates{X: c.X + other.X, Y: c.Y + other.Y}
}

func (c Coordinates) getGridSquare(grid [][]GridSquare) GridSquare{
	return grid[c.Y][c.X]
}

type GridSquare struct {
	obstruction bool
	Coords Coordinates
	explored bool
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var DirectionOffsets = map[Direction]Coordinates{
	North: {X: 0, Y: 1},
	South: {X: 0, Y: -1},
	East: {X: 1, Y: 0},
	West: {X: -1, Y: 0},
}

func getGridSquares(s []string) ([][]GridSquare, Coordinates){
	gridSquares := [][]GridSquare{}
	startPos := Coordinates{}
	for idy, line := range s{
		lineSquares := []GridSquare{}
		for idx, char := range line{
			lineSquares = append(lineSquares, GridSquare{obstruction: char == '#', Coords: Coordinates{X: idx, Y: idy}, explored: false})
			if char == '^'{
				startPos = Coordinates{X: idx, Y: idy}
			}
		}
		gridSquares = append(gridSquares, lineSquares)
	}
	return gridSquares, startPos
}

func exploreGridSquares(grid [][]GridSquare, startPos Coordinates) int{
	currentPos := startPos
	exploring := true
	currentDir := North
	totalExlored := 0
	maxLen := len(grid)
	for exploring {
		offset := DirectionOffsets[currentDir]
		nextPos := currentPos.Add(offset)
		if nextPos.X >= maxLen || nextPos.Y >= maxLen{
			break
		}

		gridSquare := nextPos.getGridSquare(grid)
		if !gridSquare.explored{
			gridSquare.explored = true
			totalExlored += 1
		}

		if gridSquare.obstruction{
			currentDir = Direction((currentDir + 1) % 4)
			continue
		}
		currentPos = nextPos
	}
	return totalExlored
}

func main(){
	lines := util.GetLines("./test.txt")
	gridSquares, startPos := getGridSquares(lines)
	totalExplored := exploreGridSquares(gridSquares, startPos)
	fmt.Println("Total Explored: ", totalExplored)
}