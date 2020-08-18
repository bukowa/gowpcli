package list

//Lists users.

type List struct {
    
    Role string // [--role=<role>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Network bool // [--network]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//Display WordPress users based on all arguments supported by
//[WP_User_Query()](https://developer.wordpress.org/reference/classes/wp_user_query/prepare_query/).
//
//## OPTIONS
//
//[--role=<role>]
//: Only display users with a certain role.
//
//[--<field>=<value>]
//: Control output by one or more arguments of WP_User_Query().
//
//[--network]
//: List all users in the network for multisite.
//
//[--field=<field>]
//: Prints the value of a single field for each user.
//
//[--fields=<fields>]
//: Limit the output to specific object fields.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - ids
//  - json
//  - count
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each user:
//
//* ID
//* user_login
//* display_name
//* user_email
//* user_registered
//* roles
//
//These fields are optionally available:
//
//* user_pass
//* user_nicename
//* user_url
//* user_activation_key
//* user_status
//* spam
//* deleted
//* caps
//* cap_key
//* allcaps
//* filter
//* url
//
//## EXAMPLES
//
//    # List user IDs
//    $ wp user list --field=ID
//    1
//
//    # List users with administrator role
//    $ wp user list --role=administrator --format=csv
//    ID,user_login,display_name,user_email,user_registered,roles
//    1,supervisor,supervisor,supervisor@gmail.com,"2016-06-03 04:37:00",administrator
//
//    # List users with only given fields
//    $ wp user list --fields=display_name,user_email --format=json
//    [{"display_name":"supervisor","user_email":"supervisor@gmail.com"}]
//
//    # List users ordered by the 'last_activity' meta value.
//    $ wp user list --meta_key=last_activity --orderby=meta_value_num
//
//