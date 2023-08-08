package get_test

import (
	"fmt"

	"github.com/rwxrob/get"
)

/*
func ExampleString_env() {

	it, err := get.String(`env:FOO`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// some
}
*/

func ExampleSchema() {

	valid := []string{
		`https://withtls`,
		`http://notls`,
		`user@server.com:file`,
		`ssh://user@server.com:12345/~/file`,
		`file:/some/path/to/file`,
		`head:/some/path/to/file`,
		`tail:/some/path/to/file`,
		`env:ENV_VAR`,
		`env.file:ENV_WITH_PATH`,
		`env.head:/some/path/to/file`,
		`env.tail:/some/path/to/file`,
		`home:file`,
		`home.head:file`,
		`home.tail:file`,
		`conf:name/file`,
		`conf.head:name/file`,
		`conf.tail:name/file`,
		`cache:name/file`,
		`cache.head:name/file`,
		`cache.tail:name/file`,
	}

	for _, it := range valid {
		fmt.Println(get.Schema(it))
	}

	// Output:
	// https:
	// http:
	// user@server.com:
	// ssh:
	// file:
	// head:
	// tail:
	// env:
	// env.file:
	// env.head:
	// env.tail:
	// home:
	// home.head:
	// home.tail:
	// conf:
	// conf.head:
	// conf.tail:
	// cache:
	// cache.head:
	// cache.tail:
}
