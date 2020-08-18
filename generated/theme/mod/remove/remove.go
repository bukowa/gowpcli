package remove

//Removes one or more theme mods.

type Remove struct {
    
    Mod []string // [<mod>...]
    
    All bool // [--all]
    
}

//## OPTIONS
//
//[<mod>...]
//: One or more mods to remove.
//
//[--all]
//: Remove all theme mods.
//
//## EXAMPLES
//
//    # Remove all theme mods.
//    $ wp theme mod remove --all
//    Success: Theme mods removed.
//
//    # Remove single theme mod.
//    $ wp theme mod remove background_color
//    Success: 1 mod removed.
//
//    # Remove multiple theme mods.
//    $ wp theme mod remove background_color header_textcolor
//    Success: 2 mods removed.
//
//