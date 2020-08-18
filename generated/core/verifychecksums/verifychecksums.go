/*
## INFO
	Downloads md5 checksums for the current version from WordPress.org, and
	compares those checksums against the currently installed files.
	For security, avoids loading WordPress when verifying checksums.
	If you experience issues verifying from this command, ensure you are
	passing the relevant `--locale` and `--version` arguments according to
	the values from the `Dashboard->Updates` menu in the admin area of the
	site.
## OPTIONS
	[--version=<version>]
	: Verify checksums against a specific version of WordPress.
	[--locale=<locale>]
	: Verify checksums against a specific locale of WordPress.
## EXAMPLES
	    # Verify checksums
	    $ wp core verify-checksums
	    Success: WordPress installation verifies against checksums.
	    # Verify checksums for given WordPress version
	    $ wp core verify-checksums --version=4.0
	    Success: WordPress installation verifies against checksums.
	    # Verify checksums for given locale
	    $ wp core verify-checksums --locale=en_US
	    Success: WordPress installation verifies against checksums.
	    # Verify checksums for given locale
	    $ wp core verify-checksums --locale=ja
	    Warning: File doesn't verify against checksum: wp-includes/version.php
	    Warning: File doesn't verify against checksum: readme.html
	    Warning: File doesn't verify against checksum: wp-config-sample.php
	    Error: WordPress installation doesn't verify against checksums.
	
 */
package verifychecksums
import utils "github.com/bukowa/gowpcli"

// VerifyChecksums //Verifies WordPress files against WordPress.org's checksums.
type VerifyChecksums struct {
    Version string // [--version=<version>]
    Locale string // [--locale=<locale>]
}

func (v VerifyChecksums) Args() []string {
    var args = []string{"core", "verify-checksums"}
    args = utils.MakeArg(args, "[--version=<version>]", v.Version)
    args = utils.MakeArg(args, "[--locale=<locale>]", v.Locale)
    return args
}

