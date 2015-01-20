package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var ncommits int
var file_name string

func init() {
	flag.IntVar(&ncommits, "n", 10, "Number of commits to perform")
	flag.StringVar(&file_name, "f", "", "The file to commit to")
}

func main() {
	flag.Parse()

	_, err := os.Stat(".git")
	if os.IsNotExist(err) {
		log.Fatal("Directory is not a git repository.\n")
	}

	_, err = exec.LookPath("git")
	if err != nil {
		log.Fatal("Git is not installed\n")
	}

	for i := 1; i <= ncommits; i++ {
		file, err := os.OpenFile(file_name, os.O_RDWR, 0666)
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.\n")
		}
		file.Seek(0, 2)
		file.WriteString(fmt.Sprint("Edit number ", i, "\n"))
		file.Close()

		cmd := exec.Command("git", "add", file_name)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		cmd = exec.Command("git", "commit", "-m", fmt.Sprint("Commit number ", i))
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		progress := (float32(i) / float32(ncommits)) * 100
		fmt.Printf("\rProgress: %6.2f%%", progress)
	}
	fmt.Printf("\n")
}
