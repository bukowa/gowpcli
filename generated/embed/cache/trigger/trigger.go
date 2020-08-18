package trigger

//Triggers the caching of all oEmbed results for a given post.

type Trigger struct {
    
    PostId string // <post_id>
    
}

//oEmbed caches for a post are stored in the post's metadata.
//
//## OPTIONS
//
//<post_id>
//: ID of the post to do the caching for.
//
//## EXAMPLES
//
//    # Triggers cache for a post
//    $ wp embed cache trigger 456
//    Success: Caching triggered!
//
//