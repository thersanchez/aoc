package main

type ReadChanger struct {
	rw ReadWriter
}

type ReadWriter interface {
	Read(addr int) (int, error)
	Write(addr, value int) error
}

func NewReadChanger(m ReadWriter) *ReadChanger {
	return nil
}

func (rc *ReadChanger) ReadAndChange(int) (int, error) {
	//rw.Read()
	//rw.Write()
	return 0, nil
}
