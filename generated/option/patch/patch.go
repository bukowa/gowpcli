/*
## OPTIONS
	<action>
	: Patch action to perform.
	---
	options:
	  - insert
	  - update
	  - delete
	---
	<key>
	: The option name.
	<key-path>...
	: The name(s) of the keys within the value to locate the value to patch.
	[<value>]
	: The new value. If omitted, the value is read from STDIN.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
## EXAMPLES
	    # Add 'bar' to the 'foo' key on an option with name 'option_name'
	    $ wp option patch insert option_name foo bar
	    Success: Updated 'option_name' option.
	    # Update the value of 'foo' key to 'new' on an option with name 'option_name'
	    $ wp option patch update option_name foo new
	    Success: Updated 'option_name' option.
	    # Set nested value of 'bar' key to value we have in the patch file on an option with name 'option_name'.
	    $ wp option patch update option_name foo bar < patch
	    Success: Updated 'option_name' option.
	    # Update the value for the key 'not-a-key' which is not exist on an option with name 'option_name'.
	    $ wp option patch update option_name foo not-a-key new-value
	    Error: No data exists for key "not-a-key"
	    # Update the value for the key 'foo' without passing value on an option with name 'option_name'.
	    $ wp option patch update option_name foo
	    Error: Please provide value to update.
	    # Delete the nested key 'bar' under 'foo' key on an option with name 'option_name'.
	    $ wp option patch delete option_name foo bar
	    Success: Updated 'option_name' option.
	
 */
package patch
import utils "github.com/bukowa/gowpcli"

// Patch //Updates a nested value in an option.
type Patch struct {
    Action string // <action>
    Key string // <key>
    KeyPath []string // <key-path>...
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (p Patch) Args() []string {
    var args = []string{"option", "patch"}
    args = utils.MakeArg(args, "<action>", p.Action)
    args = utils.MakeArg(args, "<key>", p.Key)
    args = utils.MakeArg(args, "<key-path>...", p.KeyPath)
    args = utils.MakeArg(args, "[<value>]", p.Value)
    args = utils.MakeArg(args, "[--format=<format>]", p.Format)
    return args
}

