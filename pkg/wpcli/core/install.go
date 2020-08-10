package core

import "github.com/bukowa/gowpcli/pkg/wpcli/utils"

var InstallSuccess = "Success: WordPress installed successfully."

type Install struct {
	Url           string
	Title         string
	AdminUser     string
	AdminPassword string
	AdminEmail    string
	SkipEmail     bool
}

func (i *Install) Args() (args []string) {
	var command = []string{"core", "install"}
	args = append(args, command...)

	if i.Url != "" {
		args = utils.MakeArg(args, "url", i.Url)
	}
	if i.Title != "" {
		args = utils.MakeArg(args, "title", i.Title)
	}
	if i.AdminUser != "" {
		args = utils.MakeArg(args, "admin_user", i.AdminUser)
	}
	if i.AdminPassword != "" {
		args = utils.MakeArg(args, "admin_password", i.AdminPassword)
	}
	if i.AdminEmail != "" {
		args = utils.MakeArg(args, "admin_email", i.AdminEmail)
	}
	if i.SkipEmail {
		args = utils.MakeArg(args, "skip-email", "")
	}
	return args
}
