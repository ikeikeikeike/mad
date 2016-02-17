package scientist

func recoverable(observe Observation) {
	observe.Call()
}

type Observation interface {
	Call()
}

type observation struct {
	name string
	fn   func() interface{}

	value interface{}
}

func (o *observation) Call() {
	o.value = o.fn()
}
