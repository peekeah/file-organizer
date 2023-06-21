package main

import (
	"fmt"
	"os"
	"path"
)

// error handler
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countGoFiles(folder string) (count int) {
	files, err := os.ReadDir(folder)
	check(err)

	for _, f := range files {
		fmt.Println(f.Name())
		if path.Ext(f.Name()) == ".exe" {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("welcome to file organizer!")

	// read files in folder
	directory := "."
	files, err := os.ReadDir(directory)
	check(err)

	for _, f := range files {
		fmt.Println(f.Name())
	}

}
