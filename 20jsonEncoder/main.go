package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// encodeJson()
	decodeJson()
}

/*
================================================================================
 1. JSON ENCODING (Serialization / Marshaling)

================================================================================

	WHY THIS METHOD IS IMPORTANT:
	- In web applications and microservices, data must be converted from in-memory Go
	  data structures (structs, slices, maps) into a universal text format (JSON)
	  so it can be transmitted over HTTP, saved to a database/file, or shared with
	  frontends (e.g., React, Vue, Mobile apps).
	- Go's "encoding/json" package uses reflection to inspect struct fields and
	  convert them according to struct tags and visibility rules.

	KEY STEPS:
	  Step 1: Define a struct with exported (capitalized) fields and `json` tags.
	  Step 2: Initialize sample data (struct instance or slice).
	  Step 3: Call `json.Marshal()` or `json.MarshalIndent()` to convert data to []byte.
	  Step 4: Check and handle potential encoding errors.
	  Step 5: Output or transmit the resulting JSON byte slice as a string or raw bytes.

================================================================================
*/
func encodeJson() {
	// Struct definition with JSON Tags:
	// NOTE:
	// 1. Fields MUST be Exported (Capitalized) to be visible to the `json` package.
	// 2. `json:"custom_name"` overrides the default field name in the JSON output.
	// 3. `json:"-"` ignores/omits this field completely from JSON (useful for secrets/passwords).
	// 4. `json:"tags,omitempty"` drops the field from JSON if it has a zero/nil/empty value.
	type Course struct {
		Name     string   `json:"title"`          // Renamed to "title" in JSON
		Price    int      `json:"price"`          // Serialized as "price"
		Platform string   `json:"platform"`       // Serialized as "platform"
		Password string   `json:"-"`              // "-" hides this field completely (e.g., sensitive info)
		Tags     []string `json:"tags,omitempty"` // "omitempty" omits the field if slice is nil or empty
	}

	// Step 2: Create sample data (a slice of Course structs)
	courses := []Course{
		{
			Name:     "ReactJS Bootcamp",
			Price:    299,
			Platform: "LearnCodeOnline.in",
			Password: "secretpassword123", // Will not be included in JSON output
			Tags:     []string{"web-dev", "js"},
		},
		{
			Name:     "MERN Stack BootCamp",
			Price:    199,
			Platform: "LearnCodeOnline.in",
			Password: "supersecret123",
			Tags:     []string{"full-stack", "mern"},
		},
		{
			Name:     "Golang Backend",
			Price:    499,
			Platform: "LearnCodeOnline.in",
			Password: "golangpassword",
			Tags:     nil, // Nil value with `omitempty` -> "tags" key will be omitted from JSON
		},
	}

	// Step 3 & 4: Encode data into JSON format
	// Option A: `json.Marshal(courses)` produces compact, minified JSON bytes (best for network payload).
	// finalJson, err := json.Marshal(courses)

	// Option B: `json.MarshalIndent(data, prefix, indent)` produces formatted, human-readable JSON (best for debugging/logging).
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	// Step 5: Convert []byte to string for printing/viewing
	fmt.Println("--- Encoded JSON Output ---")
	fmt.Println(string(finalJson))
}

/*
================================================================================
 2. JSON DECODING (Deserialization / Unmarshaling)

================================================================================

	WHY THIS METHOD IS IMPORTANT:
	- When building REST APIs, consumers send JSON payloads in HTTP request bodies.
	  Similarly, when consuming external APIs or reading config files, data arrives as JSON.
	- We must convert the raw incoming byte slice ([]byte) or JSON stream into Go
	  data structures so we can process, validate, and manipulate the data safely in Go.

	KEY STEPS:
	  Step 1: Validate the incoming raw JSON byte slice using `json.Valid()`.
	  Step 2 (Approach A - Typed Struct): Define a target struct and decode with `json.Unmarshal(data, &struct)`.
	          Best when the JSON schema is known and fixed.
	  Step 3 (Approach B - Dynamic Map): Decode into `map[string]any` when the JSON structure is unknown,
	          heterogeneous, or flexible.

================================================================================
*/
func decodeJson() {
	// Struct definition matching expected JSON keys
	type Course struct {
		Name     string   `json:"title"`
		Price    int      `json:"price"`
		Platform string   `json:"platform"`
		Password string   `json:"-"`
		Tags     []string `json:"tags,omitempty"`
	}

	// Sample incoming raw JSON data (e.g. from an HTTP request body or file)
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Golang Backend",
			"price": 499,
			"website": "LearnCodeOnline.in",
			"tags": ["go", "backend", "concurrency"]
		}
	`)

	// STEP 1: Verify whether the incoming data is syntactically valid JSON.
	// `json.Valid()` checks if the byte slice conforms to valid JSON syntax before attempting to parse.
	checkValid := json.Valid(jsonDataFromWeb)
	if !checkValid {
		fmt.Println("JSON was not valid!")
		return
	}

	// STEP 2: Approach A - Decode directly into a strongly-typed Struct.
	// IMPORTANT: You MUST pass a POINTER (&myCourse) to json.Unmarshal so the function
	// can mutate/populate the fields of the allocated struct variable in-place.
	//
	// NOTE ON KEY MATCHING:
	// - "price" matches `json:"price"` -> populated with 499
	// - "tags" matches `json:"tags"` -> populated with ["go", "backend", "concurrency"]
	// - "coursename" doesn't match `Name` / `json:"title"`, and "website" doesn't match `Platform` / `json:"platform"`
	//   -> Unmatched fields will safely retain their Go zero-values (empty string "").
	var myCourse Course
	err := json.Unmarshal(jsonDataFromWeb, &myCourse)
	if err != nil {
		panic(err)
	}
	fmt.Println("--- Unmarshaled into Struct ---")
	fmt.Printf("%+v\n", myCourse)
	fmt.Printf("Course Name: %s, Price: %d\n\n", myCourse.Name, myCourse.Price)

	// STEP 3: Approach B - Decode dynamic / arbitrary JSON into a Map.
	// When you don't know the exact schema or keys in advance:
	// - Use `map[string]any` (or `map[string]interface{}` in older Go versions).
	// - Numbers are decoded by default as float64, booleans as bool, strings as string, arrays as []any.
	var dynamicData map[string]any
	err = json.Unmarshal(jsonDataFromWeb, &dynamicData)
	if err != nil {
		panic(err)
	}

	fmt.Println("--- Unmarshaled into Dynamic Map ---")
	for key, value := range dynamicData {
		fmt.Printf("Key: %-12s | Value: %-20v | Type: %T\n", key, value, value)
	}
}
