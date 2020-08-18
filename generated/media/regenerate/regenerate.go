/*
## OPTIONS
	[<attachment-id>...]
	: One or more IDs of the attachments to regenerate.
	[--image_size=<image_size>]
	: Name of the image size to regenerate. Only thumbnails of this image size will be regenerated, thumbnails of other image sizes will not.
	[--skip-delete]
	: Skip deletion of the original thumbnails. If your thumbnails are linked from sources outside your control, it's likely best to leave them around. Defaults to false.
	[--only-missing]
	: Only generate thumbnails for images missing image sizes.
	[--yes]
	: Answer yes to the confirmation message. Confirmation only shows when no IDs passed as arguments.
## EXAMPLES
	    # Regenerate thumbnails for given attachment IDs.
	    $ wp media regenerate 123 124 125
	    Found 3 images to regenerate.
	    1/3 Regenerated thumbnails for "Vertical Image" (ID 123).
	    2/3 Regenerated thumbnails for "Horizontal Image" (ID 124).
	    3/3 Regenerated thumbnails for "Beautiful Picture" (ID 125).
	    Success: Regenerated 3 of 3 images.
	    # Regenerate all thumbnails, without confirmation.
	    $ wp media regenerate --yes
	    Found 3 images to regenerate.
	    1/3 Regenerated thumbnails for "Sydney Harbor Bridge" (ID 760).
	    2/3 Regenerated thumbnails for "Boardwalk" (ID 757).
	    3/3 Regenerated thumbnails for "Sunburst Over River" (ID 756).
	    Success: Regenerated 3 of 3 images.
	    # Re-generate all thumbnails that have IDs between 1000 and 2000.
	    $ seq 1000 2000 | xargs wp media regenerate
	    Found 4 images to regenerate.
	    1/4 Regenerated thumbnails for "Vertical Featured Image" (ID 1027).
	    2/4 Regenerated thumbnails for "Horizontal Featured Image" (ID 1022).
	    3/4 Regenerated thumbnails for "Unicorn Wallpaper" (ID 1045).
	    4/4 Regenerated thumbnails for "I Am Worth Loving Wallpaper" (ID 1023).
	    Success: Regenerated 4 of 4 images.
	    # Re-generate only the thumbnails of "large" image size for all images.
	    $ wp media regenerate --image_size=large
	    Do you really want to regenerate the "large" image size for all images? [y/n] y
	    Found 3 images to regenerate.
	    1/3 Regenerated "large" thumbnail for "Sydney Harbor Bridge" (ID 760).
	    2/3 No "large" thumbnail regeneration needed for "Boardwalk" (ID 757).
	    3/3 Regenerated "large" thumbnail for "Sunburst Over River" (ID 756).
	    Success: Regenerated 3 of 3 images.
	
 */
package regenerate
import utils "github.com/bukowa/gowpcli"

// Regenerate //Regenerates thumbnails for one or more attachments.
type Regenerate struct {
    AttachmentId []string // [<attachment-id>...]
    ImageSize string // [--image_size=<image_size>]
    SkipDelete bool // [--skip-delete]
    OnlyMissing bool // [--only-missing]
    Yes bool // [--yes]
}

func (r Regenerate) Args() []string {
    var args = []string{"media", "regenerate"}
    args = utils.MakeArg(args, "[<attachment-id>...]", r.AttachmentId)
    args = utils.MakeArg(args, "[--image_size=<image_size>]", r.ImageSize)
    args = utils.MakeArg(args, "[--skip-delete]", r.SkipDelete)
    args = utils.MakeArg(args, "[--only-missing]", r.OnlyMissing)
    args = utils.MakeArg(args, "[--yes]", r.Yes)
    return args
}

