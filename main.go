package main

import (
	"fmt"

	"main.go/feature1"
	"main.go/feature2"
)

func main() {
	fmt.Println("Hello Git")

	feature1.Feature1()
	feature2.Feature2()

	fmt.Println("Git main end")
}