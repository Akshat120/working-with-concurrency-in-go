package main

import "testing"

func Test_UpdateMessage(t *testing.T) {
	msg = "hello, world!"
	wg.Add(2)
	go updateMessage("x")
	go updateMessage("goodbye world!")
	wg.Wait()

	if msg != "goodbye world!" {
		t.Error("incorrect value in msg")
	}
}
