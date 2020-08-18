package gowpcli_test

import (
	"reflect"
	"testing"
	"github.com/bukowa/gowpcli/generated/core/install"
)

func TestInstall_Args(t *testing.T) {
	t.Errorf("")
	type fields struct {
		Url           string
		Title         string
		AdminUser     string
		AdminPassword string
		AdminEmail    string
		SkipEmail     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{name: "1", fields: fields{
			Url:           "http://asd.com",
			Title:         "test title",
			AdminUser:     "usr",
			AdminPassword: "passwd",
			AdminEmail:    "",
			SkipEmail:     true,
		}, want: []string{
			"core",
			"install",
			"--url=http://asd.com",
			"--title=test title",
			"--admin_user=usr",
			"--admin_password=passwd",
			"--skip-email",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := install.Install{
				Url:           tt.fields.Url,
				Title:         tt.fields.Title,
				AdminUser:     tt.fields.AdminUser,
				AdminPassword: tt.fields.AdminPassword,
				AdminEmail:    tt.fields.AdminEmail,
				SkipEmail:     tt.fields.SkipEmail,
			}
			args := i.Args()
			//sort.Strings(args)
			//sort.Strings(tt.want)
			if got := args; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Args() = %v, want %v", got, tt.want)
			}
		})
	}
}
