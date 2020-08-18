/*
## INFO
	Creates a specified number of new posts with dummy data.
## OPTIONS
	[--count=<number>]
	: How many posts to generate?
	---
	default: 100
	---
	[--post_type=<type>]
	: The type of the generated posts.
	---
	default: post
	---
	[--post_status=<status>]
	: The status of the generated posts.
	---
	default: publish
	---
	[--post_title=<post_title>]
	: The post title.
	---
	default:
	---
	[--post_author=<login>]
	: The author of the generated posts.
	---
	default:
	---
	[--post_date=<yyyy-mm-dd-hh-ii-ss>]
	: The date of the generated posts. Default: current date
	[--post_date_gmt=<yyyy-mm-dd-hh-ii-ss>]
	: The GMT date of the generated posts. Default: value of post_date (or current date if it's not set)
	[--post_content]
	: If set, the command reads the post_content from STDIN.
	[--max_depth=<number>]
	: For hierarchical post types, generate child posts down to a certain depth.
	---
	default: 1
	---
	[--format=<format>]
	: Render output in a particular format.
	---
	default: progress
	options:
	  - progress
	  - ids
	---
## EXAMPLES
	    # Generate posts.
	    $ wp post generate --count=10 --post_type=page --post_date=1999-01-04
	    Generating posts  100% [================================================] 0:01 / 0:04
	    # Generate posts with fetched content.
	    $ curl -N http://loripsum.net/api/5 | wp post generate --post_content --count=10
	      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
	                                     Dload  Upload   Total   Spent    Left  Speed
	    100  2509  100  2509    0     0    616      0  0:00:04  0:00:04 --:--:--   616
	    Generating posts  100% [================================================] 0:01 / 0:04
	    # Add meta to every generated posts.
	    $ wp post generate --format=ids | xargs -d ' ' -I % wp post meta add % foo bar
	    Success: Added custom field.
	    Success: Added custom field.
	    Success: Added custom field.
	
 */
package generate
import utils "github.com/bukowa/gowpcli"

// Generate //Generates some posts.
type Generate struct {
    Count string // [--count=<number>]
    PostType string // [--post_type=<type>]
    PostStatus string // [--post_status=<status>]
    PostTitle string // [--post_title=<post_title>]
    PostAuthor string // [--post_author=<login>]
    PostDate string // [--post_date=<yyyy-mm-dd-hh-ii-ss>]
    PostDateGmt string // [--post_date_gmt=<yyyy-mm-dd-hh-ii-ss>]
    PostContent bool // [--post_content]
    MaxDepth string // [--max_depth=<number>]
    Format string // [--format=<format>]
}

func (g Generate) Args() []string {
    var args = []string{"post", "generate"}
    args = utils.MakeArg(args, "[--count=<number>]", g.Count)
    args = utils.MakeArg(args, "[--post_type=<type>]", g.PostType)
    args = utils.MakeArg(args, "[--post_status=<status>]", g.PostStatus)
    args = utils.MakeArg(args, "[--post_title=<post_title>]", g.PostTitle)
    args = utils.MakeArg(args, "[--post_author=<login>]", g.PostAuthor)
    args = utils.MakeArg(args, "[--post_date=<yyyy-mm-dd-hh-ii-ss>]", g.PostDate)
    args = utils.MakeArg(args, "[--post_date_gmt=<yyyy-mm-dd-hh-ii-ss>]", g.PostDateGmt)
    args = utils.MakeArg(args, "[--post_content]", g.PostContent)
    args = utils.MakeArg(args, "[--max_depth=<number>]", g.MaxDepth)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

