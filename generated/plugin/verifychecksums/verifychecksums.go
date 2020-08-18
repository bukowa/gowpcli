package verifychecksums

//Verifies plugin files against WordPress.org's checksums.

type VerifyChecksums struct {
    
    Plugin []string // [<plugin>...]
    
    All bool // [--all]
    
    Strict bool // [--strict]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[<plugin>...]
//: One or more plugins to verify.
//
//[--all]
//: If set, all plugins will be verified.
//
//[--strict]
//: If set, even "soft changes" like readme.txt changes will trigger
//checksum errors.
//
//[--format=<format>]
//: Render output in a specific format.
//---
//default: table
//options:
//  - table
//  - json
//  - csv
//  - yaml
//  - count
//---
//
//## EXAMPLES
//
//    # Verify the checksums of all installed plugins
//    $ wp plugin verify-checksums --all
//    Success: Verified 8 of 8 plugins.
//
//    # Verify the checksums of a single plugin, Akismet in this case
//    $ wp plugin verify-checksums akismet
//    Success: Verified 1 of 1 plugins.
//
//