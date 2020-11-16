package main
import (
	"fmt"
	"tcpudp-server/client"
)
func main (){
	uc,err := client.UDPConnect("127.0.0.1:6666",1024)
	if err !=nil {
		fmt.Println(err)
		panic(err)
	}
	uc.ChWrite <- []byte("hello")


	tc, err := client.TCPConnect("127.0.0.1:8888",1024)
	if err != nil {
		panic(err)
	}
	tc.Write([]byte("world"))
	return
}