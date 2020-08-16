package main

import (
	"fmt"
	"learning/go-tools/pkg/mockery/simple-case/mocks"
)

func main() {
	stringer := new(mocks.Stringer)
	stringer.On("String", "hello", "dungnh", "12").Return(func(arg string, args ...interface{}) string {
		for _, value := range args {
			arg = arg + value.(string)
		}
		return arg + " dungnh"
	})

	result := stringer.String("hello", "dungnh", "12")
	fmt.Println(result)
}
