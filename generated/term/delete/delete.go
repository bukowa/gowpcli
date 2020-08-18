package delete

//Deletes an existing term.

type Delete struct {
    
    Taxonomy string // <taxonomy>
    
    Term []string // <term>...
    
    By string // [--by=<field>]
    
}

//Errors if the term doesn't exist, or there was a problem in deleting it.
//
//## OPTIONS
//
//<taxonomy>
//: Taxonomy of the term to delete.
//
//<term>...
//: One or more IDs or slugs of terms to delete.
//
//[--by=<field>]
//: Explicitly handle the term value as a slug or id.
//---
//default: id
//options:
//  - slug
//  - id
//---
//
//## EXAMPLES
//
//    # Delete post category by id
//    $ wp term delete category 15
//    Deleted category 15.
//    Success: Deleted 1 of 1 terms.
//
//    # Delete post category by slug
//    $ wp term delete category apple --by=slug
//    Deleted category 15.
//    Success: Deleted 1 of 1 terms.
//
//    # Delete all post tags
//    $ wp term list post_tag --field=term_id | xargs wp term delete post_tag
//    Deleted post_tag 159.
//    Deleted post_tag 160.
//    Deleted post_tag 161.
//    Success: Deleted 3 of 3 terms.
//
//