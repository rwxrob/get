/* Package expr contains the regular expressions used to parse the
* SSHURI. Consider using the get.ParseSSHURI* functions instead of using
* these directly.
 */
package expr

import "regexp"

// SSHURI is a variation on canonical SSH URIs that includes ssh.head
// and ssh.tail schemas. The regular expression is otherwise identical
// to that used by the git command.
var SSHURI = regexp.MustCompile(`^(ssh(?:.(?:head|tail))?|scp)://((?:([A-Za-z0-9]+)@)?([A-Za-z0-9.]+)(?::([0-9]{1,7}))?)(?:/(\S+))?$`)

// SSHURIShort is the same used by git and ssh and scp (but not
// techically a canonically compliant URI).
var SSHURIShort = regexp.MustCompile(`^(?:([A-Za-z0-9]+)@)?([A-Za-z0-9.]+)(?::(\S+))?$`)
