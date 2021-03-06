/*
## INFO
	Note: because code is executed within a method, global variables need
	to be explicitly globalized.
## OPTIONS
	<php-code>
	: The code to execute, as a string.
	[--skip-wordpress]
	: Execute code without loading WordPress.
## EXAMPLES
	    # Display WordPress content directory.
	    $ wp eval 'echo WP_CONTENT_DIR;'
	    /var/www/wordpress/wp-content
	    # Generate a random number.
	    $ wp eval 'echo rand();' --skip-wordpress
	    479620423
	
 */
package eval
import utils "github.com/bukowa/gowpcli"

// Eval //Executes arbitrary PHP code.
type Eval struct {
    PhpCode string // <php-code>
    SkipWordpress bool // [--skip-wordpress]
}

func (e Eval) Args() []string {
    var args = []string{"eval"}
    args = utils.MakeArg(args, "<php-code>", e.PhpCode)
    args = utils.MakeArg(args, "[--skip-wordpress]", e.SkipWordpress)
    return args
}

