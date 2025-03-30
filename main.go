package main

import (
	"fmt"
	"log"
	"github.com/samriddhi34/GoDrive/p2p"
)

func main(){
	tcpOps := p2p.TCPTransportOPS{
		ListenerAddr: ":3000",
		HandshakeFunc: p2p.NOPHandShakeFunc,
		Decoder: p2p.DefaultDecoder{},

	}
	tr := p2p.NewTCPTransport(tcpOps)

	go func(){
		for{
			msg:= <- tr.Consume()
			fmt.Printf("%+v\n",msg)
		}
	}()

	if err := tr.ListenAndAccept() ; err != nil{
		log.Fatal(err)
	}

	
}