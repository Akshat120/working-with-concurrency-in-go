package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_PrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	msg = "Hello, world!"
	testMessage := "test-message"

	wg.Add(1)
	go updateMessage(testMessage, &wg)
	wg.Wait()
	printMessage()

	_ = w.Close()

	bytes, _ := io.ReadAll(r)
	result := string(bytes)
	os.Stdout = stdOut

	if !strings.Contains(result, testMessage) {
		t.Errorf("Expected to match test-message, but failed")
	}
}

func Test_UpdateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	updated := "updated"
	go updateMessage("updated", &wg)

	wg.Wait()

	if msg != updated {
		t.Errorf("Expected to match updated, but failed")
	}
}
