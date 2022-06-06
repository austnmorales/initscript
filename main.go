package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	nmap_scan()
}

func nmap_scan() {
	//ip := os.Args[1]
	out, err := exec.Command("sudo", "nmap", string(os.Args[1])).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
