package structure

//Updates the permalink structure.

type Structure struct {
    
    Permastruct string // <permastruct>
    
    CategoryBase string // [--category-base=<base>]
    
    TagBase string // [--tag-base=<base>]
    
    Hard bool // [--hard]
    
}

//Sets the post permalink structure to the specified pattern.
//
//To regenerate a .htaccess file with WP-CLI, you'll need to add
//the mod_rewrite module to your [WP-CLI config](https://make.wordpress.org/cli/handbook/config/#config-files).
//For example:
//
//```
//apache_modules:
//  - mod_rewrite
//```
//
//## OPTIONS
//
//<permastruct>
//: The new permalink structure to apply.
//
//[--category-base=<base>]
//: Set the base for category permalinks, i.e. '/category/'.
//
//[--tag-base=<base>]
//: Set the base for tag permalinks, i.e. '/tag/'.
//
//[--hard]
//: Perform a hard flush - update `.htaccess` rules as well as rewrite rules in database.
//
//## EXAMPLES
//
//    $ wp rewrite structure '/%year%/%monthnum%/%postname%/'
//    Success: Rewrite structure set.
//
//