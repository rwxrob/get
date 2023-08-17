package expr_test

import (
	"fmt"

	"github.com/rwxrob/get/expr"
)

func ExampleSSHURI() {

	parts := expr.SSHURI.FindStringSubmatch(`ssh://user@host:22/some/file`)
	if parts == nil {
		fmt.Println(`failed to parse`)
		return
	}

	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])
	fmt.Println(parts[6])

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

	parts := expr.SSHURI.FindStringSubmatch(`ssh://user@host:22`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])
	fmt.Println(parts[6])

	// Output:
	// ssh://user@host:22
	// ssh
	// user@host:22
	// user
	// host
	// 22
}

func ExampleSSHURI_nouser() {

	parts := expr.SSHURI.FindStringSubmatch(`ssh://host:22/some/file`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])
	fmt.Println(parts[6])

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

	parts := expr.SSHURI.FindStringSubmatch(`ssh://host/some/file`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])
	fmt.Println(parts[6])

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

	parts := expr.SSHURI.FindStringSubmatch(`ssh://host`)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println(parts[2])
	fmt.Println(parts[3])
	fmt.Println(parts[4])
	fmt.Println(parts[5])
	fmt.Println(parts[6])

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

	parts := expr.SSHURI.FindStringSubmatch(`ssh://user@:22`)
	fmt.Println(parts)

	// Output:
	// []
}
