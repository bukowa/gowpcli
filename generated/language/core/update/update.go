package update

//Updates installed languages for core.

type Update struct {
    
    DryRun bool // [--dry-run]
    
}

//## OPTIONS
//
//[--dry-run]
//: Preview which translations would be updated.
//
//## EXAMPLES
//
//    $ wp language core update
//    Updating 'Japanese' translation for WordPress 4.9.2...
//    Downloading translation from https://downloads.wordpress.org/translation/core/4.9.2/ja.zip...
//    Translation updated successfully.
//    Success: Updated 1/1 translation.
//
//