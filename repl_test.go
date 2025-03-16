package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{ input: " hello world ", expected: []string{"hello", "world"} },
		{ input: "Hello World", expected: []string{"hello", "world"} },
		{ input: "Hello", expected: []string{"hello"} },
		{ input: " ", expected: []string{} },
	}

 	// iterate over the test cases
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q) == %q, expected %q", c.input, actual, c.expected)
			}
		}
	}
}

func TestCommandHelp(t *testing.T) {
	commands := map[string]cliCommand{
		"help":   {name: "help", description: "Displays help"},
		"exit":   {name: "exit", description: "Exits the program"},
	}

	// Capture stdout
	oldStdout := os.Stdout // Save old stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call function
	err := commandHelp(commands)
	if err != nil {
		t.Fatalf("commandHelp returned an error: %v", err)
	}

	// Close writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	expectedOutput := `Welcome to the Pokedex!
Usage:

help: Displays help
exit: Exits the program
`

	if output != expectedOutput {
		t.Errorf("commandHelp() output mismatch.\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
	}
}