package main

import (
	"cellachine/components/machine"
	"fmt"
)

func main() {
	m := &machine.Machine{}
	err := m.Install("/home/xie/share/workspace/Golang/src/cellachine/test1.sh", 1)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	err = m.Install("/home/xie/share/workspace/Golang/src/cellachine/test2.sh", 2)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
	m.Input([]int64{1, 2, 1, 1, 2})
	err = m.Exec()
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}