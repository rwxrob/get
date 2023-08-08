package get_test

import (
	"fmt"

	"github.com/rwxrob/get"
)

func ExampleString_env() {

	it, err := get.String(`env:FOO`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// some
}
