package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	count := flag.Bool("c", false, "count line ocurrences")
	repetead := flag.Bool("d", false, "show only repeated lines")
	unique := flag.Bool("u", false, "show only unique lines")
	flag.Parse()

	args := flag.Args()
	filePath := args[0]

	var lines []string
	var err error

	if filePath == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	} else {
		lines, err = readFile(filePath)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", filePath, err)
		}
	}

	result, counter, err := deleteDuplicates(lines)
	if err != nil {
		log.Fatalf("Error removing duplicates: %v", err)
	}
	printOutput(result, counter, *count, *repetead, *unique)
}

func deleteDuplicates(data []string) ([]string, map[string]int, error) {
	if len(data) == 0 {
		return nil, nil, fmt.Errorf("the data is empty")
	}
	unique := []string{}
	counter := map[string]int{}
	for _, line := range data {
		if line != "" {
			if !slices.Contains(unique, line) {
				unique = append(unique, line)
			}
			if _, ok := counter[line]; ok {
				counter[line] = counter[line] + 1
			} else {
				counter[line] = 1
			}
		}
	}
	return unique, counter, nil
}

func readFile(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func printOutput(lines []string, counter map[string]int, count, repeated, unique bool) {
	output := ""
	if count {
		for _, line := range lines {
			output += fmt.Sprintf("%d %s\n", counter[line], line)
		}
	}
	if repeated {
		for k, v := range counter {
			if v > 1 {
				output += fmt.Sprintf("%s\n", k)
			}
		}
	}
	if unique {
		for k, v := range counter {
			if v == 1 {
				output += fmt.Sprintf("%s\n", k)
			}
		}
	}
	fmt.Println(output)
}
