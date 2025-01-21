/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func getCPUUsage() {
	cmd := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu(s)' | awk '{print $2 + $4}'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total CPU Usage: %s%%\n", string(output))
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

func getMemoryUsage() {
	cmd := exec.Command("sh", "-c", "free -m | awk 'NR==2{printf \"Memory Usage: %.2f%%\", $3*100/$2 }'")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))
}

// resourceCmd represents the resource command
var resourceCmd = &cobra.Command{
	Use:   "resource [cpu|memory|disk]",
	Short: "A command to display the resource usage",
	Args:  cobra.ExactArgs(1),
	Long: `This command will display the resource usage. 
	Available parameters:
	cpu: Display CPU usage
	memory: Display memory usage
	disk: Display disk usage
	For example:
	server-stats resource cpu`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "cpu":
			getCPUUsage()
		case "memory":
			getMemoryUsage()
		case "disk":
			getDiskUsage()
		}
	},
}

func init() {
	rootCmd.AddCommand(resourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
