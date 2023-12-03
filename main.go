package main

import (
	"fmt"
	"os"

	fileIO "github.com/harturk/go-first/fileIO"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + `\` + "data" + `\` + "model.txt"
	content, err := fileIO.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getwd())
	fmt.Println(content)
	var newContent = fmt.Sprintf("Original: %v\n Double Origina: %v%v", content, content, content)
	fileIO.WriteFile(filePath+".output.txt", newContent)
}
