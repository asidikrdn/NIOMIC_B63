package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// Format menjalankan sebuah perintah dan menyimpan hasilnya ke sebuah variabel => varHasil, varError := exec.Command(perintahYangDiinginkan).Output()
	// Jika perintahnya lebih dari 1 kata, maka dipisah-pisah tiap katanya

	// menjalankan perintah "ls"
	showFiles, _ := exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(showFiles))

	// menjalankan perintah "pwd"
	showPwd, _ := exec.Command("pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(showPwd))

	// menjalankan perintah "git status"
	showGitStatus, _ := exec.Command("git", "status").Output()
	fmt.Printf(" -> git status\n%s\n", string(showGitStatus))

	// untuk mengatasi permasalahan perbedaan perintah pada beberapa OS, bisa menggunakan cara berikut
	var (
		output []byte
		err    error
	)

	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("cmd", "/C", "git log").Output()
	default:
		output, err = exec.Command("bash", "-c", "git log").Output()
	}

	if err == nil {
		fmt.Println()
		fmt.Println("Hasil Git Log :")
		fmt.Println()
		fmt.Println(string(output))
	}
}
