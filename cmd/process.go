/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

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

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process [cpu|memory]",
	Short: "A command to show top CPU or memory consuming processes",
	Args:  cobra.ExactArgs(1),
	Long: `This command is used to display the top CPU or memory consuming processes. 
	Available flags:
	-n, --number: Number of processes to display
	Parameters:
	cpu: Display top CPU consuming processes
	memory: Display top memory consuming processes
	For example:
	server-stats process cpu -n 10`,
	Run: func(cmd *cobra.Command, args []string) {
		n := cmd.Flag("number").Value.String()
		switch args[0] {
		case "cpu":
			getTopCPUProcesses(n)
		case "memory":
			getTopMemoryProcesses(n)
		}
	},
}

func init() {
	rootCmd.AddCommand(processCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	processCmd.Flags().IntP("number", "n", 5, "Number of processes to display")
}
