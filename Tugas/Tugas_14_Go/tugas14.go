package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var (
		output []byte
		err    error
	)

	// Perintah CLI 1
	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("cmd", "/C", "Systeminfo").Output()
	default:
		output, err = exec.Command("bash", "-c", "lscpu").Output()
	}
	if err == nil {
		fmt.Println()
		fmt.Println("System Info :")
		fmt.Println()
		fmt.Println(string(output))
	}

	// Perintah CLI 2
	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("cmd", "/C", "dir").Output()
	default:
		output, err = exec.Command("bash", "-c", "ls").Output()
	}
	if err == nil {
		fmt.Println()
		fmt.Println("Show File :")
		fmt.Println()
		fmt.Println(string(output))
	}

	// Perintah CLI 3
	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("cmd", "/C", "ipconfig").Output()
	default:
		output, err = exec.Command("bash", "-c", "ifconfig").Output()
	}
	if err == nil {
		fmt.Println()
		fmt.Println("Show IP :")
		fmt.Println()
		fmt.Println(string(output))
	}
}
