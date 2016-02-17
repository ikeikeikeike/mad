package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ikeikeikeike/mad/scientist"
)

func userUse(user map[string]interface{}) int {
	return 1
}

func userTry(user map[string]interface{}) int {
	return 1
}

func main() {
	user := map[string]interface{}{}

	experiment := scientist.Science("experiment-name", func(e scientist.Experiment) {
		e.Use(func() interface{} { return userUse(user) })
		e.Try(func() interface{} { return userTry(user) })
		e.Try(func() interface{} { return userTry(user) })
		e.Try(func() interface{} { return userTry(user) })
		e.Try(func() interface{} { return userTry(user) })
	})

	result := experiment.Run()

	spew.Dump(result)
}
