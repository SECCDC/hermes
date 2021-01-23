/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"hermes/hermes"
)

var (
	dbUser string
	dbPass string
	dbAddr string
	dbName string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start hermes webserver",
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("dbPass") == "" {
			fmt.Println("Password not provided. Exiting!")
			os.Exit(2)
		}
		c := hermes.Config{
			DBUser: viper.GetString("dbUser"),
			DBPass: viper.GetString("dbPass"),
			DBAddr: viper.GetString("dbAddr"),
			DBName: viper.GetString("dbName"),
		}
		hermes.Run(c)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&dbUser, "dbUser", "", dbUser, "database username")
	viper.BindPFlag("dbUser", runCmd.PersistentFlags().Lookup("dbUser"))

	runCmd.PersistentFlags().StringVarP(&dbPass, "dbPass", "", dbPass, "database password")
	viper.BindPFlag("dbPass", runCmd.PersistentFlags().Lookup("dbPass"))

	runCmd.PersistentFlags().StringVarP(&dbAddr, "dbAddr", "", dbAddr, "database address")
	viper.BindPFlag("dbAddr", runCmd.PersistentFlags().Lookup("dbAddr"))

	runCmd.PersistentFlags().StringVarP(&dbName, "dbName", "", dbName, "database to use. Will attempt to create one if it does not exist.")
	viper.BindPFlag("dbName", runCmd.PersistentFlags().Lookup("dbName"))

	viper.SetDefault("dbUser", "root")
	viper.SetDefault("dbAddr", "127.0.0.1:3306")
	viper.SetDefault("dbName", "hermes")
}
