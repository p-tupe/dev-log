// Package parse-json shows how to read/write a json file
// https://pkg.go.dev/encoding/json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const INPUT_FILE = "./input.json"
const OUTPUT_FILE = "./output.json"

func main() {
	readJSONFromFile()
	writeJSONToFile()
}

// Read file reads JSON from INPUT_FILE into inputSchema
func readJSONFromFile() {
	// This reads the entire file in memory; streaming is left as a future exercise
	file, err := os.ReadFile(INPUT_FILE)
	panicIfErr(err)
	fmt.Println("Reading from " + INPUT_FILE)

	//----------- //

	// This is for when you don't know/care about JSON structure in advance (NOT type-safe)
	var inputSchema map[string]any
	err = json.Unmarshal(file, &inputSchema)
	panicIfErr(err)
	fmt.Println(inputSchema)
	// This does mean you have to type check/cast for every key before using it's value
	_ = inputSchema["some-key"].(string)

	//----------- //

	// This is for when you do know/care about JSON structure in advance (IS type-safe)
	typedSchema := struct {
		SomeKey   string `json:"some-key"`
		IgnoreKey string `json:",omitempty"`
	}{}
	err = json.Unmarshal(file, &typedSchema)
	panicIfErr(err)
	// No need of typecasting here
	fmt.Println(typedSchema.SomeKey)
}

// Write file converts outputSchema to JSON and writes to OUTPUT_FILE
func writeJSONToFile() {
	fmt.Println("Writing to " + OUTPUT_FILE)
	outputSchema := []struct {
		SomeKey string `json:"some-key"`
	}{{SomeKey: "some-value"}, {SomeKey: "some-other-value"}}

	output, err := json.Marshal(outputSchema)
	panicIfErr(err)

	err = os.WriteFile(OUTPUT_FILE, output, 0600)
	panicIfErr(err)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
