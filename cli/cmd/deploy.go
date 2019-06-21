/*
Copyright 2019 Cortex Labs, Inc.

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

	"github.com/spf13/cobra"

	"github.com/cortexlabs/cortex/pkg/config"
	"github.com/cortexlabs/cortex/pkg/lib/debug"
	"github.com/cortexlabs/cortex/pkg/lib/errors"
)

var flagDeployForce bool

func init() {
	deployCmd.PersistentFlags().BoolVarP(&flagDeployForce, "force", "f", false, "stop all running jobs")
	addEnvFlag(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy APIs",
	Long:  "Deploy APIs.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		deploy(flagDeployForce, false)
	},
}

func deploy(force bool, ignoreCache bool) {
	root := mustCortexRoot()

	configPaths := yamlPaths(root)
	cfg, err := config.NewFromFiles(configPaths)
	if err != nil {
		errors.Exit(err)
	}

	fmt.Println(root)
	debug.Pp(cfg)
}
