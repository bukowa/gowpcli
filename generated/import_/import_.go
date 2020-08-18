/*
## INFO
	Provides a command line interface to the WordPress Importer plugin, for
	performing data migrations.
## OPTIONS
	<file>...
	: Path to one or more valid WXR files for importing. Directories are also accepted.
	--authors=<authors>
	: How the author mapping should be handled. Options are 'create', 'mapping.csv', or 'skip'. The first will create any non-existent users from the WXR file. The second will read author mapping associations from a CSV, or create a CSV for editing if the file path doesn't exist. The CSV requires two columns, and a header row like "old_user_login,new_user_login". The last option will skip any author mapping.
	[--skip=<data-type>]
	: Skip importing specific data. Supported options are: 'attachment' and 'image_resize' (skip time-consuming thumbnail generation).
## EXAMPLES
	    # Import content from a WXR file
	    $ wp import example.wordpress.2016-06-21.xml --authors=create
	    Starting the import process...
	    Processing post #1 ("Hello world!") (post_type: post)
	    -- 1 of 1
	    -- Tue, 21 Jun 2016 05:31:12 +0000
	    -- Imported post as post_id #1
	    Success: Finished importing from 'example.wordpress.2016-06-21.xml' file.
	
 */
package import_
import utils "github.com/bukowa/gowpcli"

// Import_ //Imports content from a given WXR file.
type Import_ struct {
    File []string // <file>...
    Authors string // --authors=<authors>
    Skip string // [--skip=<data-type>]
}

func (i Import_) Args() []string {
    var args = []string{"import"}
    args = utils.MakeArg(args, "<file>...", i.File)
    args = utils.MakeArg(args, "--authors=<authors>", i.Authors)
    args = utils.MakeArg(args, "[--skip=<data-type>]", i.Skip)
    return args
}

