package machine

import (
	"errors"
	"fmt"
	"os/exec"
)

type Machine struct {
	cmdMap       map[int64]string
	instructions []int64
}

func (m *Machine) Install(cmd string, nu int64) error {
	if m.cmdMap == nil {
		m.cmdMap = make(map[int64]string)
	}
	if nu == 0 {
		return errors.New(fmt.Sprintf("0 is invalid\n"))
	}
	//if v, ok := m.cmdMap[nu]; ok {
	//	return errors.New(fmt.Sprintf("%d has been registered %s\n", nu, v))
	//}
	m.cmdMap[nu] = cmd
	return nil
}

func (m *Machine) Exec() error {
	for _, v := range m.instructions {
		err := exec.Command(m.cmdMap[v]).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Machine) Input(ins []int64) {
	m.instructions = ins
}