package main

import (
	"github.com/akmalfairuz/nethex/packet"
	"time"
)

type Ping struct {
	packet.Packet

	RequestTimestamp  time.Time
	ResponseTimestamp time.Time
}

func (p *Ping) Marshal(w *packet.Writer) {
	w.Int64(p.RequestTimestamp.UnixMicro())
	w.Int64(p.ResponseTimestamp.UnixMicro())
}

func (p *Ping) Unmarshal(r *packet.Reader) {
	p.RequestTimestamp = time.UnixMicro(r.Int64())
	p.ResponseTimestamp = time.UnixMicro(r.Int64())
}

func (p *Ping) Id() byte {
	return 1
}
