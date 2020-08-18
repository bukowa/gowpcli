package gowpcli

import (
	"context"
	"fmt"
	"regexp"

	"log"
	"net/http"
	"strings"
	"time"
)

// todo special treatment
//"[--field=<field>]": "[--field=<field>]",
//"[--<field>=<value>]": "[--<field>=<value>]",

func formatKey(k string) string {
	var r = []string{"[", "]", "..."}
	for _, s := range r {
		k = strings.Replace(k, s, "", -1)
	}
	return k
}

func formatMap(m map[string]string) []string {
	var s []string
	for k, v := range m {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return s
}

// formatNamed
func formatNamed(k, v string) string {
	var reg = regexp.MustCompile("(--.*?)\\=")
	param := reg.Find([]byte(k))
	return fmt.Sprintf("%s%s", param, v)
}

// isNamed checks if parameter is in form of `--name=`
func isNamed(s string) (t bool) {
	if strings.Contains(s, "--") {
		t = true
	}
	return
}

func mult(s string) bool {
	return strings.HasSuffix(s, "...") || strings.HasSuffix(s, "...]")
}

func isRequired(s string) bool {
	return strings.HasPrefix(s, "<") || strings.HasPrefix(s, "--") || strings.HasPrefix(s, "[<")
}

func opt(s string) bool {
	return strings.HasPrefix(s, "[")
}


func MakeArg(args []string, k string, v interface{}) []string {
	Logger.Debugf("INPUT:%v, %v, %v", args, k, v)
	k = formatKey(k)
	switch t := v.(type) {
	case map[string]string:
		Logger.Debug("map[string]string")
		if len(t) != 0 {
			args = append(args, formatMap(t)...)
		}
	case []string:
		Logger.Debug("[]string")
		if len(t) != 0 {
			args = append(args, t...)
		}
	case bool:
		Logger.Debug("bool")
		if t {
			args = append(args, k)
		}
	case string:
		Logger.Debug("string")
		if t != "" {
			if isNamed(k) {
				Logger.Debug("is named")
				args = append(args, formatNamed(k, t))
			} else {
				Logger.Debug("is not named")
				args = append(args, t)
			}
		}
	}
	Logger.Debugf("RETURN:%v", args)
	return args
}

func WaitForUrl(url string, status int, duration time.Duration) (ready bool) {

	mainCtx, cancelMain := context.WithTimeout(context.Background(), duration)
	defer cancelMain()

	reqCtx, cancelReq := context.WithTimeout(context.Background(), duration)
	defer cancelReq()

	for {
		select {
		case <- mainCtx.Done():
			log.Printf("main context timeout, context error: %v", mainCtx.Err())
			return
		default:
			client := &http.Client{}
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			req = req.WithContext(reqCtx)

			log.Printf("sending http request to %s", url)
			res, err := client.Do(req)
			log.Printf("received http response from %s, errors: %v", url, err)

			if err == nil {
				log.Printf("status code for %s %v", url, res.StatusCode)
				if res.StatusCode == status {
					ready = true
					log.Printf("%s is ready", req.URL)
					return
				}
			}
			// todo we don't want to spam in our case
			time.Sleep(time.Second)
		}
	}
}
