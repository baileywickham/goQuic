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

type Client struct {
	addr string
}
type server struct {
	numConnectedClients int
	Clients             []Client
}

func (s *server) addClient(c Client) {
	s.Clients = append(s.Clients, c)
	s.numConnectedClients++
}

func serve(ctx context.Context, address string) error {
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
	for {
		buf := make([]byte, bufSize)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			doneChan <- err
		}
		println("Read: ", n, addr.String())
		println("PACKET: ", string(buf))
	}

}

func checksum(buf []byte) int16 {
	chk := int16(0)
	for _, b := range buf {
		chk += int16(b)
	}
	return chk
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serve(ctx, "localhost:8000")
	defer cancel()

}
