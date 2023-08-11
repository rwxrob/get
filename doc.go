// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package get provides simple functions for fetching data ([String], [Bytes])
from common sources in the usual ways enabling the location of such data
to be passed or set by the end user allowing the user to decide the best
location for such data through arguments and such.  For example, the
following would both be supported without additional logic to detect the
URL version:

	foo --token mytoken
	foo --token conf:thisapp/mytokenfile
	foo --token env.file.first:TOKEN_FILE_PATH

This enables a user to store their secret token more securely as a file
rather than passing it as an open argument on the command line depending
on their preference.

For more details see the [String] function.

# Consider "vendoring"

Most may prefer to cut and paste from this package and "vendor" the code
into their own to keep dependencies to a minimum. If so, please maintain
some form of the copyright acknowledgement.
*/
package get
