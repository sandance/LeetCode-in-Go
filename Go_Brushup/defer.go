package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(f string) *os.File {
	fmt.Println("creating")
	file, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "Hello World data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
