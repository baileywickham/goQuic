package main

import "bufio"
import "context"
import "net"
import "os"
import "strings"
import r "github.com/baileywickham/runner"

const bufSize = 1024

func client(ctx context.Context, address string) error {
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

	go Runner(conn, doneChan)

	select {
	case <-ctx.Done():
		println("done")
	case err = <-doneChan:
		println(err.Error())
	}

	return nil

}

func writeText(text string) {
	conn.Write(text)

}

func Runner(conn *net.UDPConn, doneChan chan error) {
	reader := bufio.NewReader(os.Stdin)
	println("PACKET SHELL: ")
	for {
		print(":|: ")
		text, _ := reader.ReadString('\n')
		tokens := strings.Fields(text)
		if len(tokens) == 0 {
			printHelp()
			continue
		}
		switch tokens[0] {
		case "w":
			// defualt size of 64
			var buf []byte
			for _, str := range tokens[1:] {
				// no idea how this works
				buf = append(buf, []byte(str)...)
			}
			_, err := conn.Write(buf)
			if err != nil {
				doneChan <- err
			}

		default:
			printHelp()
		}

	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	client(ctx, "localhost:8000")
	defer cancel()
}
func printHelp() {
	println("w: write following string to server")
}
