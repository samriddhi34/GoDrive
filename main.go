package main

import (
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

	if err := tr.ListenAndAccept() ; err != nil{
		log.Fatal(err)
	}

	select {}
}