/*
Copyright 2023 KylinSoft  Co., Ltd.

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

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	var (
		version = "0.2.2"
		arch    = fmt.Sprint(runtime.GOOS, "/", runtime.GOARCH)
	)

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display the NKD version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version:    %s\n", version)
			fmt.Printf("OS/Arch:    %s\n", arch)
		},
	}

	return cmd
}
