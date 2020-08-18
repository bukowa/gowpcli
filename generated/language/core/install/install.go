package install

//Installs a given language.

type Install struct {
    
    Language []string // <language>...
    
    Activate bool // [--activate]
    
}

//Downloads the language pack from WordPress.org.
//
//## OPTIONS
//
//<language>...
//: Language code to install.
//
//[--activate]
//: If set, the language will be activated immediately after install.
//
//## EXAMPLES
//
//    # Install the Japanese language.
//    $ wp language core install ja
//    Downloading translation from https://downloads.wordpress.org/translation/core/4.9.8/ja.zip...
//    Unpacking the update...
//    Installing the latest version...
//    Translation updated successfully.
//    Language 'ja' installed.
//    Success: Installed 1 of 1 languages.
//
//