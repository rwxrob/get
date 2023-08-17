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

func ExampleString_env_file_head() {
	os.Setenv(`FOOFILE`, `testdata/datafile`)
	defer os.Unsetenv(`FOOFILE`)

	it, err := get.String(`env.file.head:FOOFILE`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(it)

	// Output:
	// first line
}

func ExampleString_env_file_tail() {
	os.Setenv(`FOOFILE`, `testdata/datafile`)
	defer os.Unsetenv(`FOOFILE`)

	it, err := get.String(`env.file.tail:FOOFILE`)
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
		`env:`, `env.head:`, `env.tail:`,
		`env.file:`, `env.file.head:`, `env.file.tail:`,
		`file:`, `file.head:`, `file.tail:`,
		`head:`, `tail:`,
		`home:`, `home.head:`, `home.tail:`,
		`conf:`, `conf.head:`, `conf.tail:`,
		`cache:`, `cache.head:`, `cache.tail:`,
		`scp:`, `ssh.head:`, `ssh.tail:`,
		`https:`, `https.head:`, `https.tail:`,
		`http:`, `http.head:`, `http.tail:`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: ""
	// schema: "env.head" value: ""
	// schema: "env.tail" value: ""
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
	// schema: "scp" value: ""
	// schema: "ssh.head" value: ""
	// schema: "ssh.tail" value: ""
	// schema: "https" value: ""
	// schema: "https.head" value: ""
	// schema: "https.tail" value: ""
	// schema: "http" value: ""
	// schema: "http.head" value: ""
	// schema: "http.tail" value: ""

}

func ExampleSchema_with_Values() {

	valid := []string{
		`env:VALUE`, `env.head:VALUE`, `env.tail:VALUE`,
		`env.file:FILE_PATH`, `env.file.head:FILE_PATH`, `env.file.tail:FILE_PATH`,
		`file:VALUE`, `file.head:VALUE`, `file.tail:VALUE`,
		`head:FILE_PATH`, `tail:FILE_PATH`,
		`home:FILE_PATH`, `home.head:FILE_PATH`, `home.tail:FILE_PATH`,
		`conf:FILE_PATH`, `conf.head:FILE_PATH`, `conf.tail:FILE_PATH`,
		`cache:FILE_PATH`, `cache.head:FILE_PATH`, `cache.tail:FILE_PATH`,
		`scp://user@example.com:1234/some/place`,
		`ssh.head://user@example.com:1234/some/place`,
		`ssh.tail://user@example.com:1234/some/place`,
		`https://example.com/some/place`,
		`https.head://example.com/some/place`,
		`https.tail://example.com/some/place`,
		`http://example.com/some/place`,
		`http.head://example.com/some/place`,
		`http.tail://example.com/some/place`,
	}

	for _, it := range valid {
		schema, value := get.Schema(it)
		fmt.Printf("schema: %q value: %q\n", schema, value)
	}

	// Output:
	// schema: "env" value: "VALUE"
	// schema: "env.head" value: "VALUE"
	// schema: "env.tail" value: "VALUE"
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
	// schema: "scp" value: "//user@example.com:1234/some/place"
	// schema: "ssh.head" value: "//user@example.com:1234/some/place"
	// schema: "ssh.tail" value: "//user@example.com:1234/some/place"
	// schema: "https" value: "//example.com/some/place"
	// schema: "https.head" value: "//example.com/some/place"
	// schema: "https.tail" value: "//example.com/some/place"
	// schema: "http" value: "//example.com/some/place"
	// schema: "http.head" value: "//example.com/some/place"
	// schema: "http.tail" value: "//example.com/some/place"

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

	out, err := get.FirstLineOfSSH(`ssh://rwxrob@localhost:22/somefile.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	/// Output:
	// first line
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

func ExampleSSHURI_UpdateAddr() {

	uri := &get.SSHURI{
		Schema: `ssh`,
		Host:   `host`,
		User:   `user`,
		Port:   `22`,
	}

	uri.UpdateAddr()
	fmt.Println(uri.Addr)

	uri.Port = ``
	uri.UpdateAddr()
	fmt.Println(uri.Addr)

	uri.User = ``
	uri.UpdateAddr()
	fmt.Println(uri.Addr)

	uri.Host = ``
	uri.UpdateAddr()
	fmt.Println(uri.Addr)

	// Output:
	// user@host:22
	// user@host
	// host
	//

}

func ExampleParseSSHURI() {

	// canonical but not that some/path is path, not /some/path
	uri := get.ParseSSHURI(`ssh://user@host:22/some/path`)
	fmt.Println(uri)
	fmt.Println(uri.Path)

	// head and tail schema extensions supported
	uri = get.ParseSSHURI(`ssh.head://user@host:22/some/path`)
	fmt.Println(uri)

	// full path
	uri = get.ParseSSHURI(`ssh://user@host:22//etc/passwd`)
	fmt.Println(uri)

	// no path
	uri = get.ParseSSHURI(`ssh://user@host:22`)
	fmt.Println(uri)

	// no port
	uri = get.ParseSSHURI(`ssh://user@host`)
	fmt.Println(uri)

	// no user
	uri = get.ParseSSHURI(`ssh://host`)
	fmt.Println(uri)

	// short form, all
	uri = get.ParseSSHURI(`user@host:some/path`)
	fmt.Println(uri)

	// short form, full path
	uri = get.ParseSSHURI(`user@host:/etc/passwd`)
	fmt.Println(uri)

	// short form, no path
	uri = get.ParseSSHURI(`user@host`)
	fmt.Println(uri)

	// short form, no user
	uri = get.ParseSSHURI(`host`)
	fmt.Println(uri)

	// invalid
	uri = get.ParseSSHURI(`bogus://user@:22/some/path`)
	fmt.Println(uri)
	uri = get.ParseSSHURI(`ssh.head://user@:22/some/path`)
	fmt.Println(uri)
	uri = get.ParseSSHURI(`user@host:`)
	fmt.Println(uri)
	uri = get.ParseSSHURI(`user@:/some/path`)
	fmt.Println(uri)
	uri = get.ParseSSHURI(`@host`)
	fmt.Println(uri)

	// Output:
	// ssh://user@host:22/some/path
	// some/path
	// ssh.head://user@host:22/some/path
	// ssh://user@host:22//etc/passwd
	// ssh://user@host:22
	// ssh://user@host
	// ssh://host
	// ssh://user@host:22/some/path
	// ssh://user@host:22//etc/passwd
	// ssh://user@host:22
	// ssh://host:22
	// <nil>
	// <nil>
	// <nil>
	// <nil>
	// <nil>

}
