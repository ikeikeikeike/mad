package scientist

func Science(name string, fn func(e Experiment)) Experiment {
	e := &experiment{name: name}
	fn(e)
	return e
}
