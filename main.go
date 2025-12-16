package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := " 42 "
	p := strings.TrimSpace(input)

	ans, err := strconv.Atoi(p)
	_ = err

	fmt.Println(ans)
	fmt.Println("hi")
	fmt.Println("ghvscvhc")
}
