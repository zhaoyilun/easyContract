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
	"github.com/spf13/cobra"
	"github.com/zhaoyilun/easyContract/utils"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an EasyContract file in this directory.",
	Long: `Create an EasyContract file in this directory, the file will end with .ez.
You must use the "-f" or "--filename" tag to enter the file name of the file you created, 
Otherwise a default file named default.ez will be created.`,
	Run: func(cmd *cobra.Command, args []string) {
		//get flags
		title := cmd.Flag("filename").Value.String()
		if len(title) == 0 {
			title = "default.ez"
		}
		//create the file
		utils.CreateFile(title)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.Flags().StringP("filename", "f", "", "the new file's name")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
