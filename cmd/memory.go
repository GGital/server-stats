/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// memoryCmd represents the memory command
var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "A command to display memory usage",
	Long:  `This command will display the total memory usage of the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		getMemoryUsage()
	},
}

func getMemoryUsage() {
	cmd := exec.Command("sh", "-c", "free -m | awk 'NR==2{printf \"Memory Usage: %.2f%%\", $3*100/$2 }'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(memoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// memoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// memoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
