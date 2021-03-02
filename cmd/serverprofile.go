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
	Short: "A subcommand of hpeov cli for operating with server profile",
	Long: `A subcommand of hpeov cli for operating with server profile. For example:

	kubectl hpeov serverprofile get --all
	kubectl hpeov serverprofile get --profilename=<name of server profile> 
	kubectl hpeov serverprofile create --profilename=<name of server profile> --templatename=<name of server template>
	kubectl hpeov serverprofile delete --profilename=<name of server profile>.`,
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
kubectl hpeov serverprofile get --profilename=<name of server profile>.`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

// spCreateSubCmd represents the sub command of serverprofile command
var spCreateSubCmd = &cobra.Command{
	Use:   "create",
	Short: "A subcommand of hpeov serverprofile cli for creating server profile details",
	Long: `A subcommand of hpeov serverprofile cli for creating server profile details. For example:

kubectl hpeov serverprofile create --profilename=<name of server profile> --templatename=<name of server template>.`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

// spDeleteSubCmd represents the sub command of serverprofile command
var spDeleteSubCmd = &cobra.Command{
	Use:   "delete",
	Short: "A subcommand of hpeov serverprofile cli for deleting server profile using name",
	Long: `A subcommand of hpeov serverprofile cli for deleting server profile using name. For example:

kubectl hpeov serverprofile delete --profilename=<name of server profile>`,
	Run: func(cmd *cobra.Command, args []string) {
		processSpCLI(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serverprofileCmd)
	serverprofileCmd.AddCommand(spGetSubCmd)
	serverprofileCmd.AddCommand(spCreateSubCmd)
	serverprofileCmd.AddCommand(spDeleteSubCmd)

	spGetSubCmd.Flags().BoolP("all", "a", false, "Get all")
	spGetSubCmd.Flags().StringP("profilename", "n", "", "Pass name of the server profile")
	spDeleteSubCmd.Flags().StringP("profilename", "n", "", "Pass name of the server profile")
	spCreateSubCmd.Flags().StringP("profilename", "n", "", "Pass name of the server profile")
	spCreateSubCmd.Flags().StringP("templatename", "t", "", "Pass name of the server profile template")
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
	name, _ := cmd.Flags().GetString("profilename")
	if allFlag {
		spList, err := oneview.GetAllServerProfileDetails()
		if err != nil {
			fmt.Printf("Error while getting server profile list, error:%v\n", err)
		}
		if len(spList.Members) < 1 {
			fmt.Printf("No profile resources found \n")
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

	spName, _ := cmd.Flags().GetString("profilename")
	sptName, _ := cmd.Flags().GetString("templatename")

	if spName == "" || sptName == "" {
		cmd.Help()
		return
	}
	err := oneview.CreateServerProfile(sptName, spName)
	if err != nil {
		fmt.Printf("Server profile create failed for: %s, error: %v\n", spName, err)
	}

	return
}

func deleteServerProfile(cmd *cobra.Command) {

	spName, _ := cmd.Flags().GetString("profilename")

	err := oneview.DeleteServerProfile(spName)
	if err != nil {
		fmt.Printf("Server profile delete failed for: %s, error: %v\n", spName, err)
	}

	return
}
