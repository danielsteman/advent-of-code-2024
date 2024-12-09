package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2, err := readListsFromFile("data/day_1_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var sum_of_difference float64
	for i := 0; i < len(list1); i++ {
		sum_of_difference += math.Abs(float64(list1[i]) - float64(list2[i]))
	}

	fmt.Printf("Sum of difference: %f\n", sum_of_difference)

	m := make(map[int]int)

	for _, val := range list2 {
		m[val]++
	}

	var similarity_score int

	for _, val := range list1 {
		similarity_score += val * m[val]
	}

	fmt.Printf("Similarity score: %d\n", similarity_score)
}

func readListsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		if len(parts) == 2 {
			val1, err1 := strconv.Atoi(parts[0])
			val2, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return nil, nil, fmt.Errorf("failed to parse integers: %s", line)
			}
			list1 = append(list1, val1)
			list2 = append(list2, val2)
		} else {
			fmt.Println("Line format incorrect:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return list1, list2, nil
}

