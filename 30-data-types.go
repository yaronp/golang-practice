package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _= strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank"

	scanner := bufio.NewScanner(os.Stdin)

	var inputString string
	scanner.Scan()
	inputString = scanner.Text()

	var i2 uint64
	var d2 float64
	var s2 string


	// fmt.Scanln(&inputString)
	i2,_ = strconv.ParseUint(inputString, 10, 64)

	// fmt.Scanln(&inputString)
	scanner.Scan()
	inputString = scanner.Text()

	d2, _ = strconv.ParseFloat(inputString, 64)

	// fmt.Scanln(&s2)
	scanner.Scan()
	s2 = scanner.Text()

	fmt.Println(int(i2+i))
	fmt.Println(float64(d+d2))
	fmt.Println(s + " " + s2)


}

