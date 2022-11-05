package packet

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	w interface {
		io.Writer
		io.ByteWriter
	}
}

func NewWriter(w interface {
	io.Writer
	io.ByteWriter
}) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Int64(x int64) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Uint64(x uint64) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Int32(x int32) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Uint32(x uint32) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Int16(x int16) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Uint16(x uint16) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Int8(x int8) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Uint8(x uint8) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Float32(x float32) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) Float64(x float64) {
	err := binary.Write(w.w, binary.BigEndian, x)
	if err != nil {
		panic(err)
	}
}

func (w *Writer) String(x string) {
	w.Uint32(uint32(len(x)))
	_, err := w.w.Write([]byte(x))
	if err != nil {
		panic(err)
	}
}
