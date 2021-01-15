/*
## OPTIONS
	<user>...
	: The user login, user email or user ID of the user(s) to update.
	[--user_pass=<password>]
	: A string that contains the plain text password for the user.
	[--user_nicename=<nice_name>]
	: A string that contains a URL-friendly name for the user. The default is the user's username.
	[--user_url=<url>]
	: A string containing the user's URL for the user's web site.
	[--user_email=<email>]
	: A string containing the user's email address.
	[--display_name=<display_name>]
	: A string that will be shown on the site. Defaults to user's username.
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
	[--user_registered=<yyyy-mm-dd-hh-ii-ss>]
	: The date the user registered.
	[--role=<role>]
	: A string used to set the user's role.
	--<field>=<value>
	: One or more fields to update. For accepted fields, see wp_update_user().
	[--skip-email]
	: Don't send an email notification to the user.
## EXAMPLES
	    # Update user
	    $ wp user update 123 --display_name=Mary --user_pass=marypass
	    Success: Updated user 123.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates an existing user.
type Update struct {
    User []string // <user>...
    UserPass string // [--user_pass=<password>]
    UserNicename string // [--user_nicename=<nice_name>]
    UserUrl string // [--user_url=<url>]
    UserEmail string // [--user_email=<email>]
    DisplayName string // [--display_name=<display_name>]
    Nickname string // [--nickname=<nickname>]
    FirstName string // [--first_name=<first_name>]
    LastName string // [--last_name=<last_name>]
    Description string // [--description=<description>]
    RichEditing string // [--rich_editing=<rich_editing>]
    UserRegistered string // [--user_registered=<yyyy-mm-dd-hh-ii-ss>]
    Role string // [--role=<role>]
    Field string // --<field>=<value>
    SkipEmail bool // [--skip-email]
}

func (u Update) Args() []string {
    var args = []string{"user", "update"}
    args = utils.MakeArg(args, "<user>...", u.User)
    args = utils.MakeArg(args, "[--user_pass=<password>]", u.UserPass)
    args = utils.MakeArg(args, "[--user_nicename=<nice_name>]", u.UserNicename)
    args = utils.MakeArg(args, "[--user_url=<url>]", u.UserUrl)
    args = utils.MakeArg(args, "[--user_email=<email>]", u.UserEmail)
    args = utils.MakeArg(args, "[--display_name=<display_name>]", u.DisplayName)
    args = utils.MakeArg(args, "[--nickname=<nickname>]", u.Nickname)
    args = utils.MakeArg(args, "[--first_name=<first_name>]", u.FirstName)
    args = utils.MakeArg(args, "[--last_name=<last_name>]", u.LastName)
    args = utils.MakeArg(args, "[--description=<description>]", u.Description)
    args = utils.MakeArg(args, "[--rich_editing=<rich_editing>]", u.RichEditing)
    args = utils.MakeArg(args, "[--user_registered=<yyyy-mm-dd-hh-ii-ss>]", u.UserRegistered)
    args = utils.MakeArg(args, "[--role=<role>]", u.Role)
    args = utils.MakeArg(args, "--<field>=<value>", u.Field)
    args = utils.MakeArg(args, "[--skip-email]", u.SkipEmail)
    return args
}

