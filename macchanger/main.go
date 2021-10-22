package main

import (
	"flag"
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
	cmdo.Stdin = os.Stdin

	err = cmdo.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	return nil

}

func main() {
	iface := flag.String("iface", "wlan0", "interface to change MAC on")
	mac := flag.String("mac", "", "desired MAC address")
	flag.Parse()
	cmd := "sudo"
	ifc := "ifconfig"
	ecmd(cmd, []string{ifc, *iface, "down"})
	ecmd(cmd, []string{ifc, *iface, "hw", "ether", *mac})
	ecmd(cmd, []string{ifc, *iface, "up"})
}
