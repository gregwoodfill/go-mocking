package function

import "log"

var println = log.Println

func Info(v ...interface{}) {
	var level []interface{}
	level = append(level, "INFO")
	data := append(level, v...)
	println(data...)
}

type printer func(v ...interface{})

func InfoWithFunction(printFn printer, v ...interface{}) {
	var level []interface{}
	level = append(level, "INFO")
	data := append(level, v...)
	printFn(data...)
}
