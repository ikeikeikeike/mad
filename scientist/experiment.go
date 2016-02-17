package scientist

type Publisher interface {
	Publish(Result)
}

type Experiment interface {
	Use(func() interface{})
	Try(func() interface{})
	Run() Result
}

type experiment struct {
	name string

	ctrl     Observation
	observes []Observation

	publishers []Publisher
}

func (e *experiment) Use(fn func() interface{}) {
	e.ctrl = &observation{name: "Control", fn: fn}
}

func (e *experiment) Try(fn func() interface{}) {
	observe := &observation{name: "Candidate", fn: fn}
	e.observes = append(e.observes, observe)
}

func (e *experiment) Run() Result {

	e.ctrl.Call()

	for _, observe := range e.observes {
		recoverable(observe)
	}

	result := &result{
		exp: e,
	}

	for _, pub := range e.publishers {
		pub.Publish(result)
	}

	return result
}
