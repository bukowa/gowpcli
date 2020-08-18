package clear

//Deletes all oEmbed caches for a given post.

type Clear struct {
    
    PostId string // <post_id>
    
}

//oEmbed caches for a post are stored in the post's metadata.
//
//## OPTIONS
//
//<post_id>
//: ID of the post to clear the cache for.
//
//## EXAMPLES
//
//    # Clear cache for a post
//    $ wp embed cache clear 123
//    Success: Cleared oEmbed cache.
//
//