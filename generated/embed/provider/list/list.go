/*
## OPTIONS
	[--field=<field>]
	: Display the value of a single field
	[--fields=<fields>]
	: Limit the output to specific fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	---
	[--force-regex]
	: Turn the asterisk-type provider URLs into regexes.
## AVAILABLE FIELDS
	These fields will be displayed by default for each provider:
	* format
	* endpoint
	This field is optionally available:
	* regex
## EXAMPLES
	    # List format,endpoint fields of available providers.
	    $ wp embed provider list --fields=format,endpoint
	    +------------------------------+-----------------------------------------+
	    | format                       | endpoint                                |
	    +------------------------------+-----------------------------------------+
	    | #https?://youtu\.be/.*#i     | https://www.youtube.com/oembed          |
	    | #https?://flic\.kr/.*#i      | https://www.flickr.com/services/oembed/ |
	    | #https?://wordpress\.tv/.*#i | https://wordpress.tv/oembed/            |
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all available oEmbed providers.
type List struct {
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    ForceRegex bool // [--force-regex]
}

func (l List) Args() []string {
    var args = []string{"embed", "provider", "list"}
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--force-regex]", l.ForceRegex)
    return args
}

