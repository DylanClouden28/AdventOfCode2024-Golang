package util

import (
	"io"
	"log"
	"os"
	"strings"
)

func CheckError(e error){
	if e != nil{
		log.Fatal(e)
	}
}

func GetLines(filepath string) []string {

	file, err := os.Open(filepath)
	CheckError(err)
	data, err := io.ReadAll(file)
	CheckError(err)

	lines := strings.Split(string(data), "\n")
	return lines
}