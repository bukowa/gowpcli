package update

//Updates installed languages for one or more themes.

type Update struct {
    
    Theme []string // [<theme>...]
    
    All bool // [--all]
    
    DryRun bool // [--dry-run]
    
}

//## OPTIONS
//
//[<theme>...]
//: One or more themes to update languages for.
//
//[--all]
//: If set, languages for all themes will be updated.
//
//[--dry-run]
//: Preview which translations would be updated.
//
//## EXAMPLES
//
//    $ wp language theme update --all
//    Updating 'Japanese' translation for Twenty Fifteen 1.5...
//    Downloading translation from https://downloads.wordpress.org/translation/theme/twentyfifteen/1.5/ja.zip...
//    Translation updated successfully.
//    Success: Updated 1/1 translation.
//
//