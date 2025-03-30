package p2p





type HandshakeFunc func(Peer) error

func NOPHandShakeFunc(peer Peer) error{return nil}

