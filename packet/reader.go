package packet

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	r interface {
		io.Reader
		io.ByteReader
	}
}

func NewReader(r interface {
	io.Reader
	io.ByteReader
}) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Int64() int64 {
	var ret int64
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Uint64() uint64 {
	var ret uint64
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Int32() int32 {
	var ret int32
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Uint32() uint32 {
	var ret uint32
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Int16() int16 {
	var ret int16
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Uint16() uint16 {
	var ret uint16
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Int8() int8 {
	var ret int8
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Uint8() uint8 {
	var ret uint8
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Float32() float32 {
	var ret float32
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) Float64() float64 {
	var ret float64
	err := binary.Read(r.r, binary.BigEndian, &ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *Reader) String() string {
	byteStr := make([]byte, r.Uint32())
	_, err := r.r.Read(byteStr)
	if err != nil {
		panic(err)
	}
	return string(byteStr)
}
