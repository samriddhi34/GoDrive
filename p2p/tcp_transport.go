package p2p

import(
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct{
	conn net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn , outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound, 
	}
}

type TCPTransportOPS struct{
	ListenerAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder
}
type TCPTransport struct{
	TCPTransportOPS
	listenAddress string
	listener net.Listener
	shakeHands HandshakeFunc
	decoder Decoder

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(ops TCPTransportOPS) *TCPTransport{
	return &TCPTransport{
		TCPTransportOPS: ops,
	}
}

func (t *TCPTransport) ListenAndAccept() error{
	var err error
	t.listener, err = net.Listen("tcp", t.ListenerAddr)
	if err!= nil{
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop(){
	for {
		conn , err := t.listener.Accept()
		if err != nil{
			fmt.Printf("TCP accept Error : %s\n" , err)
		}
		fmt.Printf("new incoming connection: %+v", conn)
		go t.handleConn(conn)
	}
}

func(t *TCPTransport)handleConn(conn net.Conn){
	peer := NewTCPPeer(conn , true)
	if err := t.HandshakeFunc(peer) ; err != nil{
		conn.Close()
		fmt.Printf("TCP Handshake Error: %s\n" , err)
		return 

	}
	
	for{
		msg := &Message{}
		if err := t.Decoder.Decode(conn , msg); err != nil{
			fmt.Printf("TCP error %s\n" , err)
			continue
		}
		fmt.Printf("Received from %s: %s\n", msg.From, string(msg.Payload))
	}
	
	
}
