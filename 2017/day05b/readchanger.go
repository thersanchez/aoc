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

/*
}

// NewAutoInc returns a new AutoInc using the given Mem as its
// internal memory.
func NewAutoInc(rw ReadWriter) AutoInc {
	return AutoInc{
		rw: rw,
	}
}

// ReadAndInc returns the memory value at addr and increments
// that value in 1 in the memory.
// Returns an error if the address is <0 or >= than the memory size.
func (ai AutoInc) ReadAndInc(addr int) (int, error) {
	v, err := ai.rw.Read(addr)
	if err != nil {
		return 0, fmt.Errorf("reading: %v", err)
	}

	err = ai.rw.Write(addr, v+1)
	if err != nil {
		return 0, fmt.Errorf("writing: %v", err)
	}

	return v, nil
}
*/
