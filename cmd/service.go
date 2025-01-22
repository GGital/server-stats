package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service [filename]",
	Short: "Create systemd service for startup",
	Args:  cobra.ExactArgs(1),
	Run:   createSystemdService,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.Flags().StringP("name", "n", "", "Service name (default: filename)")
	serviceCmd.Flags().StringP("description", "d", "", "Service description")
}

func createSystemdService(cmd *cobra.Command, args []string) {
	filename := args[0]
	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	serviceName, _ := cmd.Flags().GetString("name")
	if serviceName == "" {
		serviceName = filepath.Base(filename)
	}

	description, _ := cmd.Flags().GetString("description")
	if description == "" {
		description = fmt.Sprintf("Autostart service for %s", serviceName)
	}

	serviceTemplate := `[Unit]
Description={{.Description}}

[Service]
ExecStart={{.ExecPath}}
Restart=always
User={{.User}}

[Install]
WantedBy=multi-user.target
`

	data := struct {
		Description string
		ExecPath    string
		User        string
	}{
		Description: description,
		ExecPath:    absPath,
		User:        os.Getenv("USER"),
	}

	servicePath := fmt.Sprintf("/etc/systemd/system/%s.service", serviceName)

	// Create service file
	f, err := os.OpenFile(servicePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating service file: %v\n", err)
		return
	}
	defer f.Close()

	tmpl := template.Must(template.New("service").Parse(serviceTemplate))
	if err := tmpl.Execute(f, data); err != nil {
		fmt.Printf("Error writing service file: %v\n", err)
		return
	}

	// Reload systemd and enable service
	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", serviceName).Run()

	fmt.Printf("Service %s created and enabled\n", serviceName)
}
