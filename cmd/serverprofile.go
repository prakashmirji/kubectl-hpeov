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
	"text/tabwriter"

	"github.com/prakashmirji/kubectl-hpeov/oneview"
	"github.com/spf13/cobra"
)

// serverprofileCmd represents the name command
var serverprofileCmd = &cobra.Command{
	Use:   "serverprofile",
	Short: "A subcommand of hpeov cli for getting server profile details",
	Long: `A subcommand of hpeov cli getting server profile details. For example:

	kubectl hpeov serverprofile get --all .`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

// spGetSubCmd represents the sub command of serverprofile command
var spGetSubCmd = &cobra.Command{
	Use:   "get",
	Short: "A subcommand of hpeov serverprofile cli for getting server profile details",
	Long: `A subcommand of hpeov serverprofile cli for getting server profile details. For example:

kubectl hpeov serverprofile get --all
kubectl hpeov serverprofile get --name=<name of server profile>.`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

// spCreateSubCmd represents the sub command of serverprofile command
var spCreateSubCmd = &cobra.Command{
	Use:   "create",
	Short: "A subcommand of hpeov serverprofile cli for creating server profile details",
	Long: `A subcommand of hpeov serverprofile cli for creating server profile details. For example:

kubectl hpeov serverprofile create --file=<json POST payload of server profile>`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

// spDeleteSubCmd represents the sub command of serverprofile command
var spDeleteSubCmd = &cobra.Command{
	Use:   "delete",
	Short: "A subcommand of hpeov serverprofile cli for deleting server profile using name",
	Long: `A subcommand of hpeov serverprofile cli for deleting server profile using name. For example:

kubectl hpeov serverprofile delete --name=<name of server profile>`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serverprofileCmd)
	serverprofileCmd.AddCommand(spGetSubCmd)
	serverprofileCmd.AddCommand(spCreateSubCmd)
	serverprofileCmd.AddCommand(spDeleteSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	spGetSubCmd.Flags().BoolP("all", "a", false, "Get all")
	spGetSubCmd.Flags().StringP("name", "n", "", "Pass name of the server profile")
	spDeleteSubCmd.Flags().StringP("name", "n", "", "Pass name of the server profile")
	spCreateSubCmd.Flags().StringP("file", "f", "", "Pass json payload as a file")
}

func processSpCLI(cmd *cobra.Command) {
	// TODO - validate the flags or args len
	switch cmd.Name() {
	case "get":
		getServerProfileData(cmd)
	case "create":
		createServerProfile(cmd)
	case "delete":
		deleteServerProfile(cmd)
	default:
		cmd.Help()
	}

}

func getServerProfileData(cmd *cobra.Command) {
	allFlag, _ := cmd.Flags().GetBool("all")
	name, _ := cmd.Flags().GetString("name")
	if allFlag {
		spList, err := oneview.GetAllServerProfileDetails()
		if err != nil {
			fmt.Printf("Error while getting server profile list, error:%v\n", err)
		}
		for idx, sp := range spList.Members {
			w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			if idx == 0 {
				fmt.Fprintln(w, "Profile Name\t\t\tStatus\t\t\tCompliance\t\t\tState\t\t\tSerial Number")
			}
			fmt.Fprintln(w, sp.Name+"\t\t\t"+sp.Status+"\t\t\t"+sp.TemplateCompliance+"\t\t\t"+sp.State+"\t\t\t"+fmt.Sprintf("%v", sp.SerialNumber))
			w.Flush()
		}
		return
	} else if name != "" {
		sp, err := oneview.GetServerProfileByName(name)
		if err != nil {
			fmt.Printf("Error while getting server profile details for profile :%s, error:%v", name, err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(w, "Profile Name\t\t\tStatus\t\t\tCompliance\t\t\tState\t\t\tSerial Number")
		fmt.Fprintln(w, sp.Name+"\t\t\t"+sp.Status+"\t\t\t"+sp.TemplateCompliance+"\t\t\t"+sp.State+"\t\t\t"+fmt.Sprintf("%v", sp.SerialNumber))
		w.Flush()
		return
	}

	cmd.Help()
}

func createServerProfile(cmd *cobra.Command) {
	fmt.Printf("This feature is in testing phase, will be updated soon")
}

func deleteServerProfile(cmd *cobra.Command) {
	fmt.Printf("This feature is in testing phase, will be updated soon")
}
