package packet

type Pool struct {
	pool map[byte]func() Packet
}

func NewPool() *Pool {
	return &Pool{
		pool: map[byte]func() Packet{},
	}
}

func (p *Pool) Register(id byte, f func() Packet) {
	p.pool[id] = f
}

func (p *Pool) Get(id byte) (Packet, bool) {
	f, ok := p.pool[id]
	if !ok {
		return nil, false
	}
	return (f)(), true
}
