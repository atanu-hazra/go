package main

import "fmt"

func main() {
	// myMap()
	mapTest()
}

func myMap() {
	marks := map[string]int{
		"Math":    95,
		"English": 90,
	}
	fmt.Println(marks)

	student := make(map[string]int)

	student["Rick"] = 23
	fmt.Println(student)
	student["Rick"] = 25
	fmt.Println(student)
	delete(marks, "Math")
	fmt.Println(marks)

	age, exists := student["Rick"]
	fmt.Println(age, exists)

	student["John"] = 24
	student["Alice"] = 23
	fmt.Println(student)

	for key, value := range student {
		fmt.Println(key, value)
	}

	users := map[string]map[string]string{
		"rick": {
			"city": "Kolkata",
			"role": "Developer",
		},

		"john": {
			"city": "Delhi",
			"role": "Designer",
		},
	}

	fmt.Println(users["rick"])
	fmt.Println(users["rick"]["role"])

	skills := map[string][]string{

		"Rick": {
			"Go",
			"Docker",
			"Redis",
		},
	}

	fmt.Println(skills["Rick"])

}

func mapTest() {
	// count common words and create a map
	words := []string{
		"go",
		"java",
		"go",
		"python",
		"go",
		"python",
		"java",
	}

	counts := map[string]int64{}

	for _, word := range words {
		// count, exists := counts[word]

		// if (exists) {
		// 	counts[word] = count + 1
		// } else {
		// 	counts[word] = 1
		// }

		counts[word]++
	}

	fmt.Println(counts)
}
