package main

import (
	"fmt"
	"log"
	"syscall"
)

const PORTNUM = 8000
const MAX_CONN = 3

type server struct {
	listen_fd   int
	addr_struct syscall.SockaddrInet4
}

func main() {
	server := open_socket()
	defer syscall.Close(server.listen_fd)

	err := syscall.Listen(server.listen_fd, MAX_CONN)
	if err != nil {
		log.Fatal(err)
	}

	conn_fd, _, err := syscall.Accept(server.listen_fd)
	defer syscall.Close(conn_fd)

	if err != nil {
		log.Fatal(err)
	}
	// Read from socket forever
	read_from_socket(conn_fd)

	fmt.Print(server.listen_fd)

}

func read_from_socket(conn_fd int) {
	buff := make([]byte, 4096)
	for {
		n, err := syscall.Read(conn_fd, buff)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(buff[:n]))
	}
}

func open_socket() (s server) {
	// opens and binds socket. fd is a flie descriptor
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal("socket", err)
	}

	s = server{listen_fd: fd}
	addr := syscall.SockaddrInet4{Port: PORTNUM}
	s.addr_struct = addr

	err = syscall.Bind(s.listen_fd, &addr)
	if err != nil {
		log.Fatal("Bind", err)
	}

	return s
}
