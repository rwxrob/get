package get

import (
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
	case "env", "env.file", "env.file.head", "env.file.tail":
		return
	case "file", "file.head", "file.tail":
		return
	case "head", "tail":
		return
	case "home", "home.head", "home.tail":
		return
	case "conf", "conf.head", "conf.tail":
		return
	case "cache", "cache.head", "cache.tail":
		return
	case "ssh", "ssh.head", "ssh.tail":
		return
	case "http", "http.head", "http.tail":
		return
	case "https", "https.head", "https.tail":
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

// String returns
func String(url string) (string, error) {
	var it string

	return it, nil
}
