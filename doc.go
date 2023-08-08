// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package get provides simple functions for fetching data (string, []byte,
and *os.File) from common sources in the usual ways enabling the
location of such data to be passed or set by the end user allowing them
to choose the best location for such data.

In all cases, the source provided in the argument signature is a URL of
the normally expected form but with some additional schema/sources
mostly for convenience:

	https     - remote HTTP (assumes net/http.DefaultClient)
	ssh       - remote ssh target (assumes exec.LookPath("ssh"))
	file      - local path, relative or fully qualified
	head      - first line from a file
	tail      - last line from a file
	env       - case-insensitive environment variable name
	env.file  - file path from an environment variable
	env.head  - first line of file from path in environment variable
	env.tail  - last line of file from path in environment variable
	home      - relative path from os.UserHomeDir
	home.head - first line of file from relative path from os.UserHomeDir
	home.tail - last line of file relative path from os.UserHomeDir
	conf      - relative path from os.UserConfigDir
	conf.head - first line of file from relative path from os.UserConfigDir
	conf.tail - last line of file relative path from os.UserConfigDir
	cache     - relative path from os.UserCacheDir
	cache.head - first line of file from relative path from os.UserCacheDir
	cache.tail - last line of file relative path from os.UserCacheDir

Most may prefer to cut and paste from this package and "vendor" the code
into their own to keep dependencies to a minimum. (Please maintain some
form of the copyright acknowledgement.)

# Security

Remote fetch methods always require some form of encryption (and always
will). This means that unencrypted HTTP Get and FTP will never be
supported.

Rather that support encrypted files "at rest" this is left to the caller
to handle since the number of possible encryption methods is simply too
great to support within this package. In such cases an ASCII-armored
text file can be used with get.String and other formats with get.Bytes.
*/
package get
