package main

import (
	"AOC/util"
	"fmt"
)

type frequency rune

type coordinates struct{
	X int
	Y int
} 

type antenna_Map struct{
	antennas map[frequency][]antenna
	width int
	height int
}

func (c coordinates) equals(otherCord coordinates)bool{
	return c.X == otherCord.X && c.Y == otherCord.Y
}

func (c coordinates) outsideOfMap(a antenna_Map)bool{
	if (c.X >= a.width || c.X < 0){
		return true
	}
	if (c.Y >= a.height || c.Y < 0){
		return true
	}
	return false
}

type antenna struct{
	frequency frequency
	coords coordinates
}

func (a antenna) findAntiPos(otherAntenna antenna)coordinates{
	diffX := (a.coords.X - otherAntenna.coords.X)
	diffY := (a.coords.Y - otherAntenna.coords.Y)
	antiPos := coordinates{X: a.coords.X + diffX, Y: a.coords.Y + diffY}
	return antiPos
}

func (a antenna) findMultiAntiPos(otherAntenna antenna, antenna_map antenna_Map)[]coordinates{
	antiNodes := []coordinates{}
	antiPos := otherAntenna.coords
	tempPos := a.coords
	diffX := (a.coords.X - antiPos.X)
	diffY := (a.coords.Y - antiPos.Y)
	for {
		tempPos = coordinates{X: tempPos.X + diffX, Y: tempPos.Y + diffY}
		if tempPos.outsideOfMap(antenna_map){
			break
		}
		antiNodes = append(antiNodes, tempPos)
	}
	return antiNodes
}



func (a *antenna_Map) setAtennaLocations(s []string){
	a.antennas = map[frequency][]antenna{}
	for idy, line := range s{
		for idx, char := range line{
			if char == '.'{
				continue
			}
			freq := frequency(char)
			a.antennas[freq] = append(a.antennas[freq], antenna{frequency: freq, coords: coordinates{X: idx, Y:idy}})
		}
	}
}

func (a *antenna_Map) findAntiNodes_p2()map[coordinates]bool{
	AntiNodes := map[coordinates]bool{}
	for _, antennas := range a.antennas{
		for _, antenna := range antennas{
			AntiNodes[antenna.coords] = true
			for _, otherAntenna := range antennas{
				//If same atenna continue
				if antenna.coords.equals(otherAntenna.coords){
					continue
				}
				antiNodesPos := antenna.findMultiAntiPos(otherAntenna, *a)
				for _, node := range antiNodesPos{
					AntiNodes[node] = true
				}
			}
		}
	}
	return AntiNodes
}

func (a *antenna_Map) findAntiNodes()map[coordinates]bool{
	AntiNodes := map[coordinates]bool{}
	for _, antennas := range a.antennas{
		for _, antenna := range antennas{
			for _, otherAntenna := range antennas{
				//If same atenna continue
				if antenna.coords.equals(otherAntenna.coords){
					continue
				}
				antiNodePos := antenna.findAntiPos(otherAntenna)
				if antiNodePos.outsideOfMap(*a){
					continue
				}
				AntiNodes[antiNodePos] = true
			}
		}
	}
	return AntiNodes
}

func (a *antenna_Map) drawMap(s []string, antiNodes map[coordinates]bool){
	for idy, line := range s{
		for idx, char := range line{
			if char != '.'{
				fmt.Printf("%c", char)
				continue
			}
			_, ok := antiNodes[coordinates{X: idx, Y: idy}]
			if ok{
				fmt.Printf("#")
				continue
			}
			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
}

func main(){
	lines := util.GetLines("./input.txt")
	newMap := antenna_Map{width: len(lines), height: len(lines[0])}
	newMap.setAtennaLocations(lines)
	antiNodes := newMap.findAntiNodes()
	antiNodesp2 := newMap.findAntiNodes_p2()
	//newMap.drawMap(lines, antiNodes)
	fmt.Printf("Total AntiNodes: %d\n", len(antiNodes))
	//newMap.drawMap(lines, antiNodesp2)
	fmt.Printf("Total AntiNodes: %d\n", len(antiNodesp2))
}