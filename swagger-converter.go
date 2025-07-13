// swagger_json_to_yaml.go
// A simple CLI tool in Go to convert a Swagger/OpenAPI JSON file to YAML.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// Command-line flags
	input := flag.String("i", "", "Path to input Swagger JSON file (default: stdin)")
	output := flag.String("o", "", "Path to output Swagger YAML file (default: stdout)")
	flag.Parse()

	// Read JSON data
	var jsonBytes []byte
	var err error
	if *input == "" {
		jsonBytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Error reading JSON from stdin: %v", err)
		}
	} else {
		jsonBytes, err = ioutil.ReadFile(*input)
		if err != nil {
			log.Fatalf("Error reading JSON file '%s': %v", *input, err)
		}
	}

	// Unmarshal into generic interface
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Marshal to YAML
	yamlBytes, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("Error converting to YAML: %v", err)
	}

	// Write YAML output
	if *output == "" {
		if _, err := os.Stdout.Write(yamlBytes); err != nil {
			log.Fatalf("Error writing YAML to stdout: %v", err)
		}
	} else {
		if err := ioutil.WriteFile(*output, yamlBytes, 0644); err != nil {
			log.Fatalf("Error writing YAML file '%s': %v", *output, err)
		}
	}

	fmt.Fprintln(os.Stderr, "Conversion complete.")
}

