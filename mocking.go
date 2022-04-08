package main

import (
	"gregwoodfill.com/mockExamples/function"
	"gregwoodfill.com/mockExamples/method"
)

func main() {
	function.Info()
	item, err := method.ReadDynamoDBItem("")
	if err != nil {
		return
	}
	function.Info(item)
}
