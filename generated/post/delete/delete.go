package delete

//Deletes an existing post.

type Delete struct {
    
    Id []string // <id>...
    
    Force bool // [--force]
    
    DeferTermCounting bool // [--defer-term-counting]
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of posts to delete.
//
//[--force]
//: Skip the trash bin.
//
//[--defer-term-counting]
//: Recalculate term count in batch, for a performance boost.
//
//## EXAMPLES
//
//    # Delete post skipping trash
//    $ wp post delete 123 --force
//    Success: Deleted post 123.
//
//    # Delete all pages
//    $ wp post delete $(wp post list --post_type='page' --format=ids)
//    Success: Trashed post 1164.
//    Success: Trashed post 1186.
//
//    # Delete all posts in the trash
//    $ wp post delete $(wp post list --post_status=trash --format=ids)
//    Success: Deleted post 1268.
//    Success: Deleted post 1294.
//
//