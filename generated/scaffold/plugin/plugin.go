/*
## INFO
	The following files are always generated:
	* `plugin-slug.php` is the main PHP plugin file.
	* `readme.txt` is the readme file for the plugin.
	* `package.json` needed by NPM holds various metadata relevant to the project. Packages: `grunt`, `grunt-wp-i18n` and `grunt-wp-readme-to-markdown`. Scripts: `start`, `readme`, `i18n`.
	* `Gruntfile.js` is the JS file containing Grunt tasks. Tasks: `i18n` containing `addtextdomain` and `makepot`, `readme` containing `wp_readme_to_markdown`.
	* `.editorconfig` is the configuration file for Editor.
	* `.gitignore` tells which files (or patterns) git should ignore.
	* `.distignore` tells which files and folders should be ignored in distribution.
	The following files are also included unless the `--skip-tests` is used:
	* `phpunit.xml.dist` is the configuration file for PHPUnit.
	* `.travis.yml` is the configuration file for Travis CI. Use `--ci=<provider>` to select a different service.
	* `bin/install-wp-tests.sh` configures the WordPress test suite and a test database.
	* `tests/bootstrap.php` is the file that makes the current plugin active when running the test suite.
	* `tests/test-sample.php` is a sample file containing test cases.
	* `.phpcs.xml.dist` is a collection of PHP_CodeSniffer rules.
## OPTIONS
	<slug>
	: The internal name of the plugin.
	[--dir=<dirname>]
	: Put the new plugin in some arbitrary directory path. Plugin directory will be path plus supplied slug.
	[--plugin_name=<title>]
	: What to put in the 'Plugin Name:' header.
	[--plugin_description=<description>]
	: What to put in the 'Description:' header.
	[--plugin_author=<author>]
	: What to put in the 'Author:' header.
	[--plugin_author_uri=<url>]
	: What to put in the 'Author URI:' header.
	[--plugin_uri=<url>]
	: What to put in the 'Plugin URI:' header.
	[--skip-tests]
	: Don't generate files for unit testing.
	[--ci=<provider>]
	: Choose a configuration file for a continuous integration provider.
	---
	default: travis
	options:
	  - travis
	  - circle
	  - gitlab
	---
	[--activate]
	: Activate the newly generated plugin.
	[--activate-network]
	: Network activate the newly generated plugin.
	[--force]
	: Overwrite files that already exist.
## EXAMPLES
	    $ wp scaffold plugin sample-plugin
	    Success: Created plugin files.
	    Success: Created test files.
	
 */
package plugin
import utils "github.com/bukowa/gowpcli"

// Plugin //Generates starter code for a plugin.
type Plugin struct {
    Slug string // <slug>
    Dir string // [--dir=<dirname>]
    PluginName string // [--plugin_name=<title>]
    PluginDescription string // [--plugin_description=<description>]
    PluginAuthor string // [--plugin_author=<author>]
    PluginAuthorUri string // [--plugin_author_uri=<url>]
    PluginUri string // [--plugin_uri=<url>]
    SkipTests bool // [--skip-tests]
    Ci string // [--ci=<provider>]
    Activate bool // [--activate]
    ActivateNetwork bool // [--activate-network]
    Force bool // [--force]
}

func (p Plugin) Args() []string {
    var args = []string{"scaffold", "plugin"}
    args = utils.MakeArg(args, "<slug>", p.Slug)
    args = utils.MakeArg(args, "[--dir=<dirname>]", p.Dir)
    args = utils.MakeArg(args, "[--plugin_name=<title>]", p.PluginName)
    args = utils.MakeArg(args, "[--plugin_description=<description>]", p.PluginDescription)
    args = utils.MakeArg(args, "[--plugin_author=<author>]", p.PluginAuthor)
    args = utils.MakeArg(args, "[--plugin_author_uri=<url>]", p.PluginAuthorUri)
    args = utils.MakeArg(args, "[--plugin_uri=<url>]", p.PluginUri)
    args = utils.MakeArg(args, "[--skip-tests]", p.SkipTests)
    args = utils.MakeArg(args, "[--ci=<provider>]", p.Ci)
    args = utils.MakeArg(args, "[--activate]", p.Activate)
    args = utils.MakeArg(args, "[--activate-network]", p.ActivateNetwork)
    args = utils.MakeArg(args, "[--force]", p.Force)
    return args
}

