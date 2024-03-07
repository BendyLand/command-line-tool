package main

import (
	"fmt"
	"golang-part2/utils"
)

func main() {
	for i := 0; i < 100; i++ {
		displayAll()
	}
}

func displayAll() {
	fmt.Println()
	fmt.Println(utils.ConstructFullRandomLogMessage())
	fmt.Println(utils.ConstructFullRandomLogMessage())
	fmt.Println(utils.ConstructFullRandomLogMessage())
	fmt.Println(utils.ConstructFullRandomLogMessage())
	fmt.Println(utils.ConstructFullRandomLogMessage())
	fmt.Println()
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println(utils.ConstructFullRandomRequestMessage())
	fmt.Println()
}
