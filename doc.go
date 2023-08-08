// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package get provides simple functions for fetching data (string, []byte)
from common sources in the usual ways enabling the location of such data
to be passed or set by the end user allowing them to choose the best
location for such data through arguments and such.  For example, the
following would both be supported without additional logic to detect the
URL version:

	    foo --token mytoken
			foo --token conf:thisapp/mytokenfile

This enables a user to store their secret token more securely as a file
rather than passing it as an open argument on the command line depending
on their preference.

In all cases, the source provided in the argument signature is a URL of
the normally expected form but with some additional schema/sources
mostly for convenience:

	  (none)    - string as is
		https     - remote HTTP (assumes net/http.DefaultClient)
		http      - (same as https)
		user@host - remote ssh target relative to user home (like git, assumes ssh/scp)
		ssh       - remote ssh target fully qualified (like git, assumes ssh/scp)
		file      - local path, relative or fully qualified
		head      - first line from a file
		tail      - last line from a file
		env       - case-insensitive environment variable name
		env.file  - file path from an environment variable
		env.head  - first line of file from path in environment variable
		env.tail  - last line of file from path in environment variable
		home      - relative path from os.UserHomeDir (for systems without ~)
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

Rather that support encrypted files "at rest" this is left to the caller
to handle since the number of possible encryption methods is simply too
great to support within this package. In such cases an ASCII-armored
text file can be used with get.String and other formats with get.Bytes.

# Caveats

Obviously, get strings that would contain any of the reserved initial
schema identifiers won't work.  See [RegexpSchema] for details. This
expression may change in future releases as additional schemas are
supported. It is therefore suggested to avoid use of this package where
potential conflicts with actual string values might be encountered.
*/
package get
