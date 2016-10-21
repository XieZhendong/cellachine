package client

import (
	"net/rpc"
	"fmt"
	"cellachine/paramanager/server"
	"cellachine/paramanager/common"
)

type ParaGetter struct {
	*rpc.Client
}

func (p *ParaGetter) SetInt(name string, value int) error {
	var ok bool
	err := p.Call(common.REGISTERNAME + ".SetInt", &server.IntType{name,value}, &ok)
	if !ok {
		return err
	}
}

func Init() error {
	cli, err := rpc.DialHTTP(common.PROTOCOL, common.ADDR)
	if err != nil {
		return err
	}
	var strs bool
	err = cli.Call(common.REGISTERNAME + ".SetInt", &server.IntType{"test",1}, &strs)
	if err != nil {
		return err
	}
	fmt.Printf("reply: %v\n", strs)
	return nil
}