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

func ExampleString_env_file_first() {
	os.Setenv(`FOOFILE`, `testdata/datafile`)
	defer os.Unsetenv(`FOOFILE`)

	it, err := get.String(`env.file.first:FOOFILE`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// first line
}

func ExampleString_env_file_last() {
	os.Setenv(`FOOFILE`, `testdata/datafile`)
	defer os.Unsetenv(`FOOFILE`)

	it, err := get.String(`env.file.last:FOOFILE`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// last line
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
		`env:`, `env.first:`, `env.last:`,
		`env.file:`, `env.file.first:`, `env.file.last:`,
		`file:`, `file.first:`, `file.last:`,
		`first:`, `last:`,
		`home:`, `home.first:`, `home.last:`,
		`conf:`, `conf.first:`, `conf.last:`,
		`cache:`, `cache.first:`, `cache.last:`,
		`ssh:`, `ssh.first:`, `ssh.last:`,
		`https:`, `https.first:`, `https.last:`,
		`http:`, `http.first:`, `http.last:`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: ""
	// schema: "env.first" value: ""
	// schema: "env.last" value: ""
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

}

func ExampleSchema_with_Values() {

	valid := []string{
		`env:VALUE`, `env.first:VALUE`, `env.last:VALUE`,
		`env.file:FILE_PATH`, `env.file.first:FILE_PATH`, `env.file.last:FILE_PATH`,
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
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: "VALUE"
	// schema: "env.first" value: "VALUE"
	// schema: "env.last" value: "VALUE"
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

}

func ExampleFirstLine() {
	str := "first line\nsecond line\n"
	fmt.Println(get.FirstLine(str))
	// Output:
	// first line
}

func ExampleFirstLineOf() {

	// Contains only:
	//	first line
	//	second line
	//	last line

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

	// Contains only:
	//	first line
	//	second line
	//	last line

	fmt.Println(get.LastLineOf(`testdata/datafile`))

	// A blank last line is perfectly valid.

	// Contains extra blank line:
	//	first line
	//	second line
	//	last line
	//

	line, err := get.LastLineOf(`testdata/datafilereturn`)
	fmt.Printf("%q %q", line, err)

	// Output:
	// last line <nil>
	// "" %!q(<nil>)
}

func ExampleHomeFile() {

	// change home to current directory for testing only
	orig, _ := os.UserHomeDir()
	os.Setenv(`HOME`, `.`)
	defer os.Setenv(`HOME`, orig)

	byt, err := get.HomeFile(`testdata/datafile`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byt))

	// Output:
	// first line
	// second line
	// last line
}

func ExampleSSHOut_short_Form() {

	out, err := get.SSHOut(`localhost`, `echo something`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	/// Output:
	// something
}

func ExampleSSHOut_long_Form() {

	out, err := get.SSHOut(`ssh://rwxrob@localhost:22`, `echo something`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	/// Output:
	// something
}

func ExampleLastLineOfSSH() {

	out, err := get.LastLineOfSSH(`ssh://rwxrob@localhost:22`, `somefile.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	/// Output:
	// last line
}

func ExampleFirstLineOfSSH() {

	out, err := get.FirstLineOfSSH(`ssh://rwxrob@localhost:22`, `somefile.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	/// Output:
	// first line
}

func ExampleRegxSSHURI() {

	parts := get.RegxSSHURI.FindStringSubmatch(`ssh://user@host:22/some/file`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])

	// Output:
	// ssh://user@host:22/some/file
	// user@host:22
	// user
	// host
	// 22
	// /some/file
}

func ExampleRegxSSHURI_nopath() {

	parts := get.RegxSSHURI.FindStringSubmatch(`ssh://user@host:22`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])

	// Output:
	// ssh://user@host:22
	// user@host:22
	// user
	// host
	// 22
}

func ExampleRegxSSHURI_nouser() {

	parts := get.RegxSSHURI.FindStringSubmatch(`ssh://host:22/some/file`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])

	// Output:
	// ssh://host:22/some/file
	// host:22
	//
	// host
	// 22
	// /some/file
}

func ExampleRegxSSHURI_domain_Only() {

	parts := get.RegxSSHURI.FindStringSubmatch(`ssh://host/some/file`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])

	// Output:
	// ssh://host/some/file
	// host
	//
	// host
	//
	// /some/file
}

func ExampleRemoteSCP_random_To() {
	path, err := get.RemoteSCP(`localhost:somefile.txt`, ``)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", path)
	/// Output:
	// /tmp/scp2992279421/
}

func ExampleRemoteSCP_specific_To() {
	path, err := get.RemoteSCP(`localhost:somefile.txt`, `/tmp`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", path)
	/// Output:
	// /tmp
}

func ExampleRemoteSCP_random_To_Many() {
	path, err := get.RemoteSCP(`localhost:*.txt`, ``)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", path)
	/// Output:
	// /tmp/scp*
}

func ExampleFirstFileIn() {
	file, err := get.FirstFileIn(`testdata`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", file)
	/// Output:
	// "datafile"
}
