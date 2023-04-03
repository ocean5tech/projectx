package main

import (
	"time"

	"github.com/ocean5tech/projectx/network"
)

func main() {
	//trLocal := network.NewLocalTransport("LOCAL")
	//trRemote := network.NewLocalTransport("REMOTE")
	aRemote := network.NewLocalTransport("A") //注册
	bRemote := network.NewLocalTransport("B") //注册
	cRemote := network.NewLocalTransport("C") //不注册
	dRemote := network.NewLocalTransport("D") //不注册

	// trLocal.Connect(trRemote)
	// trRemote.Connect(trLocal)

	aRemote.Connect(bRemote) // A-B,都在Server中注册
	bRemote.Connect(aRemote)
	aRemote.Connect(cRemote) // A-C, A注册，C不注册
	cRemote.Connect(aRemote)
	cRemote.Connect(dRemote) // C-D, C不注册，d不注册
	dRemote.Connect(cRemote)
	bRemote.Connect(dRemote) // B-> D 单项


	go func() {
		for {
			aRemote.SendMessage(bRemote.Addr(), []byte("A say Hello to B"))
			bRemote.SendMessage(aRemote.Addr(), []byte("B say Hello to A"))
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			aRemote.SendMessage(cRemote.Addr(), []byte("A say Hello to C"))
			cRemote.SendMessage(aRemote.Addr(), []byte("C say Hello to A"))
			time.Sleep(1 * time.Second)
		}
	}()	
	go func() {
		for {
			cRemote.SendMessage(dRemote.Addr(), []byte("C say Hello to D"))
			dRemote.SendMessage(cRemote.Addr(), []byte("D say Hello to C"))
			time.Sleep(1 * time.Second)
		}
	}()	
	go func() {
		for {
			bRemote.SendMessage(dRemote.Addr(), []byte("B say Hello to D"))
			dRemote.SendMessage(bRemote.Addr(), []byte("D say Hello to B"))
			time.Sleep(1 * time.Second)
		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{aRemote, bRemote},
	}

	s := network.NewServer(opts)
	s.Start()
}
