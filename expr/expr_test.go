package expr_test

import (
	"fmt"

	"github.com/rwxrob/get/expr"
)

func ExampleSSHURI() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://user@host:22/some/file`)
	if p == nil {
		fmt.Println(`failed to parse`)
		return
	}

	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])
	fmt.Println(p[4])
	fmt.Println(p[5])
	fmt.Println(p[6])

	// Output:
	// ssh://user@host:22/some/file
	// ssh
	// user@host:22
	// user
	// host
	// 22
	// some/file
}

func ExampleSSHURI_nopath() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://user@host:22`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])
	fmt.Println(p[4])
	fmt.Println(p[5])
	fmt.Println(p[6])

	// Output:
	// ssh://user@host:22
	// ssh
	// user@host:22
	// user
	// host
	// 22
}

func ExampleSSHURI_nouser() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://host:22/some/file`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])
	fmt.Println(p[4])
	fmt.Println(p[5])
	fmt.Println(p[6])

	// Output:
	// ssh://host:22/some/file
	// ssh
	// host:22
	//
	// host
	// 22
	// some/file
}

func ExampleSSHURI_domain_Only() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://host/some/file`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])
	fmt.Println(p[4])
	fmt.Println(p[5])
	fmt.Println(p[6])

	// Output:
	// ssh://host/some/file
	// ssh
	// host
	//
	// host
	//
	// some/file
}

func ExampleSSHURI_just_Host() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://host`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])
	fmt.Println(p[4])
	fmt.Println(p[5])
	fmt.Println(p[6])

	// Output:
	// ssh://host
	// ssh
	// host
	//
	// host
	//
	//
}

func ExampleSSHURI_invalid_No_Host() {

	p := expr.SSHURI.FindStringSubmatch(`ssh://user@:22`)
	fmt.Println(p)

	// Output:
	// []
}

func ExampleSSHURIShort() {

	p := expr.SSHURIShort.FindStringSubmatch(`user@host:some/file`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])

	// Output:
	// user@host:some/file
	// user
	// host
	// some/file
}

func ExampleSSHURIShort_no_Path() {

	p := expr.SSHURIShort.FindStringSubmatch(`user@host`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])

	// Output:
	// user@host
	// user
	// host
	//
}

func ExampleSSHURIShort_no_User() {

	p := expr.SSHURIShort.FindStringSubmatch(`host`)
	fmt.Println(p[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println(p[3])

	// Output:
	// host
	//
	// host
	//
}

func ExampleSSHURIShort_invalid_No_Host() {

	p := expr.SSHURIShort.FindStringSubmatch(`user@:some/file`)
	fmt.Println(p)

	// Output:
	// []
}
