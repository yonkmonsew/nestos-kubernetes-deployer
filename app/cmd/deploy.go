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
	"nestos-kubernetes-deployer/app/phases/infra"

	"github.com/spf13/cobra"
)

func NewDeployCommand() *cobra.Command {
	var dir, role string

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Use this command to deploy kubernetes node",
		RunE: func(cmd *cobra.Command, args []string) error {
			cluster := &infra.Cluster{
				Dir:  dir,
				Role: role,
			}
			return cluster.Create()
		},
	}

	cmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "directory for deployment")
	cmd.PersistentFlags().StringVarP(&role, "role", "r", "", "node role for deployment")

	return cmd
}