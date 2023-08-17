package get

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/rwxrob/get/expr"
)

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
	case `env`, `env.head`, `env.tail`:
		return
	case `env.file`, `env.file.head`, `env.file.tail`:
		return
	case `file`, `file.head`, `file.tail`:
		return
	case `head`, `tail`:
		return
	case `home`, `home.head`, `home.tail`:
		return
	case `conf`, `conf.head`, `conf.tail`:
		return
	case `cache`, `cache.head`, `cache.tail`:
		return
	case `scp`:
		return
	case `ssh.head`, `ssh.tail`:
		return
	case `http`, `http.head`, `http.tail`:
		return
	case `https`, `https.head`, `https.tail`:
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
//	env.head      - file line of value of environment variable by name
//	env.tail       - tail line of value of environment variable by name
//	env.file       - full content file at path from environment variable
//	env.file.head - head line of file at path from environment variable
//	env.file.tail  - tail line of file at path from environment variable
//	file           - full content of local file at path
//	file.head     - head line of file
//	file.tail      - tail line of file
//	head          - (same as file.head)
//	tail           - (same as file.tail)
//	home           - full content of local file relative to os.UserHomeDir (for systems without ~)
//	home.head     - head line of home
//	home.tail      - tail line of home
//	conf           - full content of local file relative os.UserConfigDir
//	conf.head     - head line of conf
//	conf.tail      - tail line of conf
//	cache          - full content of local file relative os.UserCacheDir
//	cache.head    - head line of cache
//	cache.tail     - tail line of cache
//	scp            - full content of remote file over scp
//	ssh.head      - head line of remote file with ssh head -1
//	ssh.tail       - tail line of remote file with ssh tail -1
//	http(s)        - full content of remote HTTP/TLS GET (net/http.DefaultClient)
//	http(s).head  - head line of http(s) (from full GET)
//	http(s).tail   - tail line of https(s) (from full GET)
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
// Note that except for head and tail the line endings are always
// preserved if they are included. The caller must remove these if
// needed.
//
// # Enabling comments in files
//
// By using head and tail comments of any kind may be included in the
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
func String(target string) (string, error) {
	schema, value := Schema(target)

	// not a reserved schema, must just be a string
	if len(schema) == 0 {
		return target, nil
	}

	switch schema {

	case `env`:
		return os.Getenv(value), nil

	case `env.head`:
		return FirstLine(os.Getenv(value)), nil

	case `env.tail`:
		return LastLine(os.Getenv(value)), nil

	case `env.file`:
		byt, err := os.ReadFile(os.Getenv(value))
		if err != nil {
			return "", err
		}
		return string(byt), nil

	case `env.file.head`:
		return FirstLineOf(os.Getenv(value))

	case `env.file.tail`:
		return LastLineOf(os.Getenv(value))

	case `file`:
		byt, err := os.ReadFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil

	case `file.head`, `head`:
		return FirstLineOf(value)

	case `file.tail`, `tail`:
		return LastLineOf(value)

	case `home`:
		byt, err := HomeFile(value)
		if err != nil {
			return "", err
		}
		return string(byt), nil

	case `home.head`:
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path := path.Join(home, value)
		return FirstLineOf(path)

	case `home.tail`:
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

	case `conf.head`:
		conf, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		path := path.Join(conf, value)
		return FirstLineOf(path)

	case `conf.tail`:
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

	case `cache.head`:
		cache, err := os.UserCacheDir()
		if err != nil {
			return "", err
		}
		path := path.Join(cache, value)
		return FirstLineOf(path)

	case `cache.tail`:
		cache, err := os.UserCacheDir()
		if err != nil {
			return "", err
		}
		path := path.Join(cache, value)
		return LastLineOf(path)

	case `scp`:

	case `ssh.head`:

		// if value begins with // assume a full SSH URL
		// (ssh.head://localhost/somefile.txt)
		// (ssh.head://localhost//full/path/to/somefile.txt)
		if strings.HasPrefix(value, `//`) {
			//parts := strings.SplitN(value[2:], `/`, 2)
			//if len(parts) < 2 {
			//return ``, fmt.Errorf(`missing file/path`)
			//}
			//target= `ssh://`+parts[0]+
			target = `ssh:` + value
		}

		// otherwise, assume shortform
		//   (ssh.head:localhost:somefile.txt)

		//return FirstLineOfSSH(target)

	case `ssh.tail`:

	case `http`, `https`:

	case `http.head`, `https.head`:

	case `http.tail`, `https.tail`:

	}

	// should never get here, but whatever
	return target, nil

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

// FirstLine returns the first line of the string or []byte (similar to
// the head -1 command).
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

// RemoteSCP copies one or more remote files from the remote target into
// the directory specified (to) (or a temporary directory if to is blank)
// using the scp command and returns the path to the directory. The path
// to the directory is always returned (even when there is an error).
// The to directory is never removed leaving this to the caller if
// needed. The target (from) may be any valid scp short form (user@host:
// rel/path/to/file) or long form URI (scp://user@host:22//full/path/to/file).
// Note that the scp command is executed with -r added to allow for
// recursive directory copies. Returns an error if the scp command
// cannot be found.
func RemoteSCP(from, to string) (string, error) {

	scpexe, err := exec.LookPath(`scp`)
	if err != nil {
		return to, err
	}

	if len(to) == 0 {
		to, err = os.MkdirTemp(``, `scp`)

		if err != nil {
			return to, err
		}
	}

	err = exec.Command(scpexe, `-r`, from, to).Run()
	return to, err
}

// FirstLineOfSSH returns only the first line of a remote file by
// calling head on the file over an ssh connection. Otherwise, identical
// to LastLineOfSSH.
func FirstLineOfSSH(target string) (string, error) {
	u := ParseSSHURI(target)
	if u == nil {
		return ``, fmt.Errorf(`%q is not a valid SSH URI`, target)
	}
	return SSHOut(u.Addr, `head -1 `+u.Path)
}

// SSHURI is more restrictive than SSH might allow and includes the
// addition of ssh.head and ssh.tail schemas. Note that the Addr field
// must be kept in sync with the others if any fields are changed. This
// is avoid unnecessary indirection just to join the other strings. The
// UpdateAddr method has been added for convenience to do this.
type SSHURI struct {
	Schema string // scp, ssh, ssh.head, ssh.tail
	User   string // user
	Host   string // host
	Port   string // 22
	Path   string // /some/full/path OR rel/path
	Addr   string // user@host:22
}

// UpdateAddr simply sets the Addr string to match the other fields.
// This is provided to avoid indirection from a dynamic attribute. If
// the Host is missing sets Addr to blank.
func (u *SSHURI) UpdateAddr() {
	if len(u.Host) == 0 {
		u.Addr = ``
		return
	}
	u.Addr = u.Host
	if len(u.User) > 0 {
		u.Addr = u.User + `@` + u.Host
	}
	if len(u.Port) > 0 {
		u.Addr = u.Addr + `:` + u.Port
	}
}

// String fulfills the fmt.Stringer interface by printing the canonical
// URI format.
func (u SSHURI) String() string {
	str := fmt.Sprintf(`%v://%v`, u.Schema, u.Addr)
	if len(u.Path) > 0 {
		str += `/` + u.Path
	}
	return str
}

// ParseSSHURI return a parsed long or short form SSH URI or nil if
// unable to parse (no Host found).
func ParseSSHURI(target string) *SSHURI {

	// long form (URI)
	if strings.Index(target, `://`) > 0 {
		p := expr.SSHURI.FindStringSubmatch(target)
		if len(p) < 4 || len(p[4]) == 0 { // host
			return nil
		}

		return &SSHURI{
			Schema: p[1],
			Addr:   p[2],
			User:   p[3],
			Host:   p[4],
			Port:   p[5],
			Path:   p[6],
		}
	}

	// short form (not strictly a URI)
	// TODO

	return nil
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

// FirstFileIn returns the first file in the specified local directory
// (excluding directories, which are technically also "files").
func FirstFileIn(dir string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return ``, err
	}
	for _, file := range files {
		if !file.IsDir() {
			return file.Name(), nil
		}
	}
	return ``, fmt.Errorf(`file not found`)
}
