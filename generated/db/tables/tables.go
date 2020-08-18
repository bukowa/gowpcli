package tables

//Lists the database tables.

type Tables struct {
    
    Table []string // [<table>...]
    
    Scope string // [--scope=<scope>]
    
    Network bool // [--network]
    
    AllTablesWithPrefix bool // [--all-tables-with-prefix]
    
    AllTables bool // [--all-tables]
    
    Format string // [--format=<format>]
    
}

//Defaults to all tables registered to the $wpdb database handler.
//
//## OPTIONS
//
//[<table>...]
//: List tables based on wildcard search, e.g. 'wp_*_options' or 'wp_post?'.
//
//[--scope=<scope>]
//: Can be all, global, ms_global, blog, or old tables. Defaults to all.
//
//[--network]
//: List all the tables in a multisite install.
//
//[--all-tables-with-prefix]
//: List all tables that match the table prefix even if not registered on $wpdb. Overrides --network.
//
//[--all-tables]
//: List all tables in the database, regardless of the prefix, and even if not registered on $wpdb. Overrides --all-tables-with-prefix.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: list
//options:
//  - list
//  - csv
//---
//
//## EXAMPLES
//
//    # List tables for a single site, without shared tables like 'wp_users'
//    $ wp db tables --scope=blog --url=sub.example.com
//    wp_3_posts
//    wp_3_comments
//    wp_3_options
//    wp_3_postmeta
//    wp_3_terms
//    wp_3_term_taxonomy
//    wp_3_term_relationships
//    wp_3_termmeta
//    wp_3_commentmeta
//
//    # Export only tables for a single site
//    $ wp db export --tables=$(wp db tables --url=sub.example.com --format=csv)
//    Success: Exported to wordpress_dbase.sql
//
//