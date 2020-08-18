/*
## INFO
	Helpful for diagnostic purposes, this command shares:
	* OS information.
	* Shell information.
	* PHP binary used.
	* PHP binary version.
	* php.ini configuration file used (which is typically different than web).
	* WP-CLI root dir: where WP-CLI is installed (if non-Phar install).
	* WP-CLI global config: where the global config YAML file is located.
	* WP-CLI project config: where the project config YAML file is located.
	* WP-CLI version: currently installed version.
	See [config docs](https://wp-cli.org/config/) for more details on global
	and project config YAML files.
## OPTIONS
	[--format=<format>]
	: Render output in a particular format.
	---
	default: list
	options:
	  - list
	  - json
	---
## EXAMPLES
	    # Display various data about the CLI environment.
	    $ wp cli info
	    OS:  Linux 4.10.0-42-generic #46~16.04.1-Ubuntu SMP Mon Dec 4 15:57:59 UTC 2017 x86_64
	    Shell:   /usr/bin/zsh
	    PHP binary:  /usr/bin/php
	    PHP version: 7.1.12-1+ubuntu16.04.1+deb.sury.org+1
	    php.ini used:    /etc/php/7.1/cli/php.ini
	    WP-CLI root dir:    phar://wp-cli.phar
	    WP-CLI packages dir:    /home/person/.wp-cli/packages/
	    WP-CLI global config:
	    WP-CLI project config:
	    WP-CLI version: 1.5.0
	
 */
package info
import utils "github.com/bukowa/gowpcli"

// Info //Prints various details about the WP-CLI environment.
type Info struct {
    Format string // [--format=<format>]
}

func (i Info) Args() []string {
    var args = []string{"cli", "info"}
    args = utils.MakeArg(args, "[--format=<format>]", i.Format)
    return args
}

