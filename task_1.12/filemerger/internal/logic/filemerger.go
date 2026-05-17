package logic

import (
	"io"
	"os"
	"path/filepath"
)

func GetTextFiles(dir string) []string {
	pattern := filepath.Join(dir, "*.txt")
	files, _ := filepath.Glob(pattern)
	
	result := []string{}
	for _, f := range files {
		if filepath.Base(f) != "combined.txt" {
			result = append(result, f)
		}
	}
	return result
}

func MergeFiles(files []string, dir, outputName string) {
	outputPath := filepath.Join(dir, outputName)
	output, _ := os.Create(outputPath)
	defer output.Close()
	
	for _, file := range files {
		output.WriteString("\n=== " + filepath.Base(file) + " ===\n")
		
		input, _ := os.Open(file)
		io.Copy(output, input)
		input.Close()
		
		output.WriteString("\n")
	}
}