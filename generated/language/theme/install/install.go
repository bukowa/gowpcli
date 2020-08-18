package install

//Installs a given language for a theme.

type Install struct {
    
    Theme string // [<theme>]
    
    All bool // [--all]
    
    Language []string // <language>...
    
    Format string // [--format=<format>]
    
}

//Downloads the language pack from WordPress.org.
//
//## OPTIONS
//
//[<theme>]
//: Theme to install language for.
//
//[--all]
//: If set, languages for all themes will be installed.
//
//<language>...
//: Language code to install.
//
//[--format=<format>]
//: Render output in a particular format. Used when installing languages for all themes.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//  - summary
//---
//
//## EXAMPLES
//
//    # Install the Japanese language for Twenty Seventeen.
//    $ wp language theme install twentyseventeen ja
//    Downloading translation from https://downloads.wordpress.org/translation/theme/twentyseventeen/1.3/ja.zip...
//    Unpacking the update...
//    Installing the latest version...
//    Translation updated successfully.
//    Language 'ja' installed.
//    Success: Installed 1 of 1 languages.
//
//