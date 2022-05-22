package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("This is a go utils that mimics the ls command")
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Name(), f.IsDir(), f.Mode(), f.ModTime(), f.Size())
	}
}
