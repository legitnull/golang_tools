package main

import (
	"log"
	"os"
	"os/exec"
)

func ecmd(cmd string, ara []string) (err error) {
	args := ara
	cmdo := exec.Command(cmd, args...)
	// cmd out & disp
	cmdo.Stdout = os.Stdout
	cmdo.Stderr = os.Stderr

	err = cmdo.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil

}

func main() {
	cmd := "cowsay"
	ecmd(cmd, []string{"FUCKHOLEJONES"})
}
