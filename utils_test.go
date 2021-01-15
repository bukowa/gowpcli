package gowpcli

import (
	"errors"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

type BasicHandler struct{}

func (b BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	w.Write([]byte("1"))
	w.WriteHeader(200)
}

func Test_formatParam(t *testing.T) {
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{
			k: "[--post_types=<post-types>]",
			v: "test",
		}, want: "--post_types=test"},
		{name: "2", args: args{
			k: "[--comment_status=<comment_status>]",
			v: "test1",
		}, want: "--comment_status=test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatNamed(tt.args.k, tt.args.v); got != tt.want {
				t.Errorf("formatNamed() = %v, want %v, len %v:%v", got, tt.want, len(got), len(tt.want))
			}
		})
	}
}


func TestRequiredOptional(t *testing.T) {
	// check if we handle required / optional
	for k, v := range TESTS {
		if !opt(k) && !isRequired(k) {
			t.Error(v)
		}
	}
	for k, v := range TESTS {
		if strings.HasPrefix(k, "--") && !isRequired(k) {
			t.Error(v)
		}
		if strings.HasPrefix(k, "[") && !opt(k) {
			t.Error(v)
		}
		if strings.HasPrefix(k, "--") && !isRequired(k) {
			t.Error(v)
		}
	}
}

func TestMulti(t *testing.T) {
	forkv(t, func(k, v string) error {
		if strings.Index(k, "...") != -1 && !mult(k) {
			return errors.New(k)
		}
		return nil
	})
}

func TestMakeArg(t *testing.T) {
	//LogLevel = LevelDebug
	type test struct{
		args map[string]interface{}
		want []string
	}
	tests := []test{
		{
			args: map[string]interface{}{
				"<cap>": "xa",
				"[--network_id=<network-id>]": "haw",
				"[--grant]": true,
			},
			want: []string{"xa", "--network_id=haw", "--grant"},
		},
		{
			// wp core install
			args: map[string]interface{}{
				"--url=<url>": "http://localhost.com",
				"--title=<site-title>": "testTitle",
				"--admin_user=<username>": "admin",
				"[--admin_password=<password>]": "passwd",
				"--admin_email=<email>": "example@e.com",
				"[--skip-email]": true,
			},
			want: []string{
				"--url=http://localhost.com",
				"--title=testTitle",
				"--admin_user=admin",
				"--admin_password=passwd",
				"--admin_email=example@e.com",
				"--skip-email",
			},
		},
		{
			// wp plugin install
			args: map[string]interface{}{
				"<plugin|zip|url>": []string{"1plugin1", "1plugin2"},
				"[--version=<version>]": "",
				"[--force]": false,
				"[--activate]": true,
				"[--activate-network]": false,
			},
			want: []string{
				"1plugin1",
				"1plugin2",
				"--activate",
			},
		},
	}
	for _, test := range tests {
		var args []string
		for k, v := range test.args {
			args = MakeArg(args, k, v)
		}
		sort.Strings(args)
		sort.Strings(test.want)
		if !reflect.DeepEqual(args, test.want) {
			t.Errorf("%s !should be! %s, len: %v:%v", args, test.want, len(args), len(test.want))
		}
	}
}
