/*
## EXAMPLES
	    # Dump the list of installed commands.
	    $ wp cli cmd-dump
	    {"name":"wp","description":"Manage WordPress through the command-line.","longdesc":"\n\n
 */
package cmddump


// CmdDump //Dumps the list of installed commands, as JSON.
type CmdDump struct {
}

func (c CmdDump) Args() []string {
    var args = []string{"cli", "cmd-dump"}
    return args
}

