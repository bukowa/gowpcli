package posttype

//Retrieves details on the site's registered post types.

type PostType struct {
    
}

//Get information on WordPress' built-in and the site's [custom post types](https://developer.wordpress.org/plugins/post-types/).
//
//## EXAMPLES
//
//    # Get details about a post type
//    $ wp post-type get page --fields=name,label,hierarchical --format=json
//    {"name":"page","label":"Pages","hierarchical":true}
//
//    # List post types with 'post' capability type
//    $ wp post-type list --capability_type=post --fields=name,public
//    +---------------+--------+
//    | name          | public |
//    +---------------+--------+
//    | post          | 1      |
//    | attachment    | 1      |
//    | revision      |        |
//    | nav_menu_item |        |
//    +---------------+--------+
//
//