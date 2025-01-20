/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// topMemCmd represents the topMem command
var topMemCmd = &cobra.Command{
	Use:   "topMem",
	Short: "A command to display the top memory consuming processes",
	Long: `This command will display the top memory consuming processes.
	Available flags:
	-n, --number: Number of processes to display
	For example:
	server-stats topMem -n 10`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cmd.Flag("number").Value.String()
		getTopMemoryProcesses(n)
	},
}

func getTopMemoryProcesses(n string) {
	command := "ps -eo pid,comm,%mem --sort=-%mem | head -n " + n
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Top 5 Processes by Memory Usage:\n", string(output))
}

func init() {
	rootCmd.AddCommand(topMemCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topMemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topMemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	topMemCmd.Flags().IntP("number", "n", 5, "Number of processes to display")
}
