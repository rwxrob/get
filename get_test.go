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

func ExampleSchema_values_Only() {

	valid := []string{
		`just a string`,
		`:just a string`,
		`not:just a string`,
		``,
		`234`,
		`really:not:a:schema`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "" value: "just a string"
	// schema: "" value: ":just a string"
	// schema: "" value: "not:just a string"
	// schema: "" value: ""
	// schema: "" value: "234"
	// schema: "" value: "really:not:a:schema"

}
func ExampleSchema_partial() {

	valid := []string{
		`env:`, `env.file:`, `env.file.head:`, `env.file.tail:`,
		`file:`, `file.head:`, `file.tail:`,
		`head:`, `tail:`,
		`home:`, `home.head:`, `home.tail:`,
		`conf:`, `conf.head:`, `conf.tail:`,
		`cache:`, `cache.head:`, `cache.tail:`,
		`ssh:`, `ssh.head:`, `ssh.tail:`,
		`https:`, `https.head:`, `https.tail:`,
		`http:`, `http.head:`, `http.tail:`,
		`user@example.com:`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: ""
	// schema: "env.file" value: ""
	// schema: "env.file.head" value: ""
	// schema: "env.file.tail" value: ""
	// schema: "file" value: ""
	// schema: "file.head" value: ""
	// schema: "file.tail" value: ""
	// schema: "head" value: ""
	// schema: "tail" value: ""
	// schema: "home" value: ""
	// schema: "home.head" value: ""
	// schema: "home.tail" value: ""
	// schema: "conf" value: ""
	// schema: "conf.head" value: ""
	// schema: "conf.tail" value: ""
	// schema: "cache" value: ""
	// schema: "cache.head" value: ""
	// schema: "cache.tail" value: ""
	// schema: "ssh" value: ""
	// schema: "ssh.head" value: ""
	// schema: "ssh.tail" value: ""
	// schema: "https" value: ""
	// schema: "https.head" value: ""
	// schema: "https.tail" value: ""
	// schema: "http" value: ""
	// schema: "http.head" value: ""
	// schema: "http.tail" value: ""
	// schema: "user@example.com" value: ""

}

func ExampleSchema_with_Values() {

	valid := []string{
		`env:VALUE`, `env.file:FILE_PATH`, `env.file.head:FILE_PATH`, `env.file.tail:FILE_PATH`,
		`file:VALUE`, `file.head:VALUE`, `file.tail:VALUE`,
		`head:FILE_PATH`, `tail:FILE_PATH`,
		`home:FILE_PATH`, `home.head:FILE_PATH`, `home.tail:FILE_PATH`,
		`conf:FILE_PATH`, `conf.head:FILE_PATH`, `conf.tail:FILE_PATH`,
		`cache:FILE_PATH`, `cache.head:FILE_PATH`, `cache.tail:FILE_PATH`,
		`ssh://user@example.com:1234/some/place`,
		`ssh.head://user@example.com:1234/some/place`,
		`ssh.tail://user@example.com:1234/some/place`,
		`https://example.com/some/place`,
		`https.head://example.com/some/place`,
		`https.tail://example.com/some/place`,
		`http://example.com/some/place`,
		`http.head://example.com/some/place`,
		`http.tail://example.com/some/place`,
		`user@example.com:path/to/file`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: "VALUE"
	// schema: "env.file" value: "FILE_PATH"
	// schema: "env.file.head" value: "FILE_PATH"
	// schema: "env.file.tail" value: "FILE_PATH"
	// schema: "file" value: "VALUE"
	// schema: "file.head" value: "VALUE"
	// schema: "file.tail" value: "VALUE"
	// schema: "head" value: "FILE_PATH"
	// schema: "tail" value: "FILE_PATH"
	// schema: "home" value: "FILE_PATH"
	// schema: "home.head" value: "FILE_PATH"
	// schema: "home.tail" value: "FILE_PATH"
	// schema: "conf" value: "FILE_PATH"
	// schema: "conf.head" value: "FILE_PATH"
	// schema: "conf.tail" value: "FILE_PATH"
	// schema: "cache" value: "FILE_PATH"
	// schema: "cache.head" value: "FILE_PATH"
	// schema: "cache.tail" value: "FILE_PATH"
	// schema: "ssh" value: "//user@example.com:1234/some/place"
	// schema: "ssh.head" value: "//user@example.com:1234/some/place"
	// schema: "ssh.tail" value: "//user@example.com:1234/some/place"
	// schema: "https" value: "//example.com/some/place"
	// schema: "https.head" value: "//example.com/some/place"
	// schema: "https.tail" value: "//example.com/some/place"
	// schema: "http" value: "//example.com/some/place"
	// schema: "http.head" value: "//example.com/some/place"
	// schema: "http.tail" value: "//example.com/some/place"
	// schema: "user@example.com" value: "path/to/file"

}
