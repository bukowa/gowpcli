package wpcli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Command interface {
	Args() []string
}

type ClientInterface interface {
	Execute(Command) error
	Config() *ClientConfig
	Core() *Core
	Plugin() *Plugin
}

type ClientConfig struct {
	ExecutablePath string // path to wp-cli

	Path string // Path to the WordPress files.
	Ssh  string // Perform operation against a remote server over SSH.

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// shortcut
func NewConfig(executablePath, wordpressPath string) *ClientConfig {
	if executablePath == "" {
		executablePath = "wp"
	}
	return &ClientConfig{
		ExecutablePath: executablePath,
		Path: wordpressPath,
	}
}

type Client struct {
	config *ClientConfig
	core   *Core
	plugin *Plugin
}

// shortcut
func NewClient(config *ClientConfig, osStd bool) *Client {
	if osStd {
		config.Stdin = os.Stdin
		config.Stdout = os.Stdout
		config.Stderr = os.Stderr
	}
	c := &Client{config: config}

	c.core = &Core{}
	c.plugin = &Plugin{}
	return c
}

func (c *Client) Execute(command Command) error {
	cmd := exec.Command(c.config.ExecutablePath, command.Args()...)
	fmt.Println(cmd.Path, cmd.Args)
	if c.Config().Stdin != nil {
		cmd.Stdin = c.Config().Stdin
	}
	if c.Config().Stdout != nil {
		cmd.Stdout = c.Config().Stdout
	}
	if c.Config().Stderr != nil {
		cmd.Stderr = c.Config().Stderr
	}
	return cmd.Run()
}

func (c *Client) Config() *ClientConfig {
	return c.config
}

func (c *Client) Core() *Core {
	return c.core
}

func (c *Client) Plugin() *Plugin {
	return c.plugin
}

