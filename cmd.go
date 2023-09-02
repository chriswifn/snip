package snip

import (
	"errors"
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/to"
	"github.com/rwxrob/vars"
	"os"
	"path/filepath"
)

var (
	homeDir string = getHomeDir()
	snipDir string = filepath.Join(homeDir, "repos/github.com/chriswifn/dotfiles/snippets")
)

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return string(dirname)
}

func init() {
	Z.Conf.SoftInit()
	Z.Vars.SoftInit()
	Z.Dynamic[`dsnipDir`] = func() string { return snipDir }
}

var Cmd = &Z.Cmd{
	Name:        `snip`,
	Aliases:     []string{},
	Copyright:   `Copyright 2023 Christian Hageloch`,
	Version:     `v0.1.0`,
	License:     `MIT`,
	Source:      `git@github.com:chriswifn/snip.git`,
	Issues:      `github.com/chriswifn/snip/issues`,
	Commands:    []*Z.Cmd{snipCmd, help.Cmd, vars.Cmd, conf.Cmd, initCmd, listCmd},
	Summary:     help.S(_snip),
	Description: help.D(_snip),
}

var initCmd = &Z.Cmd{
	Name:        `init`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_init),
	Description: help.D(_init),
	Call: func(x *Z.Cmd, _ ...string) error {
		val, _ := x.Caller.C(`snipDir`)
		if val == "null" {
			val = snipDir
		}
		x.Caller.Set(`snipDir`, val)
		return nil
	},
}

var snipCmd = &Z.Cmd{
	Name:        `get`,
	Usage:       `[file|stdin, args]`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_get),
	Description: help.D(_get),
	Call: func(x *Z.Cmd, args ...string) error {
		ln := len(args)
		if ln == 0 {
			fmt.Println("")
			return nil
		}
		filename := to.String(args[0])
		if _, err := os.Stat(filepath.Join(snipDir, filename)); errors.Is(err, os.ErrNotExist) {
			fmt.Println(FillIn(args...))
			return nil
		}
		fmt.Println(FillFile(filepath.Join(snipDir, filename), args[1:]...))
		return nil
	},
}

var listCmd = &Z.Cmd{
	Name:        `list`,
	Usage:       `[help]`,
	Commands:    []*Z.Cmd{help.Cmd},
	Summary:     help.S(_list),
	Description: help.D(_list),
	Call: func(x *Z.Cmd, _ ...string) error {
		snips := ListSnip(snipDir)
		for _, snip := range snips {
			fmt.Println(snip)
		}
		return nil
	},
}
