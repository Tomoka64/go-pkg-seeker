package main

import (
	"fmt"
	"os"
)

//Helper consists of a message to return when the request from user is not what is expected.
type Helper struct {
	message string
}

func newHelper(items ...string) (Driver, error) {
	return &Helper{
		fmt.Sprintf("usage1: %s <directory> <keyword> (e.g. fmt TODO)\nusage2: localhost (to connect with localhost)\nusage3: history (to see your search history)", items[0]),
	}, nil
}

//Run will write the Helper's message to terminal
func (r *Helper) Run() error {
	fmt.Fprintln(os.Stdout, r.message)
	os.Exit(1)
	return nil
}
