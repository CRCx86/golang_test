package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listen, ok := net.Listen("tcp", ":8080")
	if ok != nil {
		panic(ok)
	}
	for {
		accept, ok := listen.Accept()
		if ok != nil {
			panic(ok)
		}
		go handleConnection(accept)
	}

}

func handleConnection(con net.Conn) {
	fmt.Printf("remote address %s", con.RemoteAddr().String())
	con.Write([]byte("Hello!\n"))
	defer con.Close()
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			break
		} else {
			con.Write([]byte(text))
		}
	}
}
