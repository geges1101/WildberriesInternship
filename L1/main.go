package main

import "fmt"

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	collection := make(map[string]int)
	for _, word := range arr {
		collection[word]++
	}
	fmt.Println(collection)
}
