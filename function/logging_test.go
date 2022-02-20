package function

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	realPrintln := println

	m.Run()

	println = realPrintln
}

func TestInfo(t *testing.T) {
	expected := "INFO hello world\n"
	logString := "hello world"

	// mock a function
	println = func(v ...interface{}) {
		result := fmt.Sprintln(v...)
		assert.Equal(t, expected, result)
	}

	Info(logString)
}

func TestInfoWithFunction(t *testing.T) {
	expected := "INFO hello world\n"
	logString := "hello world"

	// mock a function
	printFn := func(v ...interface{}) {
		result := fmt.Sprintln(v...)
		assert.Equal(t, expected, result)
	}

	InfoWithFunction(printFn, logString)
}
