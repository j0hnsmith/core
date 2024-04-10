// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cli generates powerful CLIs from Go struct types and functions.
// See package cliview to create a GUI representation of a CLI.
package cli

//go:generate core generate

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"cogentcore.org/core/xlog"
)

// Run runs an app with the given options, configuration struct,
// and commands. It does not run the GUI; see [cogentcore.org/core/cliview.Run]
// for that. The configuration struct should be passed as a pointer, and
// configuration options should be defined as fields on the configuration
// struct. The commands can be specified as either functions or struct
// objects; the functions are more concise but require using gti
// (see https://cogentcore.org/core/gti). In addition to the given commands,
// Run adds a "help" command that prints the result of [Usage], which
// will also be the root command if no other root command is specified.
// Also, it adds the fields in [MetaConfig] as configuration options.
// If [Options.Fatal] is set to true, the error result of Run does
// not need to be handled. Run uses [os.Args] for its arguments.
func Run[T any, C CmdOrFunc[T]](opts *Options, cfg T, cmds ...C) error {
	cs, err := CmdsFromCmdOrFuncs[T, C](cmds)
	if err != nil {
		err := fmt.Errorf("internal/programmer error: error getting commands from given commands: %w", err)
		if opts.Fatal {
			xlog.PrintError(err)
			os.Exit(1)
		}
		return err
	}
	cmd, err := Config(opts, cfg, cs...)
	if err != nil {
		if opts.Fatal {
			xlog.PrintlnError("error: ", err)
			os.Exit(1)
		}
		return err
	}
	err = RunCmd(opts, cfg, cmd, cs...)
	if err != nil {
		if opts.Fatal {
			fmt.Println(xlog.CmdColor(cmdString(opts, cmd)) + xlog.ErrorColor(" failed: "+err.Error()))
			os.Exit(1)
		}
		return fmt.Errorf("%s failed: %w", CmdName()+" "+cmd, err)
	}
	// if the user sets level to error (via -q), we don't show the success message
	if opts.PrintSuccess && xlog.UserLevel <= slog.LevelWarn {
		fmt.Println(xlog.CmdColor(cmdString(opts, cmd)) + xlog.SuccessColor(" succeeded"))
	}
	return nil
}

// RunCmd runs the command with the given name using the given options,
// configuration information, and available commands. If the given
// command name is "", it runs the root command.
func RunCmd[T any](opts *Options, cfg T, cmd string, cmds ...*Cmd[T]) error {
	for _, c := range cmds {
		if c.Name == cmd || c.Root && cmd == "" {
			err := c.Func(cfg)
			if err != nil {
				return err
			}
			return nil
		}
	}
	if cmd == "" { // if we couldn't find the command and we are looking for the root command, we fall back on help
		fmt.Println(Usage(opts, cfg, cmd, cmds...))
		os.Exit(0)
	}
	return fmt.Errorf("command %q not found", cmd)
}

// CmdName returns the name of the command currently being run.
func CmdName() string {
	base := filepath.Base(os.Args[0])
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// cmdString is a simple helper function that returns a string
// with [CmdName] and the given command name string combined
// to form a string representing the complete command being run.
func cmdString(opts *Options, cmd string) string {
	if cmd == "" {
		return CmdName()
	}
	return CmdName() + " " + cmd
}
