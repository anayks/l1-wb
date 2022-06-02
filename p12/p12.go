package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", GetDataSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func GetDataSet(arr []string) (result map[string]bool) {
	result = map[string]bool{}
	for _, v := range arr {
		_, ok := result[v]
		if !ok {
			result[v] = true
			continue
		}
	}
	return result
}
