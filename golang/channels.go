package main

import (
	"fmt"
	"log"
	"os"
)

// FileInfo contains information about a file
type FileInfo struct {
	name string
	size int64
}

func main() {
	argsWithoutProg := os.Args[1:]
	fileInfosChannel := make(chan FileInfo)

	sendFileSizesOnChannelConcurrently(argsWithoutProg, fileInfosChannel)
	fileInfos := getFileInfosFromChannel(fileInfosChannel, len(argsWithoutProg))
	biggest := getBiggest(fileInfos)
	evens := getEvens(biggest, fileInfos)

	if len(evens) == len(fileInfos) {
		fmt.Println("All the files are even")
	} else if len(evens) > 1 {
		fmt.Printf("The biggest are %v\n", evens)
	} else {
		fmt.Printf("The biggest is %s\n", biggest.name)
	}
}

func sendFileSizesOnChannelConcurrently(filenames []string, channel chan FileInfo) {
	for _, filename := range filenames {
		go getFileSize(filename, channel)
	}
}

func getFileSize(filename string, sizeMessages chan FileInfo) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeMessages <- FileInfo{name: filename, size: fi.Size()}
}

func getFileInfosFromChannel(channel chan FileInfo, countExpected int) []FileInfo {
	var fileInfos []FileInfo
	for i := 0; i < countExpected; i++ {
		fileInfo := <-channel
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos
}

func getBiggest(fileInfos []FileInfo) FileInfo {
	var biggest FileInfo
	for _, fileInfo := range fileInfos {
		if fileInfo.size > biggest.size {
			biggest = fileInfo
		}
	}
	return biggest
}

func getEvens(biggest FileInfo, fileInfos []FileInfo) []string {
	var evens []string
	for _, fileInfo := range fileInfos {
		if fileInfo.size == biggest.size {
			evens = append(evens, fileInfo.name)
		}
	}
	return evens
}
