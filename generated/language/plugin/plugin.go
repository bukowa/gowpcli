/*
## EXAMPLES
	    # Install the Dutch plugin language pack.
	    $ wp language plugin install hello-dolly nl_NL
	    Success: Language installed.
	    # Uninstall the Dutch plugin language pack.
	    $ wp language plugin uninstall hello-dolly nl_NL
	    Success: Language uninstalled.
	    # List installed plugin language packages.
	    $ wp language plugin list --status=installed
	    +----------+--------------+-------------+-----------+-----------+---------------------+
	    | language | english_name | native_name | status    | update    | updated             |
	    +----------+--------------+-------------+-----------+-----------+---------------------+
	    | nl_NL    | Dutch        | Nederlands  | installed | available | 2016-05-13 08:12:50 |
	    +----------+--------------+-------------+-----------+-----------+---------------------+
	
 */
package plugin


// Plugin //Installs, activates, and manages plugin language packs.
type Plugin struct {
}

func (p Plugin) Args() []string {
    var args = []string{"language", "plugin"}
    return args
}

