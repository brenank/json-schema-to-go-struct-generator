package main

import (
	"fmt"
	"path/filepath"

	"github.com/brenank/json-schema-to-go-struct-generator/pkg/converter"
	"github.com/brenank/json-schema-to-go-struct-generator/pkg/utils"
)

func main() {
	flags := utils.ParseFlags() // Parsing the cl flags
	files, err := utils.ReadFiles(flags.InputDir)
	if err != nil {
		panic(err)
	}

	outPath, err := filepath.Abs(flags.OutputPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Processing files: %v\n", files)
	err = converter.Convert(files, "models", outPath, false)

	if err != nil {
		panic(err)
	}
}
