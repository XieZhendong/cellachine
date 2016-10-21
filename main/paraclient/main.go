package main

import (
	"cellachine/paramanager/client"
	"fmt"
)

func main() {
	err := client.Init()
	if err != nil {
		fmt.Printf("init error: %v\n", err)
		return
	}
}