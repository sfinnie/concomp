package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var fileSizes map[string]int64

func main() {
	argsWithoutProg := os.Args[1:]
	fileSizes = make(map[string]int64)

	populateFileSizes(argsWithoutProg)

	biggest, evens := getBiggestAndEvens(fileSizes)

	if len(evens) == len(argsWithoutProg) {
		fmt.Println("All files are even.")
	} else if len(evens) > 1 {
		fmt.Printf("The biggest are %v\n", evens)
	} else {
		fmt.Printf("The biggest is %s\n", biggest)
	}
}

func populateFileSizes(fileNames []string) {
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	for _, fileName := range fileNames {
		waitGroup.Add(1)
		go func(fileName string) {
			defer waitGroup.Done()
			fileSize := getFileSize(fileName)
			mutex.Lock()
			fileSizes[fileName] = fileSize
			mutex.Unlock()
		}(fileName)
	}

	waitGroup.Wait()
}

func getFileSize(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

func getBiggestAndEvens(fileSizes map[string]int64) (string, []string) {
	var max int64
	var biggest string
	var evens []string

	for fileName, fileSize := range fileSizes {
		if fileSize > max {
			max = fileSize
			biggest = fileName
			evens = make([]string, 0)
		}
		if fileSize == max {
			evens = append(evens, fileName)
		}
	}
	return biggest, evens
}
