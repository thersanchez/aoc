package readchanger

import "fmt"

type ReadChanger struct {
	rw ReadWriter
}

type ReadWriter interface {
	Read(addr int) (int, error)
	Write(addr, value int) error
}

func NewReadChanger(m ReadWriter) *ReadChanger {
	return &ReadChanger{rw: m}
}

func (rc *ReadChanger) ReadAndChange(addr int) (int, error) {
	v, err := rc.rw.Read(addr)
	if err != nil {
		return 0, fmt.Errorf("reading: %v", err)
	}
	// TODO substract 3 when needed
	err = rc.rw.Write(addr, v+1)
	if err != nil {
		return 0, fmt.Errorf("writing: %v", err)
	}
	return v, nil
}
