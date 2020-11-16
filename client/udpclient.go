package client

import (
	"fmt"
	"net"
	"runtime"
)

// chanLen Length of read write channel
const chanLen = 10

type (
	// Client main struct
	UDPClient struct {
		conn      net.Conn    //
		frameSize uint16      //
		ChRead    chan []byte //
		ChWrite   chan []byte //
	}
)

// New constructor of a new server
func UDPConnect(addr string, frameSize uint16) (t *UDPClient, err error) {

	if frameSize == 0 {
		frameSize = 65507
	}

	t = &UDPClient{
		ChRead:    make(chan []byte, chanLen),
		ChWrite:   make(chan []byte, chanLen),
		frameSize: frameSize,
	}

	t.conn, err = net.Dial("udp", addr)
	if err != nil {
		return
	}

	go t.reader()
	go t.writer()

	// destroy action
	stopAllGorutines := func(t *UDPClient) {
		close(t.ChWrite)
		t.conn.Close()
	}
	runtime.SetFinalizer(t, stopAllGorutines)
	return
}

func (t *UDPClient) reader() {
	var (
		e error
		n int
	)

	for {
		buf := make([]byte, t.frameSize)
		n, e = t.conn.Read(buf)
		fmt.Printf("read from : %s",string(buf[:n]))

		if e == nil {
			t.ChRead <- buf[:n]
		} else {
			return
		}
	}
}

func (t *UDPClient) writer() {
	for {
		fmt.Println("write")
		for v := range t.ChWrite{
			fmt.Printf("data:%s",string(v))
			t.conn.Write(v)
		}
	}
}