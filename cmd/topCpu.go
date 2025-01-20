/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// topCpuCmd represents the topCpu command
var topCpuCmd = &cobra.Command{
	Use:   "topCpu",
	Short: "A command to display the top CPU consuming processes",
	Long: `This command will display the top CPU consuming processes.
	Available flags:
	-n, --number: Number of processes to display
	For example:
	server-stats topCpu -n 10`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cmd.Flag("number").Value.String()
		getTopCPUProcesses(n)
	},
}

func getTopCPUProcesses(n string) {
	command := "ps -eo pid,comm,%cpu --sort=-%cpu | head -n " + n
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Top 5 Processes by CPU Usage:\n", string(output))
}

func init() {
	rootCmd.AddCommand(topCpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topCpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topCpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	topCpuCmd.Flags().IntP("number", "n", 5, "Number of processes to display")
}
