package main

import "testing"

func Test_processCmdLn(t *testing.T) {
	o := processCmdLn()
	if len(o) > 0 { // TODO dummy test, replace it.
		t.Error()
	}
}

func Test_initiateHost_noOptions(t *testing.T) {
	options := []Option{} // no options
	_, c, err := initiateHost(options)
	defer c()

	if err != nil {
		t.Error(err)
	}
}
