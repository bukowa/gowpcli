/*
## OPTIONS
	<file>...
	: Path to file or files to be imported. Supports the glob(3) capabilities of the current shell.
	    If file is recognized as a URL (for example, with a scheme of http or ftp), the file will be
	    downloaded to a temp file before being sideloaded.
	[--post_id=<post_id>]
	: ID of the post to attach the imported files to.
	[--title=<title>]
	: Attachment title (post title field).
	[--caption=<caption>]
	: Caption for attachent (post excerpt field).
	[--alt=<alt_text>]
	: Alt text for image (saved as post meta).
	[--desc=<description>]
	: "Description" field (post content) of attachment post.
	[--skip-copy]
	: If set, media files (local only) are imported to the library but not moved on disk.
	File names will not be run through wp_unique_filename() with this set.
	[--preserve-filetime]
	: Use the file modified time as the post published & modified dates.
	Remote files will always use the current time.
	[--featured_image]
	: If set, set the imported image as the Featured Image of the post its attached to.
	[--porcelain]
	: Output just the new attachment ID.
## EXAMPLES
	    # Import all jpgs in the current user's "Pictures" directory, not attached to any post.
	    $ wp media import ~/Pictures/**\/*.jpg
	    Imported file '/home/person/Pictures/landscape-photo.jpg' as attachment ID 1751.
	    Imported file '/home/person/Pictures/fashion-icon.jpg' as attachment ID 1752.
	    Success: Imported 2 of 2 items.
	    # Import a local image and set it to be the post thumbnail for a post.
	    $ wp media import ~/Downloads/image.png --post_id=123 --title="A downloaded picture" --featured_image
	    Imported file '/home/person/Downloads/image.png' as attachment ID 1753 and attached to post 123 as featured image.
	    Success: Imported 1 of 1 images.
	    # Import a local image, but set it as the featured image for all posts.
	    # 1. Import the image and get its attachment ID.
	    # 2. Assign the attachment ID as the featured image for all posts.
	    $ ATTACHMENT_ID="$(wp media import ~/Downloads/image.png --porcelain)"
	    $ wp post list --post_type=post --format=ids | xargs -d ' ' -I % wp post meta add % _thumbnail_id $ATTACHMENT_ID
	    Success: Added custom field.
	    Success: Added custom field.
	    # Import an image from the web.
	    $ wp media import http://s.wordpress.org/style/images/wp-header-logo.png --title='The WordPress logo' --alt="Semantic personal publishing"
	    Imported file 'http://s.wordpress.org/style/images/wp-header-logo.png' as attachment ID 1755.
	    Success: Imported 1 of 1 images.
	    # Get the URL for an attachment after import.
	    $ wp media import http://s.wordpress.org/style/images/wp-header-logo.png --porcelain | xargs -I {} wp post list --post__in={} --field=url --post_type=attachment
	    http://wordpress-develop.dev/wp-header-logo/
	
 */
package import_
import utils "github.com/bukowa/gowpcli"

// Import_ //Creates attachments from local files or URLs.
type Import_ struct {
    File []string // <file>...
    PostId string // [--post_id=<post_id>]
    Title string // [--title=<title>]
    Caption string // [--caption=<caption>]
    Alt string // [--alt=<alt_text>]
    Desc string // [--desc=<description>]
    SkipCopy bool // [--skip-copy]
    PreserveFiletime bool // [--preserve-filetime]
    FeaturedImage bool // [--featured_image]
    Porcelain bool // [--porcelain]
}

func (i Import_) Args() []string {
    var args = []string{"media", "import"}
    args = utils.MakeArg(args, "<file>...", i.File)
    args = utils.MakeArg(args, "[--post_id=<post_id>]", i.PostId)
    args = utils.MakeArg(args, "[--title=<title>]", i.Title)
    args = utils.MakeArg(args, "[--caption=<caption>]", i.Caption)
    args = utils.MakeArg(args, "[--alt=<alt_text>]", i.Alt)
    args = utils.MakeArg(args, "[--desc=<description>]", i.Desc)
    args = utils.MakeArg(args, "[--skip-copy]", i.SkipCopy)
    args = utils.MakeArg(args, "[--preserve-filetime]", i.PreserveFiletime)
    args = utils.MakeArg(args, "[--featured_image]", i.FeaturedImage)
    args = utils.MakeArg(args, "[--porcelain]", i.Porcelain)
    return args
}

