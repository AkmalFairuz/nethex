package main

import (
	"bufio"
	"fmt"
	"github.com/akmalfairuz/nethex/packet"
	"github.com/akmalfairuz/nethex/server"
	"io"
	"net"
	"os"
)

func main() {
	srv, err := server.NewServer("127.0.0.1:7773")
	srv.PacketPool().Register((&Text{}).Id(), func() packet.Packet {
		return &Text{}
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Accepting...")
	c, err := srv.Accept()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			pkt, err := c.ReadPacket()
			if err == io.EOF || err == net.ErrClosed {
				panic("Disconnected")
			}
			if err != nil {
				panic(err)
			}
			switch pk := pkt.(type) {
			case *Text:
				fmt.Printf("Client: %s\n", pk.Text)
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Printf("\n")
		c.WritePacket(&Text{Text: text})
		c.Flush()
	}
}
