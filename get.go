package get

import (
	"bufio"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var RegxSSHURI = regexp.MustCompile(`ssh://((?:([A-Za-z0-9]+)@)?([A-Za-z0-9.]+)(?::([0-9]{1,7}))?)(\S+)?`)

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
	case `env`, `env.first`, `env.last`:
		return
	case `env.file`, `env.file.first`, `env.file.last`:
		return
	case `file`, `file.first`, `file.last`:
		return
	case `first`, `last`:
		return
	case `home`, `home.first`, `home.last`:
		return
	case `conf`, `conf.first`, `conf.last`:
		return
	case `cache`, `cache.first`, `cache.last`:
		return
	case `ssh`, `ssh.first`, `ssh.last`:
		return
	case `http`, `http.first`, `http.last`:
		return
	case `https`, `https.first`, `https.last`:
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
//	(none)         - string as is
//	env            - value of environment variable by name
//	env.first      - file line of value of environment variable by name
//	env.last       - last line of value of environment variable by name
//	env.file       - full content file at path from environment variable
//	env.file.first - first line of file at path from environment variable
//	env.file.last  - last line of file at path from environment variable
//	file           - full content of local file at path
//	file.first     - first line of file
//	file.last      - last line of file
//	first          - (same as file.first)
//	last           - (same as file.last)
//	home           - full content of local file relative to os.UserHomeDir (for systems without ~)
//	home.first     - first line of home
//	home.last      - last line of home
//	conf           - full content of local file relative os.UserConfigDir
//	conf.first     - first line of conf
//	conf.last      - last line of conf
//	cache          - full content of local file relative os.UserCacheDir
//	cache.first    - first line of cache
//	cache.last     - last line of cache
//	ssh            - full content of remote file over ssh (like scp)
//	ssh.first      - first line of ssh (with head)
//	ssh.last       - last line of ssh (with tail)
//	http(s)        - full content of remote HTTP/TLS GET (net/http.DefaultClient)
//	http(s).first  - first line of http(s) (from full GET)
//	http(s).last   - last line of https(s) (from full GET)
//
// For more information about how the data is acquired and parsed see
// the relevant helper functions ([HomeFile], [CacheFile], [ConfFile]
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
// # External dependencies
//
// The ssh schemas require ssh and scp to be installed and available
// in the PATH for the host system. Most all systems that have ssh
// installed (openssh, for example) automatically have scp installed as
// well.
func String(a string) (string, error) {
	var it string
	schema, value := Schema(a)

	if len(schema) == 0 {
		return value, nil
	}

	switch schema {

	case `env`:
		return os.Getenv(value), nil
	case `env.first`:
		return FirstLine(os.Getenv(value)), nil
	case `env.last`:
		return LastLine(os.Getenv(value)), nil

	case `env.file`:
		byt, err := os.ReadFile(os.Getenv(value))
		if err != nil {
			return "", err
		}
		return string(byt), nil
	case `env.file.first`:
		return FirstLineOf(os.Getenv(value))
	case `env.file.last`:
		return LastLineOf(os.Getenv(value))

	case `file`:
		byt, err := os.ReadFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	case `file.first`, `first`:
		return FirstLineOf(value)
	case `file.last`, `last`:
		return LastLineOf(value)

	case `home`:
		byt, err := HomeFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	case `home.first`:
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path := path.Join(home, value)
		return FirstLineOf(path)
	case `home.last`:
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path := path.Join(home, value)
		return LastLineOf(path)

	case `conf`:
		byt, err := ConfFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	case `conf.first`:
		conf, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		path := path.Join(conf, value)
		return FirstLineOf(path)
	case `conf.last`:
		conf, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		path := path.Join(conf, value)
		return LastLineOf(path)

	case `cache`:
		byt, err := CacheFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil
	case `cache.first`:
		cache, err := os.UserCacheDir()
		if err != nil {
			return "", err
		}
		path := path.Join(cache, value)
		return FirstLineOf(path)
	case `cache.last`:
		cache, err := os.UserCacheDir()
		if err != nil {
			return "", err
		}
		path := path.Join(cache, value)
		return LastLineOf(path)

	case `ssh`:
	case `ssh.first`:
		//TODO
		//return FirstLineOfSSH(target, path)
	case `ssh.last`:

	case `http`, `https`:
	case `http.first`, `https.first`:
	case `http.last`, `https.last`:

	}
	return it, nil
}

// HomeFile returns the []byte content of a file within the
// os.UserHomeDir.
func HomeFile(relpath string) ([]byte, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path.Join(dir, relpath))
}

// CacheFile returns the []byte content of a file within the
// os.UserCacheDir.
func CacheFile(relpath string) ([]byte, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path.Join(dir, relpath))
}

// ConfFile returns the []byte content of a file within the
// os.UserConfigDir.
func ConfFile(relpath string) ([]byte, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	return os.ReadFile(path.Join(dir, relpath))
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

// RemoteFileSCPTemp returns an file opened for reading after having
// copied it with scp to a temporary file and opened it there. No
// attempt to remove the copied temporary file is done leaving this task
// to the caller.

//RemoteFileSCP returns the entire content of the remote file by first
//copying it with scp into a temporary local file and then buffering the
//entire file into the []byte slice returned.

// FirstLineOfSSH returns only the first line of a remote file by
// calling head on the file over an ssh connection. Otherwise, identical
// to LastLineOfSSH.
func FirstLineOfSSH(target, path string) (string, error) {
	return SSHOut(target, `head -1 `+path)
}

// LastLineOfSSH returns the last line of a remote file by calling tail
// on the file at the path indicated making it safe for grabbing
// exactly one last line of a large remote file (log, etc.).
//
// If the remote system does not support the tail command returns an error
// stating as much. See SSHOut for valid target formats. The path can be
// relative to the login home directory or fully qualified (beginning
// with slash).
func LastLineOfSSH(target, path string) (string, error) {
	return SSHOut(target, `tail -1 `+path)
}

// SSHOut sends the command string to the target using the ssh command
// on the host system (not the ssh package) and returns the standard
// output. It is the equivalent of the following command line:
//
//	ssh '<target>' '<command>'
//
// Note that the <target> may be either a relative ssh shortcut (ex:
// user@localhost) or a fully qualified ssh URI (ex: ssh:
// //user@localhost:22). See the documentation on the ssh command itself
// for more details.
func SSHOut(target, command string) (string, error) {
	sshexe, err := exec.LookPath(`ssh`)
	if err != nil {
		return "", err
	}
	byt, err := exec.Command(sshexe, target, command).Output()
	return string(byt), err
}
