/*
## INFO
	Defaults to updating WordPress to the latest version.
	If you see "Error: Another update is currently in progress.", you may
	need to run `wp option delete core_updater.lock` after verifying another
	update isn't actually running.
## OPTIONS
	[<zip>]
	: Path to zip file to use, instead of downloading from wordpress.org.
	[--minor]
	: Only perform updates for minor releases (e.g. update from WP 4.3 to 4.3.3 instead of 4.4.2).
	[--version=<version>]
	: Update to a specific version, instead of to the latest version. Alternatively accepts 'nightly'.
	[--force]
	: Update even when installed WP version is greater than the requested version.
	[--locale=<locale>]
	: Select which language you want to download.
## EXAMPLES
	    # Update WordPress
	    $ wp core update
	    Updating to version 4.5.2 (en_US)...
	    Downloading update from https://downloads.wordpress.org/release/wordpress-4.5.2-no-content.zip...
	    Unpacking the update...
	    Cleaning up files...
	    No files found that need cleaning up
	    Success: WordPress updated successfully.
	    # Update WordPress to latest version of 3.8 release
	    $ wp core update --version=3.8 ../latest.zip
	    Updating to version 3.8 ()...
	    Unpacking the update...
	    Cleaning up files...
	    File removed: wp-admin/js/tags-box.js
	    ...
	    File removed: wp-admin/js/updates.min.
	    377 files cleaned up
	    Success: WordPress updated successfully.
	    # Update WordPress to 3.1 forcefully
	    $ wp core update --version=3.1 --force
	    Updating to version 3.1 (en_US)...
	    Downloading update from https://wordpress.org/wordpress-3.1.zip...
	    Unpacking the update...
	    Warning: Checksums not available for WordPress 3.1/en_US. Please cleanup files manually.
	    Success: WordPress updated successfully.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates WordPress to a newer version.
type Update struct {
    Zip string // [<zip>]
    Minor bool // [--minor]
    Version string // [--version=<version>]
    Force bool // [--force]
    Locale string // [--locale=<locale>]
}

func (u Update) Args() []string {
    var args = []string{"core", "update"}
    args = utils.MakeArg(args, "[<zip>]", u.Zip)
    args = utils.MakeArg(args, "[--minor]", u.Minor)
    args = utils.MakeArg(args, "[--version=<version>]", u.Version)
    args = utils.MakeArg(args, "[--force]", u.Force)
    args = utils.MakeArg(args, "[--locale=<locale>]", u.Locale)
    return args
}

