package generate

//Generates some number of new dummy comments.

type Generate struct {
    
    Count string // [--count=<number>]
    
    PostId string // [--post_id=<post-id>]
    
    Format string // [--format=<format>]
    
}

//Creates a specified number of new comments with dummy data.
//
//## OPTIONS
//
//[--count=<number>]
//: How many comments to generate?
//---
//default: 100
//---
//
//[--post_id=<post-id>]
//: Assign comments to a specific post.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: progress
//options:
//  - progress
//  - ids
//---
//
//## EXAMPLES
//
//    # Generate comments for the given post.
//    $ wp comment generate --format=ids --count=3 --post_id=123
//    138 139 140
//
//    # Add meta to every generated comment.
//    $ wp comment generate --format=ids --count=3 | xargs -d ' ' -I % wp comment meta add % foo bar
//    Success: Added custom field.
//    Success: Added custom field.
//    Success: Added custom field.
//
//