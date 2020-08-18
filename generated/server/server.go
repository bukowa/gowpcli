/*
## INFO
	Uses `php -S` to launch a web server serving the WordPress webroot.
	<http://php.net/manual/en/features.commandline.webserver.php>
	Importantly, PHP's built-in web server doesn't support `.htaccess` files.
	If this is a requirement, please use a more advanced web server.
## OPTIONS
	[--host=<host>]
	: The hostname to bind the server to.
	---
	default: localhost
	---
	[--port=<port>]
	: The port number to bind the server to.
	---
	default: 8080
	---
	[--docroot=<path>]
	: The path to use as the document root. If the path global parameter is
	set, the default value is it.
	[--config=<file>]
	: Configure the server with a specific .ini file.
## EXAMPLES
	    # Make the instance available on any address (with port 8080)
	    $ wp server --host=0.0.0.0
	    PHP 5.6.9 Development Server started at Tue May 24 01:27:11 2016
	    Listening on http://0.0.0.0:8080
	    Document root is /
	    Press Ctrl-C to quit.
	    # Run on port 80 (for multisite)
	    $ wp server --host=localhost.localdomain --port=80
	    PHP 5.6.9 Development Server started at Tue May 24 01:30:06 2016
	    Listening on http://localhost1.localdomain1:80
	    Document root is /
	    Press Ctrl-C to quit.
	    # Configure the server with a specific .ini file
	    $ wp server --config=development.ini
	    PHP 7.0.9 Development Server started at Mon Aug 22 12:09:04 2016
	    Listening on http://localhost:8080
	    Document root is /
	    Press Ctrl-C to quit.
	
 */
package server
import utils "github.com/bukowa/gowpcli"

// Server //Launches PHP's built-in web server for a specific WordPress installation.
type Server struct {
    Host string // [--host=<host>]
    Port string // [--port=<port>]
    Docroot string // [--docroot=<path>]
    Config string // [--config=<file>]
}

func (s Server) Args() []string {
    var args = []string{"server"}
    args = utils.MakeArg(args, "[--host=<host>]", s.Host)
    args = utils.MakeArg(args, "[--port=<port>]", s.Port)
    args = utils.MakeArg(args, "[--docroot=<path>]", s.Docroot)
    args = utils.MakeArg(args, "[--config=<file>]", s.Config)
    return args
}

