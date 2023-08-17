/* Package expr contains the regular expressions used to parse the
* SSHURI. Consider using the get.ParseSSHURI* functions instead of using
* these directly.
 */
package expr

import "regexp"

var SSHURI = regexp.MustCompile(`^(ssh(?:.(?:head|tail))?|scp)://((?:([A-Za-z0-9]+)@)?([A-Za-z0-9.]+)(?::([0-9]{1,7}))?)(?:/(\S+))?$`)

var SSHURIShort = regexp.MustCompile(`^(?:([A-Za-z0-9]+)@)?([A-Za-z0-9.]+)(?::(\S+))?$`)
