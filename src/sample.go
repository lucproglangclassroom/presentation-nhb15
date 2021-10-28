package main

import "fmt"

func main() {
	frequencyMap := make(map[string]int)
	stringList := []string{"red", "blue", "blue", "green", "purple", "blue", "green"}

	for _, stringExample := range stringList {
		current_frequency, ok := frequencyMap[stringExample]
		if ok {
			frequencyMap[stringExample] = current_frequency + 1
		} else {
			frequencyMap[stringExample] = 1
		}
	}

	fmt.Println(fmt.Sprintf("%v", frequencyMap))
}
