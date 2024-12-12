package main

import (
	"aoc2024/solutions"
	"fmt"
)

func main() {
	solution, err := solutions.SolveDay3()
	if err != nil {
		fmt.Println("Error:", err)
        return
	}

    fmt.Println(solution)
}
