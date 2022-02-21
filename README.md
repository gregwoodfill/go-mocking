# go-mocking

Project to demonstrate go mocking strategies

## Function Mocking

### Replacing with mock at package level

You can mock a function by replacing it by extracting your
function you want mocked to a package level var and replacing it
with a new function in the test.

```go
// production
var println = fmt.Println

func MyFuncThatPrints() {
    println("some message")
}
```

```go
// test
func TestMyFuncThatPrints() {
    expectedLog := "some message"

    // replace the package level function with your own
    println = func (v ...[]interface{}) {
        // assert here
    }
}
```

See the [logging_test](function/logging_test.go)

### Passing in function to be mocked to function under test

```go
// prod file
type printer func(v ...interface{})

func InfoWithFunction(printFn printer, v ...interface{}) {
    var level []interface{}
    level = append(level, "INFO")
    data := append(level, v...)
    printFn(data...)
}
```

```go
// test file
func TestInfoWithFunction(t *testing.T) {
    expected := "INFO hello world\n"
    logString := "hello world"

    // mock a function
    printFn := func(v ...interface{}) {
        result := fmt.Sprintln(v...)
        assert.Equal(t, expected, result)
    }
    // and pass it to the function under test
    InfoWithFunction(printFn, logString)
}
```

## Mocking Methods
