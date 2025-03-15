package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Colors for beautiful terminal logs
const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	reset  = "\033[0m"
)

func main() {
	// Definition of command line arguments
	dirPtr := flag.String("dir", "", "Base directory for creating the barrel file (required)")
	outputPtr := flag.String("o", "", "Output file path (optional, default: <folder>_exports.dart)")
	flag.Parse()

	// Directory validation
	if *dirPtr == "" {
		fmt.Println(red + "❌ Error: You must specify a directory with `-dir`" + reset)
		os.Exit(1)
	}

	dirPath := *dirPtr
	outputFile := *outputPtr

	// If the output is not defined, generate <folder_name>_exports.dart
	if outputFile == "" {
		baseName := filepath.Base(dirPath) // Folder name
		outputFile = filepath.Join(dirPath, baseName+"_exports.dart")
	}

	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println(red + "🚨 Error: The folder “ + dirPath + ” was not found.”" + reset)
		os.Exit(1)
	}

	var dartFiles []string

	// Walk through the directory and get all .dart files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".dart") {
			if strings.HasSuffix(info.Name(), ".freezed.dart") || strings.HasSuffix(info.Name(), ".g.dart") || filepath.Base(path) == filepath.Base(outputFile) {
				return nil // Ignores automatically generated files and the barrel file itself
			}
			relativePath, _ := filepath.Rel(dirPath, path) // Relative path
			dartFiles = append(dartFiles, relativePath)
		}
		return nil
	})

	if err != nil {
		fmt.Println(red + "❌ Error reading files:", err, reset)
		os.Exit(1)
	}

	// If no .dart files are found, exit the program
	if len(dartFiles) == 0 {
		fmt.Println(yellow + "⚠️ No .dart file found in " + dirPath + reset)
		os.Exit(1)
	}

	// Sort the files alphabetically
	sort.Strings(dartFiles)

	// Create the barrel file
	var builder strings.Builder
	builder.WriteString("// 📦 Automatically generated file. DO NOT EDIT MANUALLY.\n\n")
	for _, file := range dartFiles {
		builder.WriteString(fmt.Sprintf("export '%s';\n", file))
	}

	// Write the barrel file
	err = os.WriteFile(outputFile, []byte(builder.String()), 0644)
	if err != nil {
		fmt.Println(red + "❌ Error writing the file:", err, reset)
		os.Exit(1)
	}

	fmt.Println(green + "✅ File " + filepath.Base(outputFile) + " updated successfully!" + reset)

}