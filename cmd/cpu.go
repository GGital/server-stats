/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "A command to display cpu usage",
	Long:  `This command will display the total CPU usage of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		getCPUUsage()
	},
}

func getCPUUsage() {
	cmd := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu(s)' | awk '{print $2 + $4}'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total CPU Usage: %s%%\n", string(output))
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
