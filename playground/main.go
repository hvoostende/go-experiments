package main

import "fmt"

func main() {
	coolMap := make(map[int]string)

	coolMap[0] = "Nice"
	coolMap[1] = "Work"
	coolMap[2] = "Man"
	coolMap[3] = "!!"

	for k, v := range coolMap {
		fmt.Println(k, v)
	}

	var awesomeMap = map[string]bool{
		"han":   true,
		"joost": false,
	}

	bestMap := map[string]map[string]string{
		"Han": map[string]string{
			"name": "Han van Oostende",
			"sex":  "male"},
		"Jacqueline": map[string]string{
			"name": "Jacqueline Longayroux",
			"sex":  "female"},
	}

	for i := 0; i < len(coolMap); i++ {
		fmt.Println(coolMap[i])
	}

	fmt.Println("is han cool?", awesomeMap["han"])

	// delete(bestMap, "Jacqueline")

	if jacqueline, ok := bestMap["Jacqueline"]; ok {
		fmt.Printf("Jacqueline's full name is %v and she is a %v\n", jacqueline["name"], jacqueline["sex"])
	}

}
