# Server Stats CLI

## Project Details

**Server Stats CLI** is a lightweight, easy-to-use command-line tool built using the Go programming language and the Cobra framework. It enables Linux server administrators to monitor essential performance metrics quickly. The tool helps to keep track of server health and performance by providing insights into CPU usage, memory usage, disk usage, and processes.

### **Features**
- **CPU Monitoring**: Check total CPU usage percentage.
- **Memory Monitoring**: Check total memory usage percentage.
- **Disk Monitoring**: Display disk usage details for mounted filesystems.
- **Top Processes**: Identify the top 5 processes by CPU and memory usage.
- **Extensible Design**: Built with the Cobra CLI framework, making it easy to add new commands or features.

---

## Getting Started

### **Prerequisites**

To use this tool, ensure the following are installed on your Linux server:

1. **Go (Golang)**: Version 1.16 or higher.
   - [Download Go](https://go.dev/dl/)
2. **Linux Shell Utilities**: Commands like `top`, `free`, `df`, `ps`, and `awk` should be available (commonly pre-installed on most Linux distributions).

### **Installation**

1. Clone this repository:
   ```bash
   git clone https://github.com/GGital/server-stats.git
   cd server-stats
   ```

2. Build the binary:
   ```bash
   go build -o server-stats
   ```

3. Make the binary executable and move it to a directory in your PATH:
   ```bash
   chmod +x server-stats
   sudo mv server-stats /usr/local/bin/
   ```
   OR
   ```bash
   go install
   ```

5. Verify installation:
   ```bash
   server-stats --help
   ```

### **Usage**

- **Listing commands**:
  ```bash
  server-stats --help
  ```

- **CPU Usage**:
  ```bash
  server-stats cpu
  ```
  Displays the total CPU usage as a percentage.


