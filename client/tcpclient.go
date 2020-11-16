package client

import (
	"net"
	"runtime"
)

type TCPClient struct {
	conn net.Conn
	Address string
	frameSize int

}

func TCPConnect(addr string,frameSize int ) (t *TCPClient,err error) {
 t.conn, err = net.Dial("tcp",addr)
 if err != nil {
 	return nil,err
 }
 t.frameSize = frameSize

	// destroy action
	stopAllGorutines := func(t *TCPClient) {
		t.conn.Close()
	}
	runtime.SetFinalizer(t, stopAllGorutines)
	return t,nil
}

func (t *TCPClient) Write(data []byte)  (err error){
	_,err = t.conn.Write(data)
	if err != nil {
		return err
	}
return nil
}
