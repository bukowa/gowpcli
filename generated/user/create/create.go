/*
## OPTIONS
	<user-login>
	: The login of the user to create.
	<user-email>
	: The email address of the user to create.
	[--role=<role>]
	: The role of the user to create. Default: default role. Possible values
	include 'administrator', 'editor', 'author', 'contributor', 'subscriber'.
	[--user_pass=<password>]
	: The user password. Default: randomly generated.
	[--user_registered=<yyyy-mm-dd-hh-ii-ss>]
	: The date the user registered. Default: current date.
	[--display_name=<name>]
	: The display name.
	[--user_nicename=<nice_name>]
	: A string that contains a URL-friendly name for the user. The default is the user's username.
	[--user_url=<url>]
	: A string containing the user's URL for the user's web site.
	[--nickname=<nickname>]
	: The user's nickname, defaults to the user's username.
	[--first_name=<first_name>]
	: The user's first name.
	[--last_name=<last_name>]
	: The user's last name.
	[--description=<description>]
	: A string containing content about the user.
	[--rich_editing=<rich_editing>]
	: A string for whether to enable the rich editor or not. False if not empty.
	[--send-email]
	: Send an email to the user with their new account details.
	[--porcelain]
	: Output just the new user id.
## EXAMPLES
	    # Create user
	    $ wp user create bob bob@example.com --role=author
	    Success: Created user 3.
	    Password: k9**&I4vNH(&
	    # Create user without showing password upon success
	    $ wp user create ann ann@example.com --porcelain
	    4
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new user.
type Create struct {
    UserLogin string // <user-login>
    UserEmail string // <user-email>
    Role string // [--role=<role>]
    UserPass string // [--user_pass=<password>]
    UserRegistered string // [--user_registered=<yyyy-mm-dd-hh-ii-ss>]
    DisplayName string // [--display_name=<name>]
    UserNicename string // [--user_nicename=<nice_name>]
    UserUrl string // [--user_url=<url>]
    Nickname string // [--nickname=<nickname>]
    FirstName string // [--first_name=<first_name>]
    LastName string // [--last_name=<last_name>]
    Description string // [--description=<description>]
    RichEditing string // [--rich_editing=<rich_editing>]
    SendEmail bool // [--send-email]
    Porcelain bool // [--porcelain]
}

func (c Create) Args() []string {
    var args = []string{"user", "create"}
    args = utils.MakeArg(args, "<user-login>", c.UserLogin)
    args = utils.MakeArg(args, "<user-email>", c.UserEmail)
    args = utils.MakeArg(args, "[--role=<role>]", c.Role)
    args = utils.MakeArg(args, "[--user_pass=<password>]", c.UserPass)
    args = utils.MakeArg(args, "[--user_registered=<yyyy-mm-dd-hh-ii-ss>]", c.UserRegistered)
    args = utils.MakeArg(args, "[--display_name=<name>]", c.DisplayName)
    args = utils.MakeArg(args, "[--user_nicename=<nice_name>]", c.UserNicename)
    args = utils.MakeArg(args, "[--user_url=<url>]", c.UserUrl)
    args = utils.MakeArg(args, "[--nickname=<nickname>]", c.Nickname)
    args = utils.MakeArg(args, "[--first_name=<first_name>]", c.FirstName)
    args = utils.MakeArg(args, "[--last_name=<last_name>]", c.LastName)
    args = utils.MakeArg(args, "[--description=<description>]", c.Description)
    args = utils.MakeArg(args, "[--rich_editing=<rich_editing>]", c.RichEditing)
    args = utils.MakeArg(args, "[--send-email]", c.SendEmail)
    args = utils.MakeArg(args, "[--porcelain]", c.Porcelain)
    return args
}

