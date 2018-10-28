package main

import "testing"

func Test_processCmdLn(t *testing.T) {
	o := processCmdLn()
	if len(o) > 0 {
		t.Error()
	}
}
