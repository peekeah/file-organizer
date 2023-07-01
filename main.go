package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// error handler
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDirectory(extension string) string {
	if extension == ".exe" {
		return "executables"
	} else if extension == ".jpg" || extension == ".jpeg" || extension == ".png" {
		return "images"
	} else if extension == ".mp3" || extension == ".wav" {
		return "audio"
	} else if extension == ".mp4" || extension == ".avi" || extension == ".mov" {
		return "video"
	} else if extension == ".doc" || extension == ".docx" {
		return "documents"
	} else if extension == ".pdf" {
		return "pdf"
	} else if extension == ".zip" || extension == ".rar" {
		return "compressed"
	} else if extension == ".xlsx" || extension == ".csv" {
		return "documents"
	} else if extension == ".html" || extension == ".htm" {
		return "html"
	} else if extension == ".css" {
		return "css"
	} else if extension == ".js" {
		return "javascript"
	} else {
		return "others"
	}
}

func main() {
	fmt.Println("welcome to file organizer!")

	var sourceDirectory = "e:/file-organizer-test" // source directory
	// read files in folder
	files, err := os.ReadDir(sourceDirectory)
	check(err)

	for _, f := range files {
		var source = filepath.Join(sourceDirectory, f.Name())
		var destDir = filepath.Join(sourceDirectory, getDirectory(path.Ext(f.Name())))
		var destination = filepath.Join(destDir, f.Name())
		err = os.MkdirAll(destDir, os.ModePerm)
		err := os.Rename(source, destination)
		check(err)
	}

	fmt.Println("Succefully organized the files!")
}
