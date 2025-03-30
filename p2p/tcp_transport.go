package p2p

import(
	"fmt"
	"net"
	
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

type TCPTransportOPS struct {
    ListenerAddr  string
    HandshakeFunc HandshakeFunc
    Decoder       Decoder
    OnPeer        func(Peer) error
}

type TCPTransport struct{
	TCPTransportOPS
	listenAddress string
	listener net.Listener
	shakeHands HandshakeFunc
	rpcch chan RPC
	decoder Decoder
}

//Consume implements the Transport interface which will reaturn only read-only channel 
//For reading incoming messages coming from other peer in network
func (t *TCPTransport) Consume() <- chan RPC{
	return t.rpcch
} 

//Close implements the peer interface
func (p *TCPPeer) Close() error{
	return p.conn.Close()
}
func NewTCPTransport(ops TCPTransportOPS) *TCPTransport{
	return &TCPTransport{
		TCPTransportOPS: ops,
		rpcch: make(chan RPC),
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

func(t *TCPTransport) handleConn(conn net.Conn){
	var err error

	defer func(){
		fmt.Printf("drooping peer connection: %s",err)
		conn.Close()
	}()

	peer := NewTCPPeer(conn , true)

	if err := t.HandshakeFunc(peer) ; err != nil{
		return 

	}
	rpc := RPC{}
	for{
		err = t.Decoder.Decode(conn , &rpc)
		if err == net.ErrClosed{
			
			return
		}
		if err != nil{
			fmt.Printf("TCP read error: %s", err)
			continue
		}

		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
	
	
}
