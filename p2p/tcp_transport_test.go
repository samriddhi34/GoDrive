package p2p

import(
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestTCPTransport(t *testing.T){
	opc := TCPTransportOPS{
		ListenerAddr: ":3000",
		HandshakeFunc: NOPHandShakeFunc,
		Decoder: DefaultDecoder{},
	}
	tr := NewTCPTransport(opc)
	assert.Equal(t , tr.ListenerAddr , ":3000")

	assert.Nil(t,tr.ListenAndAccept())


}