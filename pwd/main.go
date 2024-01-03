package main

import (
	"fmt"
	"os"
)

func main () {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to get current working directory")
		return
	}
  fmt.Println(pwd)
}