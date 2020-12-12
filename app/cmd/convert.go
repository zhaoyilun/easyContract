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

// convertEOSCmd represents the convert command
var convertEosCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the contract code into EOSIO-specific code",
	Long: `Convert the contract code into EOSIO-specific code.

Here you need to use the "-o" or "--output" tag to name the generated code. 
	
The generated object code file will be in the "/EOS/" directory under this level directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert eos called")
	},
}

// convertETHCmd represents the convert command
var convertETHCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the contract code into ETH-specific code",
	Long: `Convert the contract code into ETH-specific code.

Here you need to use the "-o" or "--output" tag to name the generated code. 

The generated object code file will be in the "/ETH/" directory under this level directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert eth called")
	},
}

// convertBOSCmd represents the convert command
var convertBosCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert the contract code into BOS-specific code",
	Long: `Convert the contract code into BOS-specific code.

Here you need to use the "-o" or "--output" tag to name the generated code. 

The generated object code file will be in the "/BOS/" directory under this level directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert bos called")
	},
}

func init() {
	eosCmd.AddCommand(convertEosCmd)
	ethCmd.AddCommand(convertETHCmd)
	bosCmd.AddCommand(convertBosCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
