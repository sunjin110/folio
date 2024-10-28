package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("===== hello lambda thumbnail")
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
}
