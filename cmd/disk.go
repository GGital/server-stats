/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "A command to display disk usage",
	Long:  `This command will display the total disk usage of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		getDiskUsage()
	},
}

func getDiskUsage() {
	cmd := exec.Command("sh", "-c", "df -h | grep '^/' | awk '{ print $1, $5 }'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Disk Usage:\n", string(output))
}

func init() {
	rootCmd.AddCommand(diskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
