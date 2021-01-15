/*
## INFO
	The following files are generated by default:
	* `phpunit.xml.dist` is the configuration file for PHPUnit.
	* `.travis.yml` is the configuration file for Travis CI. Use `--ci=<provider>` to select a different service.
	* `bin/install-wp-tests.sh` configures the WordPress test suite and a test database.
	* `tests/bootstrap.php` is the file that makes the current theme active when running the test suite.
	* `tests/test-sample.php` is a sample file containing the actual tests.
	* `.phpcs.xml.dist` is a collection of PHP_CodeSniffer rules.
	Learn more from the [plugin unit tests documentation](https://make.wordpress.org/cli/handbook/plugin-unit-tests/).
## ENVIRONMENT
	The `tests/bootstrap.php` file looks for the WP_TESTS_DIR environment
	variable.
## OPTIONS
	[<theme>]
	: The name of the theme to generate test files for.
	[--dir=<dirname>]
	: Generate test files for a non-standard theme path. If no theme slug is specified, the directory name is used.
	[--ci=<provider>]
	: Choose a configuration file for a continuous integration provider.
	---
	default: travis
	options:
	  - travis
	  - circle
	  - gitlab
	  - bitbucket
	---
	[--force]
	: Overwrite files that already exist.
## EXAMPLES
	    # Generate unit test files for theme 'twentysixteenchild'.
	    $ wp scaffold theme-tests twentysixteenchild
	    Success: Created test files.
	
 */
package themetests
import utils "github.com/bukowa/gowpcli"

// ThemeTests //Generates files needed for running PHPUnit tests in a theme.
type ThemeTests struct {
    Theme string // [<theme>]
    Dir string // [--dir=<dirname>]
    Ci string // [--ci=<provider>]
    Force bool // [--force]
}

func (t ThemeTests) Args() []string {
    var args = []string{"scaffold", "theme-tests"}
    args = utils.MakeArg(args, "[<theme>]", t.Theme)
    args = utils.MakeArg(args, "[--dir=<dirname>]", t.Dir)
    args = utils.MakeArg(args, "[--ci=<provider>]", t.Ci)
    args = utils.MakeArg(args, "[--force]", t.Force)
    return args
}

