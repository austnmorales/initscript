package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	go nmap_scan()
	go nikto_scan()
}

func nmap_scan() {
	out, err := exec.Command("sudo", "nmap", "-A", "-T4", string(os.Args[1])).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func nikto_scan() {
	out, err := exec.Command("sudo", "nikto", "-host", string(os.Args[1])).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
