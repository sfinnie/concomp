# Golang implementation of ConComp

##To run:

1. Install [Golang](https://golang.org/doc/install) if you haven't already.
2. Clone this repo
3. Run the program:

### Channels
This implementation uses channels to get the sizes back from the Go routines
```
~/concomp/golang$ go run channels.go file1.txt file2.txt file3.txt
```

### WaitGroup & Mutex
This implementation uses a WaitGroup to wait for all Go routines to finish and a Mutex to avoid concurrent writes of the sizes
```
~/concomp/golang$ go run waitgroup.go file1.txt file2.txt file3.txt
```