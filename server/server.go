package main

import (
    "fmt"
    "net"
    "os"
)

func check(e error){
    if e != nil {
        panic(e)
    }
}

func main(){
    if len(os.Args) != 2 {
	fmt.Printf("usage: %s port\n", os.Args[0])
        os.Exit(1)
    }

    msg := make([]byte, 256)

    addr := ":" + os.Args[1]

    TCPAddr,err := net.ResolveTCPAddr("tcp", addr)
    check(err)

    lsocket,err := net.ListenTCP("tcp", TCPAddr)
    check(err)

    for {
        conn, err := lsocket.Accept()
	check(err)
	_,err = conn.Read(msg)
	check(err)
	fmt.Println("Received: " + string(msg))
	conn.Close()
    }
}
