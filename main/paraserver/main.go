package main

import (
	"cellachine/paramanager/server"
	"fmt"
)

func main() {
	err := server.Init()
	if err != nil {
		fmt.Printf("server init error: %v\n", err)
		return
	}
}
