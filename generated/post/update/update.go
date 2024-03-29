/*
## OPTIONS
	<id>...
	: One or more IDs of posts to update.
	[--post_author=<post_author>]
	: The ID of the user who added the post. Default is the current user ID.
	[--post_date=<post_date>]
	: The date of the post. Default is the current time.
	[--post_date_gmt=<post_date_gmt>]
	: The date of the post in the GMT timezone. Default is the value of $post_date.
	[--post_content=<post_content>]
	: The post content. Default empty.
	[--post_content_filtered=<post_content_filtered>]
	: The filtered post content. Default empty.
	[--post_title=<post_title>]
	: The post title. Default empty.
	[--post_excerpt=<post_excerpt>]
	: The post excerpt. Default empty.
	[--post_status=<post_status>]
	: The post status. Default 'draft'.
	[--post_type=<post_type>]
	: The post type. Default 'post'.
	[--comment_status=<comment_status>]
	: Whether the post can accept comments. Accepts 'open' or 'closed'. Default is the value of 'default_comment_status' option.
	[--ping_status=<ping_status>]
	: Whether the post can accept pings. Accepts 'open' or 'closed'. Default is the value of 'default_ping_status' option.
	[--post_password=<post_password>]
	: The password to access the post. Default empty.
	[--post_name=<post_name>]
	: The post name. Default is the sanitized post title when creating a new post.
	[--to_ping=<to_ping>]
	: Space or carriage return-separated list of URLs to ping. Default empty.
	[--pinged=<pinged>]
	: Space or carriage return-separated list of URLs that have been pinged. Default empty.
	[--post_modified=<post_modified>]
	: The date when the post was last modified. Default is the current time.
	[--post_modified_gmt=<post_modified_gmt>]
	: The date when the post was last modified in the GMT timezone. Default is the current time.
	[--post_parent=<post_parent>]
	: Set this for the post it belongs to, if any. Default 0.
	[--menu_order=<menu_order>]
	: The order the post should be displayed in. Default 0.
	[--post_mime_type=<post_mime_type>]
	: The mime type of the post. Default empty.
	[--guid=<guid>]
	: Global Unique ID for referencing the post. Default empty.
	[--post_category=<post_category>]
	: Array of category names, slugs, or IDs. Defaults to value of the 'default_category' option.
	[--tags_input=<tags_input>]
	: Array of tag names, slugs, or IDs. Default empty.
	[--tax_input=<tax_input>]
	: Array of taxonomy terms keyed by their taxonomy name. Default empty.
	[--meta_input=<meta_input>]
	: Array in JSON format of post meta values keyed by their post meta key. Default empty.
	[<file>]
	: Read post content from <file>. If this value is present, the
	    `--post_content` argument will be ignored.
	  Passing `-` as the filename will cause post content to
	  be read from STDIN.
	--<field>=<value>
	: One or more fields to update. See wp_insert_post().
	[--defer-term-counting]
	: Recalculate term count in batch, for a performance boost.
## EXAMPLES
	    $ wp post update 123 --post_name=something --post_status=draft
	    Success: Updated post 123.
	    # Update a post with multiple meta values.
	    $ wp post update 123 --meta_input='{"key1":"value1","key2":"value2"}'
	    Success: Updated post 123.
	    # Update multiple posts at once.
	    $ wp post update 123 456 --post_author=789
	    Success: Updated post 123.
	    Success: Updated post 456.
	    # Update all posts of a given post type at once.
	    $ wp post update $(wp post list --post_type=page --format=ids) --post_author=123
	    Success: Updated post 123.
	    Success: Updated post 456.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates one or more existing posts.
type Update struct {
    Id []string // <id>...
    PostAuthor string // [--post_author=<post_author>]
    PostDate string // [--post_date=<post_date>]
    PostDateGmt string // [--post_date_gmt=<post_date_gmt>]
    PostContent string // [--post_content=<post_content>]
    PostContentFiltered string // [--post_content_filtered=<post_content_filtered>]
    PostTitle string // [--post_title=<post_title>]
    PostExcerpt string // [--post_excerpt=<post_excerpt>]
    PostStatus string // [--post_status=<post_status>]
    PostType string // [--post_type=<post_type>]
    CommentStatus string // [--comment_status=<comment_status>]
    PingStatus string // [--ping_status=<ping_status>]
    PostPassword string // [--post_password=<post_password>]
    PostName string // [--post_name=<post_name>]
    ToPing string // [--to_ping=<to_ping>]
    Pinged string // [--pinged=<pinged>]
    PostModified string // [--post_modified=<post_modified>]
    PostModifiedGmt string // [--post_modified_gmt=<post_modified_gmt>]
    PostParent string // [--post_parent=<post_parent>]
    MenuOrder string // [--menu_order=<menu_order>]
    PostMimeType string // [--post_mime_type=<post_mime_type>]
    Guid string // [--guid=<guid>]
    PostCategory string // [--post_category=<post_category>]
    TagsInput string // [--tags_input=<tags_input>]
    TaxInput string // [--tax_input=<tax_input>]
    MetaInput string // [--meta_input=<meta_input>]
    File string // [<file>]
    Field string // --<field>=<value>
    DeferTermCounting bool // [--defer-term-counting]
}

func (u Update) Args() []string {
    var args = []string{"post", "update"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    args = utils.MakeArg(args, "[--post_author=<post_author>]", u.PostAuthor)
    args = utils.MakeArg(args, "[--post_date=<post_date>]", u.PostDate)
    args = utils.MakeArg(args, "[--post_date_gmt=<post_date_gmt>]", u.PostDateGmt)
    args = utils.MakeArg(args, "[--post_content=<post_content>]", u.PostContent)
    args = utils.MakeArg(args, "[--post_content_filtered=<post_content_filtered>]", u.PostContentFiltered)
    args = utils.MakeArg(args, "[--post_title=<post_title>]", u.PostTitle)
    args = utils.MakeArg(args, "[--post_excerpt=<post_excerpt>]", u.PostExcerpt)
    args = utils.MakeArg(args, "[--post_status=<post_status>]", u.PostStatus)
    args = utils.MakeArg(args, "[--post_type=<post_type>]", u.PostType)
    args = utils.MakeArg(args, "[--comment_status=<comment_status>]", u.CommentStatus)
    args = utils.MakeArg(args, "[--ping_status=<ping_status>]", u.PingStatus)
    args = utils.MakeArg(args, "[--post_password=<post_password>]", u.PostPassword)
    args = utils.MakeArg(args, "[--post_name=<post_name>]", u.PostName)
    args = utils.MakeArg(args, "[--to_ping=<to_ping>]", u.ToPing)
    args = utils.MakeArg(args, "[--pinged=<pinged>]", u.Pinged)
    args = utils.MakeArg(args, "[--post_modified=<post_modified>]", u.PostModified)
    args = utils.MakeArg(args, "[--post_modified_gmt=<post_modified_gmt>]", u.PostModifiedGmt)
    args = utils.MakeArg(args, "[--post_parent=<post_parent>]", u.PostParent)
    args = utils.MakeArg(args, "[--menu_order=<menu_order>]", u.MenuOrder)
    args = utils.MakeArg(args, "[--post_mime_type=<post_mime_type>]", u.PostMimeType)
    args = utils.MakeArg(args, "[--guid=<guid>]", u.Guid)
    args = utils.MakeArg(args, "[--post_category=<post_category>]", u.PostCategory)
    args = utils.MakeArg(args, "[--tags_input=<tags_input>]", u.TagsInput)
    args = utils.MakeArg(args, "[--tax_input=<tax_input>]", u.TaxInput)
    args = utils.MakeArg(args, "[--meta_input=<meta_input>]", u.MetaInput)
    args = utils.MakeArg(args, "[<file>]", u.File)
    args = utils.MakeArg(args, "--<field>=<value>", u.Field)
    args = utils.MakeArg(args, "[--defer-term-counting]", u.DeferTermCounting)
    return args
}

