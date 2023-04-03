package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.peers[trb.Addr()], trb)
	assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)
}

func TestSendMessageServer(t *testing.T){
	aRemote := network.NewLocalTransport("A") //注册
	bRemote := network.NewLocalTransport("B") //注册
	cRemote := network.NewLocalTransport("C") //不注册
	dRemote := network.NewLocalTransport("D") //不注册
	aRemote.Connect(bRemote) // A-B,都在Server中注册
	bRemote.Connect(aRemote)
	aRemote.Connect(cRemote) // A-C, A注册，C不注册
	cRemote.Connect(aRemote)
	cRemote.Connect(dRemote) // C-D, C不注册，d不注册
	dRemote.Connect(cRemote)
	bRemote.Connect(dRemote) // B-> D 单项

	
}