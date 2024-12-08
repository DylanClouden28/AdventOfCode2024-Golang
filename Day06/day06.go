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

func (c Coordinates) getGridSquare(grid [][]GridSquare) *GridSquare{
	return &grid[c.Y][c.X]
}

type GridSquare struct {
	obstruction bool
	Coords Coordinates
	explored bool
	lastExploredDir Direction
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var DirectionOffsets = map[Direction]Coordinates{
	North: {X: 0, Y: -1},
	South: {X: 0, Y: 1},
	East: {X: 1, Y: 0},
	West: {X: -1, Y: 0},
}

func getGridSquares(s []string) ([][]GridSquare, Coordinates){
	gridSquares := [][]GridSquare{}
	startPos := Coordinates{}
	for idy, line := range s{
		lineSquares := []GridSquare{}
		for idx, char := range line{
			lineSquares = append(lineSquares, GridSquare{obstruction: char == '#', Coords: Coordinates{X: idx, Y: idy}, explored: false, lastExploredDir: -1})
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

	startSquare := startPos.getGridSquare(grid)
    if !startSquare.explored {
        startSquare.explored = true
        totalExlored += 1
    }

	for exploring {
		offset := DirectionOffsets[currentDir]
		nextPos := currentPos.Add(offset)
		if nextPos.X < 0 || nextPos.Y < 0{
			break
		}
		if nextPos.X >= maxLen || nextPos.Y >= maxLen{
			break
		}

		gridSquare := nextPos.getGridSquare(grid)
		if gridSquare.obstruction{
			if gridSquare.lastExploredDir == currentDir{
				return -1
			}
			gridSquare.lastExploredDir = currentDir
			currentDir = Direction((currentDir + 1) % 4)
			continue
		}
		if !gridSquare.explored{
			gridSquare.explored = true
			totalExlored += 1
		}
		currentPos = nextPos
	}
	return totalExlored
}

func PrintGrid(grid [][]GridSquare, startPos Coordinates){
	for idy, line := range grid{
		for idx, square := range line{
			if idy == startPos.Y && idx == startPos.X{
				fmt.Printf("^")
			} else if square.obstruction{
				fmt.Printf("#")
			} else if square.explored{
				fmt.Printf("X")
			} else{
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func countExplored(grid [][]GridSquare) int{
	total := 0
	for _, line := range grid{
		for _, square := range line{
			if square.explored{
				total += 1
			} 
		}
	}
	return total
}

func totalPossObstrunctions(grid [][]GridSquare, startPos Coordinates) int {
	total := 0
	for idy, line := range grid{
		for idx, square := range line{
			clonedGridSquares := deepCloneGrid(grid)
			if square.obstruction {
				continue
			}
			clonedGridSquares[idy][idx].obstruction = true
			result := exploreGridSquares(clonedGridSquares, startPos)
			if result == -1{
				total += 1
			}
		}
	}
	return total
}

func deepCloneGrid(grid [][]GridSquare) [][]GridSquare {
    newGrid := make([][]GridSquare, len(grid))
    for i := range grid {
        newGrid[i] = make([]GridSquare, len(grid[i]))
        for j := range grid[i] {
            
            newGrid[i][j] = GridSquare{
                obstruction: grid[i][j].obstruction,
                Coords: Coordinates{  
                    X: grid[i][j].Coords.X,
                    Y: grid[i][j].Coords.Y,
                },
                explored: grid[i][j].explored,
                lastExploredDir: grid[i][j].lastExploredDir,
            }
        }
    }
    return newGrid
}

func main(){
	lines := util.GetLines("./input.txt")
	gridSquares, startPos := getGridSquares(lines)
	clonedGridSquares := deepCloneGrid(gridSquares)
	totalExplored := exploreGridSquares(gridSquares, startPos)
	PrintGrid(gridSquares, startPos)
	fmt.Println("Total Explored: ", totalExplored)
	fmt.Println("Acutal Explored: ", countExplored(gridSquares))

	totalObstruction := totalPossObstrunctions(clonedGridSquares, startPos)
	fmt.Println("Total Obstruction: ", totalObstruction)
}