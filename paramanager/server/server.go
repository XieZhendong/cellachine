package server

import (
	"net/rpc"
	"net/http"
	"net"
	"fmt"
	"github.com/pkg/errors"
	"cellachine/paramanager/common"
	"sync"
)

type IntType struct {
	Name string
	Value int
}

type intWithLock struct {
	*sync.RWMutex
	value int
}

type Para struct {
	intValues map[string]*intWithLock
}

func (p *Para) Test(args int, reply *string) error {
	fmt.Printf("Test: %d\n", args)
	*reply = "test"
	return nil
}

func (p *Para) SetInt(args *IntType, reply *bool) error {
	if _, ok := p.intValues[args.Name]; ok {
		*reply = false
		return errors.New(common.ERR_REGISTERED)
	}
	*reply = true
	p.intValues[args.Name] = &intWithLock{&sync.RWMutex{}, args.Value}
	return nil
}

func (p *Para) GetInt(name string, reply *int) error {
	if _, ok := p.intValues[name]; !ok {
		return errors.New(common.ERR_NONE)
	}
	p.intValues[name].RLock()
	*reply = p.intValues[name].value
	p.intValues[name].RUnlock()
	return nil
}

func (p *Para) UpdateInt(args *IntType, reply *int) error {
	if _, ok := p.intValues[args.Name]; !ok {
		return errors.New(common.ERR_NONE)
	}
	p.intValues[args.Name].Lock()
	p.intValues[args.Name].value = args.Value
	p.intValues[args.Name].Unlock()
	return nil
}

func Init() error {
	p := new(Para)
	p.intValues = make(map[string]*intWithLock)
	rpc.RegisterName(common.REGISTERNAME, p)
	rpc.HandleHTTP()
	l, err := net.Listen(common.PROTOCOL, common.ADDR)
	if err != nil {
		return err
	}
	http.Serve(l, nil)
	return nil
}