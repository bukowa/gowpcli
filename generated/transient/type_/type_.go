package type_

//Determines the type of transients implementation.

type Type_ struct {
    
}

//Indicates whether the transients API is using an object cache or the
//database.
//
//For a more complete explanation of the transient cache, including the
//network|site cache, please see docs for `wp transient`.
//
//## EXAMPLES
//
//    $ wp transient type
//    Transients are saved to the database.
//
//