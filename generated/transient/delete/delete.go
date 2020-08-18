package delete

//Deletes a transient value.

type Delete struct {
    
    Key string // [<key>]
    
    Network bool // [--network]
    
    All bool // [--all]
    
    Expired bool // [--expired]
    
}

//For a more complete explanation of the transient cache, including the
//network|site cache, please see docs for `wp transient`.
//
//## OPTIONS
//
//[<key>]
//: Key for the transient.
//
//[--network]
//: Delete the value of a network|site transient. On single site, this is
//is a specially-named cache key. On multisite, this is a global cache
//(instead of local to the site).
//
//[--all]
//: Delete all transients.
//
//[--expired]
//: Delete all expired transients.
//
//## EXAMPLES
//
//    # Delete transient.
//    $ wp transient delete sample_key
//    Success: Transient deleted.
//
//    # Delete expired transients.
//    $ wp transient delete --expired
//    Success: 12 expired transients deleted from the database.
//
//    # Delete expired site transients.
//    $ wp transient delete --expired --network
//    Success: 1 expired transient deleted from the database.
//
//    # Delete all transients.
//    $ wp transient delete --all
//    Success: 14 transients deleted from the database.
//
//    # Delete all site transients.
//    $ wp transient delete --all --network
//    Success: 2 transients deleted from the database.
//
//    # Delete all transients in a multsite.
//    $ wp transient delete --all --network && wp site list --field=url | xargs -n1 -I % wp --url=% transient delete --all
//
//