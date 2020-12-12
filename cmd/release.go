/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
)

// releaseEOSCmd represents the release command
var releaseEosCmd = &cobra.Command{
	Use:   "release",
	Short: "Compile the locally successfully converted code into an executable file for the target platform",
	Long: `Compile the locally successfully converted code into an executable file for the target platform.

Please make sure that the corresponding compilation environment has been installed on your computer, and the environment variables have been configured.

Please switch to the project root directory or EOS directory when using this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release eos called")
	},
}

// releaseETHCmd represents the release command
var releaseETHCmd = &cobra.Command{
	Use:   "release",
	Short: "Compile the locally successfully converted code into an executable file for the target platform",
	Long: `Compile the locally successfully converted code into an executable file for the target platform.

Please make sure that the corresponding compilation environment has been installed on your computer, and the environment variables have been configured.

Please switch to the project root directory or ETH directory when using this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release eth called")
	},
}

// releaseBOSCmd represents the release command
var releaseBosCmd = &cobra.Command{
	Use:   "release",
	Short: "Compile the locally successfully converted code into an executable file for the target platform",
	Long: `Compile the locally successfully converted code into an executable file for the target platform.

Please make sure that the corresponding compilation environment has been installed on your computer, and the environment variables have been configured.

Please switch to the project root directory or BOS directory when using this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("release bos called")
	},
}

func init() {
	eosCmd.AddCommand(releaseEosCmd)
	ethCmd.AddCommand(releaseETHCmd)
	bosCmd.AddCommand(releaseBosCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// releaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// releaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
