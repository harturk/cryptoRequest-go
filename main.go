package main

import "github.com/harturk/go-first/utils"

func main() {
	content, err := utils.ReadFile("data/model.txt")
	if err != nil {
		panic(err)
	}
	println(content)
}
