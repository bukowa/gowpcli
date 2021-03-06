/*
## OPTIONS
	<url>
	: URL to retrieve provider for.
	[--discover]
	: Whether to use oEmbed discovery or not. Defaults to true.
	[--limit-response-size=<size>]
	: Limit the size of the resulting HTML when using discovery. Default 150 KB (the standard WordPress limit). Not compatible with 'no-discover'.
	[--link-type=<json|xml>]
	: Whether to accept only a certain link type when using discovery. Defaults to any (json or xml), preferring json. Not compatible with 'no-discover'.
	---
	options:
	  - json
	  - xml
	---
## EXAMPLES
	    # Get the matching provider for the URL.
	    $ wp embed provider match https://www.youtube.com/watch?v=dQw4w9WgXcQ
	    https://www.youtube.com/oembed
	
 */
package match
import utils "github.com/bukowa/gowpcli"

// Match //Gets the matching provider for a given URL.
type Match struct {
    Url string // <url>
    Discover bool // [--discover]
    LimitResponseSize string // [--limit-response-size=<size>]
    LinkType string // [--link-type=<json|xml>]
}

func (m Match) Args() []string {
    var args = []string{"embed", "provider", "match"}
    args = utils.MakeArg(args, "<url>", m.Url)
    args = utils.MakeArg(args, "[--discover]", m.Discover)
    args = utils.MakeArg(args, "[--limit-response-size=<size>]", m.LimitResponseSize)
    args = utils.MakeArg(args, "[--link-type=<json|xml>]", m.LinkType)
    return args
}

