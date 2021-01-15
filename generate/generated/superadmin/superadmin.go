/*
## EXAMPLES
	    # List user with super-admin capabilities
	    $ wp super-admin list
	    supervisor
	    administrator
	    # Grant super-admin privileges to the user.
	    $ wp super-admin add superadmin2
	    Success: Granted super-admin capabilities.
	    # Revoke super-admin privileges to the user.
	    $ wp super-admin remove superadmin2
	    Success: Revoked super-admin capabilities.
	
 */
package superadmin


// SuperAdmin //Lists, adds, or removes super admin users on a multisite installation.
type SuperAdmin struct {
}

func (s SuperAdmin) Args() []string {
    var args = []string{"super-admin"}
    return args
}

