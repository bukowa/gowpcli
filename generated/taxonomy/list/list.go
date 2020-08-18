/*
## OPTIONS
	[--<field>=<value>]
	: Filter by one or more fields (see get_taxonomies() first parameter for a list of available fields).
	[--field=<field>]
	: Prints the value of a single field for each taxonomy.
	[--fields=<fields>]
	: Limit the output to specific taxonomy fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - count
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each term:
	* name
	* label
	* description
	* object_type
	* show_tagcloud
	* hierarchical
	* public
	These fields are optionally available:
	* count
## EXAMPLES
	    # List all taxonomies.
	    $ wp taxonomy list --format=csv
	    name,label,description,object_type,show_tagcloud,hierarchical,public
	    category,Categories,,post,1,1,1
	    post_tag,Tags,,post,1,,1
	    nav_menu,"Navigation Menus",,nav_menu_item,,,
	    link_category,"Link Categories",,link,1,,
	    post_format,Format,,post,,,1
	    # List all taxonomies with 'post' object type.
	    $ wp taxonomy list --object_type=post --fields=name,public
	    +-------------+--------+
	    | name        | public |
	    +-------------+--------+
	    | category    | 1      |
	    | post_tag    | 1      |
	    | post_format | 1      |
	    +-------------+--------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists registered taxonomies.
type List struct {
    FieldMap map[string]string // [--<field>=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"taxonomy", "list"}
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

