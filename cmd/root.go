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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hpeov",
	Short: "A kubectl extension for hpe bare metal servers",
	Long: `A kubectl extension for hpe bare metal servers. For example:

	kubectl hpeov serverhardware get --all
    kubectl hpeov serverhardware get --name=<name of server hardware> 
    kubectl hpeov serverhardware power --name=<server name> --powerstate=On
    kubectl hpeov serverprofile get --all
    kubectl hpeov serverprofile get --profilename=<name of server profile> 
    kubectl hpeov serverprofile create --profilename=<name of server profile> --templatename=<name of server template>
    kubectl hpeov serverprofile delete --profilename=<name of server profile>
    kubectl hpeov servertemplate get --all
    kubectl hpeov servertemplate get --name=<templa name>.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubectl-hpeov.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".kubectl-hpeov" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kubectl-hpeov")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
