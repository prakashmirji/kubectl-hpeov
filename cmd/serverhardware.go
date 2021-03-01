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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/prakashmirji/kubectl-hpeov/oneview"
	"github.com/spf13/cobra"
)

// serverhardwareCmd represents the name command
var serverhardwareCmd = &cobra.Command{
	Use:   "serverhardware",
	Short: "A subcommand of hpeov cli for getting server hardware details",
	Long: `A subcommand of hpeov cli getting server hardware details. For example:

	kubectl hpeov serverhardware get all .`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("name called")
		fstatus, _ := cmd.Flags().GetBool("all")
		fmt.Println("all :", fstatus)
		fstr, _ := cmd.Flags().GetString("name")
		fmt.Println("fstr :", fstr)
		processCLI(cmd)
	},
}

// shGetSubCmd represents the sub command of serverhardware command
var shGetSubCmd = &cobra.Command{
	Use:   "get",
	Short: "A subcommand of hpeov serverhardware cli for getting server hardware details",
	Long: `A subcommand of hpeov serverhardware cli for getting server hardware details. For example:

kubectl hpeov serverhardware get all.`,
	Run: func(cmd *cobra.Command, args []string) {
		processCLI(cmd)
	},
}

// shPowerSubCmd represents the sub command of serverhardware command
var shPowerSubCmd = &cobra.Command{
	Use:   "power",
	Short: "A subcommand of hpeov serverhardware cli for powering ON/OFF server hardware details",
	Long: `A subcommand of hpeov serverhardware cli for powering ON/OFF server hardware details. For example:

kubectl hpeov serverhardware power --name <server hardware name> --powerstatus=ON.`,
	Run: func(cmd *cobra.Command, args []string) {
		processCLI(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serverhardwareCmd)
	serverhardwareCmd.AddCommand(shGetSubCmd)
	serverhardwareCmd.AddCommand(shPowerSubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	shGetSubCmd.Flags().BoolP("all", "a", false, "Get all")
	shGetSubCmd.Flags().StringP("name", "n", "", "Pass name of the server hardware")
	shPowerSubCmd.Flags().StringP("name", "n", "", "Pass name of the server hardware")
	shPowerSubCmd.Flags().StringP("powerstate", "p", "", "Pass value as On or Off")
}

func processCLI(cmd *cobra.Command) {
	// TODO - validate the flags or args len
	switch cmd.Name() {
	case "get":
		getServerHardwareData(cmd)
	case "power":
		updateServerPowerState(cmd)
	}

}

func getServerHardwareData(cmd *cobra.Command) {
	//fmt.Println("Server hardware list")
	allFlag, _ := cmd.Flags().GetBool("all")
	name, _ := cmd.Flags().GetString("name")
	if allFlag {
		serverList, err := oneview.GetAllServerHardwareDetails()
		if err != nil {
			fmt.Printf("Error while getting server list, error:%v\n", err)
		}
		for idx, svr := range serverList.Members {
			w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			if idx == 0 {
				fmt.Fprintln(w, "Server Name\tPower State\tModel\tMemory\tStatus\tiLO Address")
			}
			fmt.Fprintln(w, svr.Name+"\t"+svr.PowerState+"\t"+svr.ShortModel+"\t"+fmt.Sprintf("%d", svr.MemoryMb)+"\t"+svr.Status+"\t"+svr.GetIloIPAddress())
			w.Flush()
		}
		return
	} else if name != "" {
		svr, err := oneview.GetServerHardwareByName(name)
		if err != nil {
			fmt.Printf("Error while getting server hardware details for :%s, error:%v", name, err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(w, "Server Name\tPower State\tModel\tMemory\tStatus\tiLO Address")
		fmt.Fprintln(w, svr.Name+"\t"+svr.PowerState+"\t"+svr.ShortModel+"\t"+fmt.Sprintf("%d", svr.MemoryMb)+"\t"+svr.Status+"\t"+svr.GetIloIPAddress())
		w.Flush()
		return
	}

	cmd.Help()
}

func updateServerPowerState(cmd *cobra.Command) {
	name, _ := cmd.Flags().GetString("name")
	powerState, _ := cmd.Flags().GetString("powerstate")
	if name == "" || powerState == "" {
		cmd.Help()
		return
	}
	if err := oneview.UpdatePowerState(name, powerState); err != nil {
		fmt.Printf("Failed to update power state for server: %s, error: %v\n", name, err)
	}
	fmt.Printf("Server :%s power state changed succesfully to :%s\n", name, powerState)
}

func hexToName(args []string) {
	var hexMap map[string]string

	// read the color.min.json
	content, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading the file  %v", err)
	}

	_ = json.Unmarshal(content, &hexMap)

	name, ok := hexMap[args[0]]

	if ok {
		fmt.Printf("Name: %s, Hex: %s\n", name, args[0])
	} else {
		fmt.Printf("color name not found\n")
	}
}
