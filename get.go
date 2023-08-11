package get

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var SSHShortForm = regexp.MustCompile(`[A-Za-z0-9]@[A-Za-z0-9.]`)

// Schema returns the schema up to the first colon if found. Only valid
// schema combinations will be returned (see source switch for all
// possible schemas). All others return an empty string for the schema.
// The second string is always the remaining value (after the colon).
func Schema(a string) (schema, value string) {
	parts := strings.SplitN(a, `:`, 2)
	count := len(parts)

	switch count {
	case 2:
		schema = parts[0]
		value = parts[1]
	case 1:
		value = parts[0]
		return
	default:
		value = a
		return
	}

	switch schema {
	case "env", "env.first", "env.last":
		return
	case "env.file", "env.file.first", "env.file.last":
		return
	case "file", "file.first", "file.last":
		return
	case "first", "last":
		return
	case "home", "home.first", "home.last":
		return
	case "conf", "conf.first", "conf.last":
		return
	case "cache", "cache.first", "cache.last":
		return
	case "ssh", "ssh.first", "ssh.last":
		return
	case "http", "http.first", "http.last":
		return
	case "https", "https.first", "https.last":
		return
	}

	if SSHShortForm.MatchString(schema) {
		return
	}

	// looks like we just have a plain string
	schema = ""
	value = a

	return

}

// String returns it's string argument unless one of the following special URL
// schemas is detected at the beginning of that string in which case the
// string is considered a URL and the string returned is derived from
// that URL providing a mechanism for users to decide where best to
// persist string values, for example secrets in files or on secured,
// remote https or ssh locations :
//
//	(none)    - string as is
//	env       - value of environment variable by name
//	file      - full content of local file at path
//	first     - first line from a local file at path (no line endings)
//	last      - last line from a local file at path (no line endings)
//	home      - full content of local file relative to os.UserHomeDir (for systems without ~)
//	conf      - full content of local file relative os.UserConfigDir
//	cache     - full content of local file relative os.UserCacheDir
//	user@host - full content of remote file over ssh (like git, assumes scp)
//	ssh       - full content of remote file over ssh (like git, assumes scp)
//	https     - full content of remote HTTP/TLS GET (net/http.DefaultClient)
//	http      - full content of remote HTTP GET (net/http.DefaultClient)
//
// In all cases, the source provided in the argument signature is a URL
// of the normally expected form but with some additional schema/sources
// mostly for convenience (ordered by simplest to most complex)
//
// The detection of a special URL source string is done by identifying
// any of the reserved schema type combinations up to the first colon.
// Therefore, use of this package where colons might be value string
// values should be used with caution.
//
// All of these methods (except the user@host ssh shortcut which
// technically doesn't qualify as a URL) can be combined using a simple
// dotted notation pipeline. In such cases, the string resulting from
// each schema type is used as the value for the next. For example,
// env.file.first:TOKEN_FILE would first lookup the value of the
// environment variable TOKEN_FILE and then use that as the value for
// file. So if TOKEN_FILE environment variable contained ~/.mytoken then
// file: ~/.mytoken would be evaluated and only the first chomped line
// returned.
//
// # Line endings
//
// Note that except for first and last the line endings are always
// preserved if they are included. The caller must remove these if
// needed.
//
// # Enabling comments in files
//
// By using first and last comments of any kind may be included in the
// data itself. For example, the command for how to retrieve a token
// might be included on the line after the first line that contains the
// token itself; or, a strong warning never to divulge a secret could be
// created; or, a comment noting that a particular file is generated
// automatically from some other automated process (cronjob, etc.)
//
// # Security
//
// Rather that support encrypted files "at rest" this is left to the
// caller to handle since the number of possible encryption methods is
// simply too great to support within this package. In such cases an
// ASCII-armored text file can be used with get.String and other formats
// with get.Bytes.
//
// # Design considerations
//
// Parsing of the data is isolated to the first or last line
// (last) and always will be. Additional parsing and/or unmarshalling is
// appropriately left to the caller.
func String(a string) (string, error) {
	var it string
	schema, value := Schema(a)

	if len(schema) == 0 {
		return value, nil
	}

	switch schema {
	case `env`:
		return os.Getenv(value), nil
	case `env.file`:
		file := os.Getenv(value)
		byt, err := os.ReadFile(file)
		if err != nil {
			return string(byt), err
		}
		return string(byt), nil
	case `env.first`:
		// TODO
	case `env.last`:
		// TODO
	}

	return it, nil
}

// FirstLine returns the first line of the string or []byte.
func FirstLine[T string | []byte](a T) string {
	r := strings.NewReader(string(a))
	s := bufio.NewScanner(r)
	s.Scan()
	return s.Text()
}

// FirstLineOf returns the first line (ending in \r?\n) of file at path
// without buffering the entire file.
func FirstLineOf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	s := bufio.NewScanner(f)
	s.Scan()
	return s.Text(), nil
}

// LastLine returns the last line of the string or []byte without
// buffering anything more than the last line read.
func LastLine[T string | []byte](a T) string {
	r := strings.NewReader(string(a))
	s := bufio.NewScanner(r)
	var prev string
	for s.Scan() {
		prev = s.Text()
	}
	return prev
}

// LastLineOf returns the last line (ending in \r?\n) of file at path
// without buffering the entire file.
func LastLineOf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	s := bufio.NewScanner(f)
	var prev string
	for s.Scan() {
		prev = s.Text()
	}
	return prev, nil
}
