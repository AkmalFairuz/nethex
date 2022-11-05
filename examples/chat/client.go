package main

import (
	"bufio"
	"fmt"
	"github.com/akmalfairuz/nethex/client"
	"github.com/akmalfairuz/nethex/packet"
	"io"
	"net"
	"os"
)

func main() {
	c, err := client.NewClient("127.0.0.1:7773")
	c.PacketPool().Register((&Text{}).Id(), func() packet.Packet {
		return &Text{}
	})
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
				fmt.Printf("Server: %s\n", pk.Text)
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
