package packet

type Packet interface {
	Marshal(w *Writer)
	Unmarshal(r *Reader)
	Id() byte
}
