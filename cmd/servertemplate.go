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

// servertemplateCmd represents the name command
var servertemplateCmd = &cobra.Command{
	Use:   "servertemplate",
	Short: "A subcommand of hpeov cli for getting server template details",
	Long: `A subcommand of hpeov cli getting server template details. For example:

	kubectl hpeov servertemplate get --all 
	kubectl hpeov servertemplate get --name=<name of sever template> .`,
	Run: func(cmd *cobra.Command, args []string) {
		processSptCLI(cmd)
	},
}

// sptGetSubCmd represents the sub command of servertemplate command
var sptGetSubCmd = &cobra.Command{
	Use:   "get",
	Short: "A subcommand of hpeov servertemplate cli for getting server template details",
	Long: `A subcommand of hpeov servertemplate cli for getting server template details. For example:

kubectl hpeov serverprofile get --all
kubectl hpeov serverprofile get --name=<name of server template>.`,
	Run: func(cmd *cobra.Command, args []string) {
		processSptCLI(cmd)
	},
}

func init() {
	rootCmd.AddCommand(servertemplateCmd)
	servertemplateCmd.AddCommand(sptGetSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sptGetSubCmd.Flags().BoolP("all", "a", false, "Get all")
	sptGetSubCmd.Flags().StringP("name", "n", "", "Pass name of the server template")
}

func processSptCLI(cmd *cobra.Command) {
	// TODO - validate the flags or args len
	switch cmd.Name() {
	case "get":
		getServerTemplateData(cmd)
	default:
		cmd.Help()
	}

}

func getServerTemplateData(cmd *cobra.Command) {
	allFlag, _ := cmd.Flags().GetBool("all")
	name, _ := cmd.Flags().GetString("name")
	if allFlag {
		sptList, err := oneview.GetAllServerTemplateDetails()
		if err != nil {
			fmt.Printf("Error while getting server template list, error:%v\n", err)
		}
		if len(sptList.Members) < 1 {
			fmt.Printf("No template resources found \n")
		}
		for idx, spt := range sptList.Members {
			w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			if idx == 0 {
				fmt.Fprintln(w, "Template Name\t\t\tStatus\t\t\tCompliance")
			}
			fmt.Fprintln(w, spt.Name+"\t\t\t"+spt.Status+"\t\t\t"+spt.TemplateCompliance)
			w.Flush()
		}
		return
	} else if name != "" {
		spt, err := oneview.GetServerTemplateByName(name)
		if err != nil {
			fmt.Printf("Error while getting server template details for template :%s, error:%v", name, err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(w, "Template Name\t\t\tStatus\t\t\tCompliance")
		fmt.Fprintln(w, spt.Name+"\t\t\t"+spt.Status+"\t\t\t"+spt.TemplateCompliance)
		w.Flush()
		return
	}

	cmd.Help()
}
