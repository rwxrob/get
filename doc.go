// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

/*
Package get provides simple functions for fetching data (string, []byte)
from common sources in the usual ways enabling the location of such data
to be passed or set by the end user allowing the user to decide the best
location for such data through arguments and such.  For example, the
following would both be supported without additional logic to detect the
URL version:

	    foo --token mytoken
			foo --token conf:thisapp/mytokenfile
			foo --token env.file.head:TOKEN_FILE_PATH

This enables a user to store their secret token more securely as a file
rather than passing it as an open argument on the command line depending
on their preference.

The detection of a special URL source string is done by identifying any
of the reserved schema type combinations up to the first colon.
Therefore, use of this package where colons might be value string values
should be used with caution. See [RegexpSchema] for exact details.

	  (none)    - string as is
		env       - value of environment variable by name
		file      - full content of local file at path
		head      - first line from a local file at path (no line endings)
		tail      - last line from a local file at path (no line endings)
		home      - full content of local file relative to os.UserHomeDir (for systems without ~)
		conf      - full content of local file relative os.UserConfigDir
		cache     - full content of local file relative os.UserCacheDir
		user@host - full content of remote file over ssh (like git, assumes scp)
		ssh       - full content of remote file over ssh (like git, assumes scp)
		https     - full content of remote HTTP/TLS GET (net/http.DefaultClient)
		http      - full content of remote HTTP GET (net/http.DefaultClient)

All of these methods (except the user@host ssh shortcut which
technically doesn't qualify as a URL) can be combined using a simple
dotted notation pipeline. In such cases, the string resulting from each
schema type is used as the value for the next. For example,
env.file.head:TOKEN_FILE would first lookup the value of the environment
variable TOKEN_FILE and then use that as the value for file. So if
TOKEN_FILE environment variable contained ~/.mytoken then file:
~/.mytoken would be evaluated and only the first chomped line (head)
returned.

In all cases, the source provided in the argument signature is a URL of
the normally expected form but with some additional schema/sources
mostly for convenience (ordered by simplest to most complex)

# Line endings

Note that except for head and tail the line endings are always preserved
if they are included. The caller must remove these if needed.

# Enabling comments in files

By using head and tail comments of any kind may be included in the data
itself. For example, the command for how to retrieve a token might be
included on the line after the first line that contains the token itself;
or, a strong warning never to divulge a secret could be created; or,
a comment noting that a particular file is generated automatically from
some other automated process (cronjob, etc.)

# Security

Rather that support encrypted files "at rest" this is left to the caller
to handle since the number of possible encryption methods is simply too
great to support within this package. In such cases an ASCII-armored
text file can be used with get.String and other formats with get.Bytes.

# Design considerations

Parsing of the data is isolated to the first (head) or last line (tail)
and always will be. Additional parsing and/or unmarshalling is
appropriately left to the caller.

# Consider "vendoring"

Most may prefer to cut and paste from this package and "vendor" the code
into their own to keep dependencies to a minimum. If so, please maintain
some form of the copyright acknowledgement.
*/
package get
