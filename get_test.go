package get_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/get"
)

func ExampleString_env() {
	os.Setenv(`FOO`, `something`)
	defer os.Unsetenv(`FOO`)

	it, err := get.String(`env:FOO`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// something
}

func ExampleString_env_file() {
	os.Setenv(`FOOFILE`, `testdata/somefile`)
	defer os.Unsetenv(`FOOFILE`)

	it, err := get.String(`env.file:FOOFILE`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// something
}

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
		`env:`, `env.file:`, `env.file.first:`, `env.file.last:`,
		`file:`, `file.first:`, `file.last:`,
		`first:`, `last:`,
		`home:`, `home.first:`, `home.last:`,
		`conf:`, `conf.first:`, `conf.last:`,
		`cache:`, `cache.first:`, `cache.last:`,
		`ssh:`, `ssh.first:`, `ssh.last:`,
		`https:`, `https.first:`, `https.last:`,
		`http:`, `http.first:`, `http.last:`,
		`user@example.com:`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: ""
	// schema: "env.file" value: ""
	// schema: "env.file.first" value: ""
	// schema: "env.file.last" value: ""
	// schema: "file" value: ""
	// schema: "file.first" value: ""
	// schema: "file.last" value: ""
	// schema: "first" value: ""
	// schema: "last" value: ""
	// schema: "home" value: ""
	// schema: "home.first" value: ""
	// schema: "home.last" value: ""
	// schema: "conf" value: ""
	// schema: "conf.first" value: ""
	// schema: "conf.last" value: ""
	// schema: "cache" value: ""
	// schema: "cache.first" value: ""
	// schema: "cache.last" value: ""
	// schema: "ssh" value: ""
	// schema: "ssh.first" value: ""
	// schema: "ssh.last" value: ""
	// schema: "https" value: ""
	// schema: "https.first" value: ""
	// schema: "https.last" value: ""
	// schema: "http" value: ""
	// schema: "http.first" value: ""
	// schema: "http.last" value: ""
	// schema: "user@example.com" value: ""

}

func ExampleSchema_with_Values() {

	valid := []string{
		`env:VALUE`, `env.file:FILE_PATH`, `env.file.first:FILE_PATH`, `env.file.last:FILE_PATH`,
		`file:VALUE`, `file.first:VALUE`, `file.last:VALUE`,
		`first:FILE_PATH`, `last:FILE_PATH`,
		`home:FILE_PATH`, `home.first:FILE_PATH`, `home.last:FILE_PATH`,
		`conf:FILE_PATH`, `conf.first:FILE_PATH`, `conf.last:FILE_PATH`,
		`cache:FILE_PATH`, `cache.first:FILE_PATH`, `cache.last:FILE_PATH`,
		`ssh://user@example.com:1234/some/place`,
		`ssh.first://user@example.com:1234/some/place`,
		`ssh.last://user@example.com:1234/some/place`,
		`https://example.com/some/place`,
		`https.first://example.com/some/place`,
		`https.last://example.com/some/place`,
		`http://example.com/some/place`,
		`http.first://example.com/some/place`,
		`http.last://example.com/some/place`,
		`user@example.com:path/to/file`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: "VALUE"
	// schema: "env.file" value: "FILE_PATH"
	// schema: "env.file.first" value: "FILE_PATH"
	// schema: "env.file.last" value: "FILE_PATH"
	// schema: "file" value: "VALUE"
	// schema: "file.first" value: "VALUE"
	// schema: "file.last" value: "VALUE"
	// schema: "first" value: "FILE_PATH"
	// schema: "last" value: "FILE_PATH"
	// schema: "home" value: "FILE_PATH"
	// schema: "home.first" value: "FILE_PATH"
	// schema: "home.last" value: "FILE_PATH"
	// schema: "conf" value: "FILE_PATH"
	// schema: "conf.first" value: "FILE_PATH"
	// schema: "conf.last" value: "FILE_PATH"
	// schema: "cache" value: "FILE_PATH"
	// schema: "cache.first" value: "FILE_PATH"
	// schema: "cache.last" value: "FILE_PATH"
	// schema: "ssh" value: "//user@example.com:1234/some/place"
	// schema: "ssh.first" value: "//user@example.com:1234/some/place"
	// schema: "ssh.last" value: "//user@example.com:1234/some/place"
	// schema: "https" value: "//example.com/some/place"
	// schema: "https.first" value: "//example.com/some/place"
	// schema: "https.last" value: "//example.com/some/place"
	// schema: "http" value: "//example.com/some/place"
	// schema: "http.first" value: "//example.com/some/place"
	// schema: "http.last" value: "//example.com/some/place"
	// schema: "user@example.com" value: "path/to/file"

}

func ExampleFirstLine() {
	str := "first line\nsecond line\n"
	fmt.Println(get.FirstLine(str))
	// Output:
	// first line
}

func ExampleFirstLineOf() {
	fmt.Println(get.FirstLineOf(`testdata/datafile`))
	// Output:
	// first line <nil>
}

func ExampleLastLine() {
	str := "first line\nsecond line\nlast line\n"
	fmt.Println(get.LastLine(str))
	str = "first line\nsecond line\nlast line"
	fmt.Println(get.LastLine(str))
	// Output:
	// last line
	// last line
}

func ExampleLastLineOf() {
	fmt.Println(get.LastLineOf(`testdata/datafile`))

	// note that a blank last line is perfectly valid
	line, err := get.LastLineOf(`testdata/datafilereturn`)
	fmt.Printf("%q %q", line, err)

	// Output:
	// last line <nil>
	// "" %!q(<nil>)
}
