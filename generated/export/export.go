/*
## INFO
	Generates one or more WXR files containing authors, terms, posts,
	comments, and attachments. WXR files do not include site configuration
	(options) or the attachment files themselves.
## OPTIONS
	[--dir=<dirname>]
	: Full path to directory where WXR export files should be stored. Defaults
	to current working directory.
	[--stdout]
	: Output the whole XML using standard output (incompatible with --dir=)
	[--skip_comments]
	: Don't include comments in the WXR export file.
	[--max_file_size=<MB>]
	: A single export file should have this many megabytes. -1 for unlimited.
	---
	default: 15
	---
## FILTERS
	[--start_date=<date>]
	: Export only posts published after this date, in format YYYY-MM-DD.
	[--end_date=<date>]
	: Export only posts published before this date, in format YYYY-MM-DD.
	[--post_type=<post-type>]
	: Export only posts with this post_type. Separate multiple post types with a
	comma.
	---
	default: any
	---
	[--post_type__not_in=<post-type>]
	: Export all post types except those identified. Separate multiple post types
	with a comma. Defaults to none.
	[--post__in=<pid>]
	: Export all posts specified as a comma- or space-separated list of IDs.
	Post's attachments won't be exported unless --with_attachments is specified.
	[--with_attachments]
	: Force including attachments in case --post__in has been specified.
	[--start_id=<pid>]
	: Export only posts with IDs greater than or equal to this post ID.
	[--max_num_posts=<num>]
	: Export no more than <num> posts (excluding attachments).
	[--author=<author>]
	: Export only posts by this author. Can be either user login or user ID.
	[--category=<name>]
	: Export only posts in this category.
	[--post_status=<status>]
	: Export only posts with this status.
	[--filename_format=<format>]
	: Use a custom format for export filenames. Defaults to '{site}.wordpress.{date}.{n}.xml'.
## EXAMPLES
	    # Export posts published by the user between given start and end date
	    $ wp export --dir=/tmp/ --user=admin --post_type=post --start_date=2011-01-01 --end_date=2011-12-31
	    Starting export process...
	    Writing to file /tmp/staging.wordpress.2016-05-24.000.xml
	    Success: All done with export.
	    # Export posts by IDs
	    $ wp export --dir=/tmp/ --post__in=123,124,125
	    Starting export process...
	    Writing to file /tmp/staging.wordpress.2016-05-24.000.xml
	    Success: All done with export.
	    # Export a random subset of content
	    $ wp export --post__in="$(wp post list --post_type=post --orderby=rand --posts_per_page=8 --format=ids)"
	    Starting export process...
	    Writing to file /var/www/example.com/public_html/staging.wordpress.2016-05-24.000.xml
	    Success: All done with export.
	
 */
package export
import utils "github.com/bukowa/gowpcli"

// Export //Exports WordPress content to a WXR file.
type Export struct {
    Dir string // [--dir=<dirname>]
    Stdout bool // [--stdout]
    SkipComments bool // [--skip_comments]
    MaxFileSize string // [--max_file_size=<MB>]
    StartDate string // [--start_date=<date>]
    EndDate string // [--end_date=<date>]
    PostType string // [--post_type=<post-type>]
    PostTypeNotIn string // [--post_type__not_in=<post-type>]
    PostIn string // [--post__in=<pid>]
    WithAttachments bool // [--with_attachments]
    StartId string // [--start_id=<pid>]
    MaxNumPosts string // [--max_num_posts=<num>]
    Author string // [--author=<author>]
    Category string // [--category=<name>]
    PostStatus string // [--post_status=<status>]
    FilenameFormat string // [--filename_format=<format>]
}

func (e Export) Args() []string {
    var args = []string{"export"}
    args = utils.MakeArg(args, "[--dir=<dirname>]", e.Dir)
    args = utils.MakeArg(args, "[--stdout]", e.Stdout)
    args = utils.MakeArg(args, "[--skip_comments]", e.SkipComments)
    args = utils.MakeArg(args, "[--max_file_size=<MB>]", e.MaxFileSize)
    args = utils.MakeArg(args, "[--start_date=<date>]", e.StartDate)
    args = utils.MakeArg(args, "[--end_date=<date>]", e.EndDate)
    args = utils.MakeArg(args, "[--post_type=<post-type>]", e.PostType)
    args = utils.MakeArg(args, "[--post_type__not_in=<post-type>]", e.PostTypeNotIn)
    args = utils.MakeArg(args, "[--post__in=<pid>]", e.PostIn)
    args = utils.MakeArg(args, "[--with_attachments]", e.WithAttachments)
    args = utils.MakeArg(args, "[--start_id=<pid>]", e.StartId)
    args = utils.MakeArg(args, "[--max_num_posts=<num>]", e.MaxNumPosts)
    args = utils.MakeArg(args, "[--author=<author>]", e.Author)
    args = utils.MakeArg(args, "[--category=<name>]", e.Category)
    args = utils.MakeArg(args, "[--post_status=<status>]", e.PostStatus)
    args = utils.MakeArg(args, "[--filename_format=<format>]", e.FilenameFormat)
    return args
}

