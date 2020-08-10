package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func MakeArg(args []string, key, value string) []string {
	return append(args, fmt.Sprintf("--%s=%s", key, value))
}

func WaitForUrl(url string, status int, duration time.Duration) (ready bool) {

	main_ctx, cancel_main := context.WithTimeout(context.Background(), duration)
	defer cancel_main()

	req_ctx, cancel_req := context.WithTimeout(context.Background(), duration)
	defer cancel_req()

	for {
		select {
		case <- main_ctx.Done():
			log.Printf("main context timeout, context error: %v", main_ctx.Err())
			return
		default:
			client := &http.Client{}
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			req = req.WithContext(req_ctx)

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

