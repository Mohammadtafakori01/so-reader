package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type SymbolFile struct {
	Filename  string   `json:"filename"`
	Functions []string `json:"functions"`
}

func main() {
	// Check if the operating system is Linux
	if runtime.GOOS != "linux" {
		fmt.Println("Error: This program must be run on Linux.")
		return
	}

	// Check if enough arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: soreader file.so [output_path]")
		return
	}

	// Get the filename from command line argument
	filename := os.Args[1]

	var outputPath string
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	} else {
		// If output path is not provided, use the current directory
		outputPath, _ = os.Getwd()
	}

	// Run nm -D command to get functions from the shared object file
	cmd := exec.Command("nm", "-D", filename)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running nm command: %v\n", err)
		return
	}

	// Parse the output
	functions := parseNMOutput(string(output))

	// Create a SymbolFile struct
	symbolFile := SymbolFile{
		Filename:  filename,
		Functions: functions,
	}

	// Save the functions to a JSON file
	err = saveToJSON(symbolFile, outputPath)
	if err != nil {
		fmt.Printf("Error saving functions to JSON file: %v\n", err)
		return
	}

	fmt.Println("available functions saved to", outputPath)
}

// Function to parse the output of nm command
func parseNMOutput(output string) []string {
	lines := strings.Split(output, "\n")
	var functions []string
	for _, line := range lines {
		// Extract function names
		fields := strings.Fields(line)
		if len(fields) > 2 && (fields[1] == "T" || fields[1] == "D") {
			function := fields[2]
			functions = append(functions, function)
		}
	}
	return functions
}

// Function to save functions to a JSON file
func saveToJSON(symbolFile SymbolFile, outputPath string) error {
	// Marshal the SymbolFile into JSON
	data, err := json.Marshal(symbolFile)
	if err != nil {
		return err
	}

	// Determine the filename for the JSON file
	var filename string
	if outputPath != "" {
		filename = filepath.Base(symbolFile.Filename) + ".json"
	} else {
		filename = filepath.Base(symbolFile.Filename)
	}

	// Construct the full path for the JSON file
	fullPath := filepath.Join(outputPath, filename)

	// Write the JSON data to the file
	err = ioutil.WriteFile(fullPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
