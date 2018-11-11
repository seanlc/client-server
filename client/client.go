package main

import (
    "net"
    "os"
    "fmt"
)

func check(e error){
    if e != nil {
        panic(e)
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("usage: %s IP port\n", os.Args[0])
	os.Exit(1)
    }

    addr := os.Args[1]
    port := os.Args[2]

    addrString := addr + ":" + port
    TCPAddr, err := net.ResolveTCPAddr("tcp", addrString)
    check(err)

    conn,err := net.DialTCP("tcp", nil, TCPAddr)
    check(err)

    _,err = conn.Write([]byte("Hello, world!"))
    check(err)

    os.Exit(0)
}
