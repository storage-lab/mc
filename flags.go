/*
 * Minio Client (C) 2014, 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import "github.com/minio/cli"

// Collection of mc commands currently supported are
var commands = []cli.Command{}

// Collection of mc flags currently supported
var flags = []cli.Flag{}

var (
	configFlag = cli.StringFlag{
		Name:  "config-folder, C",
		Value: mustGetMcConfigDir(),
		Usage: "Path to configuration folder",
	}

	quietFlag = cli.BoolFlag{
		Name:  "quiet, q",
		Usage: "Suppress chatty console output",
	}

	mimicFlag = cli.BoolFlag{
		Name:  "mimic",
		Usage: "Mimic operating system toolchain behavior wherever it makes sense",
	}

	jsonFlag = cli.BoolFlag{
		Name:  "json",
		Usage: "Enable json formatted output",
	}

	debugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "Enable debugging output",
	}

	// Add your new flags starting here
)

// registerCmd registers a cli command
func registerCmd(cmd cli.Command) {
	commands = append(commands, cmd)
}

// registerFlag registers a cli flag
func registerFlag(flag cli.Flag) {
	flags = append(flags, flag)
}