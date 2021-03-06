/*
Copyright (c) 2016-2017 Bitnami

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// VERSION will be overwritten automatically by the build system
	VERSION = "devel"

	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
)

//VersionCmd contains first-class command for version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Kubeless",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kubeless version: " + VERSION + " (" + GITCOMMIT + ")")
	},
}
