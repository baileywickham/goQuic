package main

import (
	"bufio"
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

const bufSize = 1024
const headerSize = 9 // odd size?

type pktHeader struct {
	version  uint8
	length   uint16
	cmd      uint8
	checksum uint32
}

func serve(ctx context.Context, address string) error {

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	doneChan := make(chan error, 1)
	go read_from_udp(doneChan, conn)

	select {
	case <-ctx.Done():
		println("done")
	case err = <-doneChan:
	}
	return nil
}

func read_from_udp(doneChan chan error, conn *net.UDPConn) {
	dec := gob.NewDecoder(conn)

	for {
		var header pktHeader
		err := dec.Decode(header)
		if err != nil {
			println("WARN: packet dropped")
		}

	}

}

func read_into_header(buf [headerSize]byte) {

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serve(ctx, "localhost:8000")
	defer cancel()

}
