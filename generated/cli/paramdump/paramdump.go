/*
## OPTIONS
	[--with-values]
	: Display current values also.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: json
	options:
	  - var_export
	  - json
	---
## EXAMPLES
	    # Dump the list of global parameters.
	    $ wp cli param-dump --format=var_export
	    array (
	      'path' =>
	      array (
	        'runtime' => '=<path>',
	        'file' => '<path>',
	        'synopsis' => '',
	        'default' => NULL,
	        'multiple' => false,
	        'desc' => 'Path to the WordPress files.',
	      ),
	      'url' =>
	      array (
	
 */
package paramdump
import utils "github.com/bukowa/gowpcli"

// ParamDump //Dumps the list of global parameters, as JSON or in var_export format.
type ParamDump struct {
    WithValues bool // [--with-values]
    Format string // [--format=<format>]
}

func (p ParamDump) Args() []string {
    var args = []string{"cli", "param-dump"}
    args = utils.MakeArg(args, "[--with-values]", p.WithValues)
    args = utils.MakeArg(args, "[--format=<format>]", p.Format)
    return args
}

