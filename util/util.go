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

	return SplitLines(string(data))
}

func GetAllDataString(filepath string) string {
	file, err := os.Open(filepath)
	CheckError(err)
	data, err := io.ReadAll(file)
	CheckError(err)

	return string(data)
}

func SplitLines(data string) []string{
	lines := strings.Split(string(data), "\n")
	return lines
}