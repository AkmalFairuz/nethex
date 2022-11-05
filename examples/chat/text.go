package main

import "github.com/akmalfairuz/nethex/packet"

type Text struct {
	packet.Packet

	Text string
}

func (t *Text) Marshal(w *packet.Writer) {
	w.String(t.Text)
}

func (t *Text) Unmarshal(r *packet.Reader) {
	t.Text = r.String()
}

func (t *Text) Id() byte {
	return 2
}
