package main

import (
	"os"
)

func main() {
	panic("a problem")

	_, err := os.Create("E:\\1.txt")
	if err != nil {
		panic(err)
	}
}
