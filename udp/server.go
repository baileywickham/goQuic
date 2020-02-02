package main

import (
	"context"
	"net"
)

const bufSize = 1024

type pktHeader struct {
	version  int8
	length   int16
	checksum int16
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server(ctx, "localhost:2000")
	defer cancel()

}

func server(ctx context.Context, address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	doneChan := make(chan error, 1)
	go read_from_udp(doneChan, pc)

	select {
	case <-ctx.Done():
		println("done")
	case err = <-doneChan:
	}
	return nil
}

func read_from_udp(doneChan chan error, pc net.PacketConn) {
	buf := make([]byte, bufSize)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			doneChan <- err
		}
		println("Read: ", n, addr.String())
	}

}

func checksum(buf []byte) int16 {
	chk := int16(0)
	for _, b := range buf {
		chk += int16(b)
	}
	return chk
}
