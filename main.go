package main

import (
    "fmt"
    "os"
    
    "github.com/t0mk/gometal/client"
)

func main() {
    command := client.NewCommandWithDefaultClient()
    if _, err := command.ExecuteC(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(-1)
	}
}}
