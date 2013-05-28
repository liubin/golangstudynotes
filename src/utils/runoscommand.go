package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	//"os/Process"
)

func main() {
	// API 1. LookPath
	// func LookPath(file string) (string, error)
	path, err := exec.LookPath("id")
	if err != nil {
		log.Fatal("id not found!")
	}
	fmt.Printf("id is available at %s\n", path)

	// API 2. new Cmd struct (via exec.Command)
	// here cmd is a Cmd struct with Reader,Writer,Process etc.
	cmd := exec.Command("id")

	// cmd get also is cmd + arguments .
	// cmd := exec.Command("ps", "-ef")
	// cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out

	// API 3. Cmd.Run
	// API 4. Cmd.Start
	// Cmd.Run will run the command and wait for the command to complete
	// but Cmd.Start will not wait for the command to complete
	// and your can call Cmd.Wait after and only after once 
	// you have called Cmd.Start
	// Please note the different between Run and Start of Cmd.
	// See http://golang.org/pkg/os/exec/#example_Cmd_Start for more info.
	err2 := cmd.Run()
	if err2 != nil {
		log.Fatal(err2)
	} else{
		// API 5. Cmd structs
		// see http://golang.org/pkg/os/exec/#Cmd for more info.
		process := cmd.Process
		fmt.Printf("pid for native os command is %d\n",process.Pid)
	}
	fmt.Printf("command result: %q\n", out.String())
}

