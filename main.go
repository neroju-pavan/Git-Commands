package main

import (
	"fmt"
	"strconv"
)
func main() {
	input := " 42 "

	ans, err := strconv.Atoi(input) 
	_ = err

	fmt.Println(ans)
}

