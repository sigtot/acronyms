package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// getAcronymMapFromFile returns a map containing lists of acronyms with their common sorted word as the key
func getAcronymMapFromFile(fileName string) (map[string][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	acronymMap := make(map[string][]string)
	for scanner.Scan() {
		word := scanner.Text()
		sortedString := sortString(word)
		if acronyms, exists := acronymMap[sortedString]; exists {
			acronymMap[sortedString] = append(acronyms, word)
		} else {
			acronymMap[sortedString] = []string{word}
		}
	}
	return acronymMap, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no filename specified")
		os.Exit(1)
	}
	acronymMap, err := getAcronymMapFromFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for sortedWord, acronyms := range acronymMap {
		if len(acronyms) > 1 && len(sortedWord) > 0 {
			fmt.Println(strings.Join(acronyms, ", "))
		}
	}
}
