package main

import "context"
import "io"
import "net"
import "time"

const bufSize = 1024

func client(ctx context.Context, address string, reader io.Reader) error {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	doneChan := make(chan error, 1)

	go writeBytes()

	select {
	case <-ctx.Done():
		println("done")
	case err = <-doneChan:
		println(err.Error())
	}

	return nil

}

func writeBytes() {

}
