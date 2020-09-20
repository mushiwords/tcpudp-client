package main
import (
	"fmt"
	"tcpudp-server/client"
)
func main (){
	fmt.Printf("ddd")
	uc,err := client.UDPConnect("127.0.0.1:6666",1024)
	if err !=nil {
		fmt.Println(err)
		panic(err)
	}
	uc.ChWrite <- []byte("hello")
	return
}