/*
## INFO
	Starting with WordPress 4.9, embeds that aren't associated with a specific post will be cached in
	a new oembed_cache post type. There can be more than one such entry for a url depending on attributes and context.
	Not to be confused with oEmbed caches for a given post which are stored in the post's metadata.
## OPTIONS
	<url>
	: URL to retrieve oEmbed data for.
	[--width=<width>]
	: Width of the embed in pixels. Part of cache key so must match. Defaults to `content_width` if set else 500px, so is theme and context dependent.
	[--height=<height>]
	: Height of the embed in pixels. Part of cache key so must match. Defaults to 1.5 * default width (`content_width` or 500px), to a maximum of 1000px.
	[--discover]
	: Whether to search with the discover attribute set or not. Part of cache key so must match. If not given, will search with attribute: unset, '1', '0', returning first.
## EXAMPLES
	    # Find cache post ID for a given URL.
	    $ wp embed cache find https://www.youtube.com/watch?v=dQw4w9WgXcQ --width=500
	    123
	
 */
package find
import utils "github.com/bukowa/gowpcli"

// Find //Finds an oEmbed cache post ID for a given URL.
type Find struct {
    Url string // <url>
    Width string // [--width=<width>]
    Height string // [--height=<height>]
    Discover bool // [--discover]
}

func (f Find) Args() []string {
    var args = []string{"embed", "cache", "find"}
    args = utils.MakeArg(args, "<url>", f.Url)
    args = utils.MakeArg(args, "[--width=<width>]", f.Width)
    args = utils.MakeArg(args, "[--height=<height>]", f.Height)
    args = utils.MakeArg(args, "[--discover]", f.Discover)
    return args
}

