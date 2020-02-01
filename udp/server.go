package main

import (
	"context"
	"net"
)

const bufSize = 1024

func main() {

}

func server(ctx context.Context, address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		panic(err)
	}
	defer pc.Close()
	doneChan := make(chan error, 1)
	buf := make([]byte, bufSize)

}
