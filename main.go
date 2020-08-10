package main

import (
	"flag"
	"github.com/bukowa/gowpcli/pkg/wpcli/utils"
	"log"
	"os"
	"os/exec"
	"time"
)

var WaitURL = flag.String("wait-url", "", "waits for url to be ready ")
var WaitStatus = flag.Int("wait-status", 200, "waits for this status to be ready")
var WaitTimeout = flag.Duration("wait-timeout", time.Second*20, "max wait time")
var WordpressPath = flag.String("path", "/var/www/html", "path to wordpress installation")

func main() {
	flag.Parse()
	time.Sleep(time.Second)
	if *WaitURL != "" {
		log.Printf("Waiting for %s to be ready...", *WaitURL)
		utils.WaitForUrl(*WaitURL, *WaitStatus, *WaitTimeout)
	}
	cmd := exec.Command("wp", "core", "install")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println(cmd.Run())
	//
	//cfg := wpcli.NewConfig("wp", *WordpressPath)
	//c := wpcli.NewClient(cfg, true)
	//
	//c.Execute(c.Core().Install(core.Install{
	//	Url:           "http://localhost:8090",
	//	Title:         "test",
	//	AdminUser:     "test",
	//	AdminPassword: "test",
	//	AdminEmail: 	"test@example.com",
	//	SkipEmail:     true,
	//}))
}
