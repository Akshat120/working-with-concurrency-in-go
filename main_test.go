package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

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

func Test_PrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "testing"
	printMessage()

	_ = w.Close()

	bytes, _ := io.ReadAll(r)
	result := string(bytes)
	os.Stdout = stdOut

	if !strings.Contains(result, "testing") {
		t.Errorf("Expected to match test-message, but failed")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	bytes, _ := io.ReadAll(r)
	result := string(bytes)
	os.Stdout = stdOut

	if !strings.Contains(result, "Hello, universe!") {
		t.Errorf("Expected to contain \"Hello, universe!\", but failed")
	}

	if !strings.Contains(result, "Hello, cosmos!") {
		t.Errorf("Expected to contain \"Hello, cosmos!\", but failed")
	}

	if !strings.Contains(result, "Hello, world!") {
		t.Errorf("Expected to contain \"Hello, world!\", but failed")
	}
}
