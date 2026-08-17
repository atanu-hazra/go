package main

import "fmt"

func main() {
	// myMap()
	mapTest()
}

// Demonstrates core map operations: creation, lookup, deletion, and nesting
func myMap() {
	// 1. Map literal initialization: map[KeyType]ValueType{ key: value, ... }
	marks := map[string]int{
		"Math":    95,
		"English": 90,
	}
	fmt.Println(marks) // output: map[English:90 Math:95]

	// 2. Map creation with make(): allocates an empty, usable map
	student := make(map[string]int)

	// Insert new key-value pair
	student["Rick"] = 23
	fmt.Println(student) // output: map[Rick:23]

	// Update existing key
	student["Rick"] = 25
	fmt.Println(student) // output: map[Rick:25]

	// 3. delete(map, key): removes key from map (does nothing if key doesn't exist)
	delete(marks, "Math")
	fmt.Println(marks) // output: map[English:90]

	// 4. "Comma ok" idiom: safely checks if a key exists
	// age = value (or zero-value if missing), exists = true/false
	age, exists := student["Rick"]
	fmt.Println(age, exists) // output: 25 true

	// Add more students
	student["John"] = 24
	student["Alice"] = 23
	fmt.Println(student) // output: map[Alice:23 John:24 Rick:25]

	// 5. Iterating over map with for-range (Note: map iteration order is random in Go)
	for key, value := range student {
		fmt.Println(key, value)
		// output: Rick 25 \n John 24 \n Alice 23 (order may vary)
	}

	// 6. Nested Maps: map where values are themselves maps
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

	fmt.Println(users["rick"])         // output: map[city:Kolkata role:Developer]
	fmt.Println(users["rick"]["role"]) // output: Developer

	// 7. Map with Slice values: map[KeyType][]ValueType
	skills := map[string][]string{
		"Rick": {
			"Go",
			"Docker",
			"Redis",
		},
	}

	fmt.Println(skills["Rick"]) // output: [Go Docker Redis]
}

// Demonstrates frequency counting using maps
func mapTest() {
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

	fmt.Println(counts) // output: map[go:3 java:2 python:2]
}
