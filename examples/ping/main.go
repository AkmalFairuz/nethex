package main

import (
	"fmt"
	"github.com/akmalfairuz/nethex/client"
	"github.com/akmalfairuz/nethex/packet"
	"github.com/akmalfairuz/nethex/server"
	"io"
	"net"
	"time"
)

func main() {
	srv, err := server.NewServer("127.0.0.1:7773")
	srv.PacketPool().Register((&Ping{}).Id(), func() packet.Packet {
		return &Ping{}
	})
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			fmt.Printf("Accepting...\n")
			c, err := srv.Accept()
			if err != nil {
				panic(err)
			}
			go handleClient(c)
		}
	}()

	c, err := client.NewClient("127.0.0.1:7773")
	c.PacketPool().Register((&Ping{}).Id(), func() packet.Packet {
		return &Ping{}
	})
	go func() {
		for {
			fmt.Println("Sending ping packet...")
			c.WritePacket(&Ping{RequestTimestamp: time.Now()})
			c.Flush()
			fmt.Println("Waiting for response...")
			pkt, err := c.ReadPacket()
			if err != nil {
				if err == io.EOF || err == net.ErrClosed {
					panic("Disconnected")
				}
				panic(err)
			}
			switch pk := pkt.(type) {
			case *Ping:
				ping := pk.ResponseTimestamp.Sub(pk.RequestTimestamp).Microseconds()
				fmt.Printf("Ping: %d microseconds\n", ping)
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		time.Sleep(time.Second * 30)
		c.Close()
	}()

	time.Sleep(time.Minute * 10)
}

func handleClient(c *server.Client) {
	for {
		pkt, err := c.ReadPacket()
		if err != nil {
			if err == io.EOF || err == net.ErrClosed {
				panic("Disconnected")
			}
			panic(err)
		}
		switch pk := pkt.(type) {
		case *Ping:
			pk.ResponseTimestamp = time.Now()
			c.WritePacket(pk)
			c.Flush()
		}
	}
}
