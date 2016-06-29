package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fileSizes := getFileSizes(argsWithoutProg)
	biggest, evens := getBiggestAndEvens(fileSizes)

	if len(evens) == len(argsWithoutProg) {
		fmt.Println("All files are even.")
	} else if len(evens) > 1 {
		fmt.Printf("The biggest are %v\n", evens)
	} else {
		fmt.Printf("The biggest is %s\n", biggest)
	}
}

func getFileSizes(fileNames []string) map[string]int64 {
	fileSizes := make(map[string]int64)
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	for _, fileName := range fileNames {
		// Start a go routine for each file
		waitGroup.Add(1)
		go func(fileName string) {
			defer waitGroup.Done()
			fileSize := getFileSize(fileName)

			// Mutex to avoid concurrent writes to map
			mutex.Lock()
			fileSizes[fileName] = fileSize
			mutex.Unlock()
		}(fileName)
	}

	// Wait for all Go routines to finish
	waitGroup.Wait()
	return fileSizes
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
