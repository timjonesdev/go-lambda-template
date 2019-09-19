package main

import "testing"

func TestHandleRequest(t *testing.T) {
	returnVal,a := HandleRequest(nil, MyEvent{"Kappa"})
        testVal := "Hello Kappa!"
        if returnVal != testVal {
		t.Errorf("Return incorrect, got: %s, expected: %s",returnVal,testVal)
	}
	_ = a
}
