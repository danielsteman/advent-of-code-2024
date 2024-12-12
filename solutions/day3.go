package solutions

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func readStringFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

func SolveDay3() (int, error) {
	memory, err := readStringFromFile("data/day3.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read memory: %w", err)
	}

	fmt.Printf("File content length: %d\n", len(memory))
	fmt.Println(memory)

	pattern := `mul\((\d+),(\d+)\)|do\(\)|don'?t\(\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(memory, -1)
	sum := 0
	condition := true

	for _, match := range matches {
		if match[1] != "" && match[2] != "" { // Matches mul(x, y)
			fmt.Println("Matched mul:")
			fmt.Println("  First number:", match[1])
			fmt.Println("  Second number:", match[2])

			first_int, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			second_int, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			if condition {
				sum += first_int * second_int
			}
		} else if match[0] == "do()" {
			fmt.Println("Matched do()")
			condition = true
		} else if match[0] == "don't()" || match[0] == "dont()" {
			fmt.Println("Matched don't()")
			condition = false
		}
	}

	return sum, nil

}
