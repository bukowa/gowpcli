package search

//Searches the WordPress.org plugin directory.

type Search struct {
    
    Search string // <search>
    
    Page string // [--page=<page>]
    
    PerPage string // [--per-page=<per-page>]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//Displays plugins in the WordPress.org plugin directory matching a given
//search query.
//
//## OPTIONS
//
//<search>
//: The string to search for.
//
//[--page=<page>]
//: Optional page to display.
//---
//default: 1
//---
//
//[--per-page=<per-page>]
//: Optional number of results to display.
//---
//default: 10
//---
//
//[--field=<field>]
//: Prints the value of a single field for each plugin.
//
//[--fields=<fields>]
//: Ask for specific fields from the API. Defaults to name,slug,author_profile,rating. Acceptable values:
//
//    **name**: Plugin Name
//    **slug**: Plugin Slug
//    **version**: Current Version Number
//    **author**: Plugin Author
//    **author_profile**: Plugin Author Profile
//    **contributors**: Plugin Contributors
//    **requires**: Plugin Minimum Requirements
//    **tested**: Plugin Tested Up To
//    **compatibility**: Plugin Compatible With
//    **rating**: Plugin Rating in Percent and Total Number
//    **ratings**: Plugin Ratings for each star (1-5)
//    **num_ratings**: Number of Plugin Ratings
//    **homepage**: Plugin Author's Homepage
//    **description**: Plugin's Description
//    **short_description**: Plugin's Short Description
//    **sections**: Plugin Readme Sections: description, installation, FAQ, screenshots, other notes, and changelog
//    **downloaded**: Plugin Download Count
//    **last_updated**: Plugin's Last Update
//    **added**: Plugin's Date Added to wordpress.org Repository
//    **tags**: Plugin's Tags
//    **versions**: Plugin's Available Versions with D/L Link
//    **donate_link**: Plugin's Donation Link
//    **banners**: Plugin's Banner Image Link
//    **icons**: Plugin's Icon Image Link
//    **active_installs**: Plugin's Number of Active Installs
//    **contributors**: Plugin's List of Contributors
//    **url**: Plugin's URL on wordpress.org
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - count
//  - json
//  - yaml
//---
//
//## EXAMPLES
//
//    $ wp plugin search dsgnwrks --per-page=20 --format=json
//    Success: Showing 3 of 3 plugins.
//    [{"name":"DsgnWrks Instagram Importer Debug","slug":"dsgnwrks-instagram-importer-debug","rating":0},{"name":"DsgnWrks Instagram Importer","slug":"dsgnwrks-instagram-importer","rating":84},{"name":"DsgnWrks Twitter Importer","slug":"dsgnwrks-twitter-importer","rating":80}]
//
//    $ wp plugin search dsgnwrks --fields=name,version,slug,rating,num_ratings
//    Success: Showing 3 of 3 plugins.
//    +-----------------------------------+---------+-----------------------------------+--------+-------------+
//    | name                              | version | slug                              | rating | num_ratings |
//    +-----------------------------------+---------+-----------------------------------+--------+-------------+
//    | DsgnWrks Instagram Importer Debug | 0.1.6   | dsgnwrks-instagram-importer-debug | 0      | 0           |
//    | DsgnWrks Instagram Importer       | 1.3.7   | dsgnwrks-instagram-importer       | 84     | 23          |
//    | DsgnWrks Twitter Importer         | 1.1.1   | dsgnwrks-twitter-importer         | 80     | 1           |
//    +-----------------------------------+---------+-----------------------------------+--------+-------------+
//
//