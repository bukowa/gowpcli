package install

//Installs one or more plugins.

type Install struct {
    
    Pluginzipurl []string // <plugin|zip|url>...
    
    Version string // [--version=<version>]
    
    Force bool // [--force]
    
    Activate bool // [--activate]
    
    ActivateNetwork bool // [--activate-network]
    
}

//## OPTIONS
//
//<plugin|zip|url>...
//: One or more plugins to install. Accepts a plugin slug, the path to a local zip file, or a URL to a remote zip file.
//
//[--version=<version>]
//: If set, get that particular version from wordpress.org, instead of the
//stable version.
//
//[--force]
//: If set, the command will overwrite any installed version of the plugin, without prompting
//for confirmation.
//
//[--activate]
//: If set, the plugin will be activated immediately after install.
//
//[--activate-network]
//: If set, the plugin will be network activated immediately after install
//
//## EXAMPLES
//
//    # Install the latest version from wordpress.org and activate
//    $ wp plugin install bbpress --activate
//    Installing bbPress (2.5.9)
//    Downloading install package from https://downloads.wordpress.org/plugin/bbpress.2.5.9.zip...
//    Using cached file '/home/vagrant/.wp-cli/cache/plugin/bbpress-2.5.9.zip'...
//    Unpacking the package...
//    Installing the plugin...
//    Plugin installed successfully.
//    Activating 'bbpress'...
//    Plugin 'bbpress' activated.
//    Success: Installed 1 of 1 plugins.
//
//    # Install the development version from wordpress.org
//    $ wp plugin install bbpress --version=dev
//    Installing bbPress (Development Version)
//    Downloading install package from https://downloads.wordpress.org/plugin/bbpress.zip...
//    Unpacking the package...
//    Installing the plugin...
//    Plugin installed successfully.
//    Success: Installed 1 of 1 plugins.
//
//    # Install from a local zip file
//    $ wp plugin install ../my-plugin.zip
//    Unpacking the package...
//    Installing the plugin...
//    Plugin installed successfully.
//    Success: Installed 1 of 1 plugins.
//
//    # Install from a remote zip file
//    $ wp plugin install http://s3.amazonaws.com/bucketname/my-plugin.zip?AWSAccessKeyId=123&Expires=456&Signature=abcdef
//    Downloading install package from http://s3.amazonaws.com/bucketname/my-plugin.zip?AWSAccessKeyId=123&Expires=456&Signature=abcdef
//    Unpacking the package...
//    Installing the plugin...
//    Plugin installed successfully.
//    Success: Installed 1 of 1 plugins.
//
//    # Update from a remote zip file
//    $ wp plugin install https://github.com/envato/wp-envato-market/archive/master.zip --force
//    Downloading install package from https://github.com/envato/wp-envato-market/archive/master.zip
//    Unpacking the package...
//    Installing the plugin...
//    Renamed Github-based project from 'wp-envato-market-master' to 'wp-envato-market'.
//    Plugin updated successfully
//    Success: Installed 1 of 1 plugins.
//
//    # Forcefully re-install all installed plugins
//    $ wp plugin install $(wp plugin list --field=name) --force
//    Installing Akismet (3.1.11)
//    Downloading install package from https://downloads.wordpress.org/plugin/akismet.3.1.11.zip...
//    Unpacking the package...
//    Installing the plugin...
//    Removing the old version of the plugin...
//    Plugin updated successfully
//    Success: Installed 1 of 1 plugins.
//
//