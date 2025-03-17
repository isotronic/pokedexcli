package main

import (
	"bytes"
	"os"
	"strings"
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
		"exit":   {name: "exit", description: "Exits the program"},
		"help":   {name: "help", description: "Displays help"},
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

exit: Exits the program
help: Displays help
`

	if output != expectedOutput {
		t.Errorf("commandHelp() output mismatch.\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
	}
}
func TestCommandMap(t *testing.T) {
	config := &configType{}

	// Capture stdout
	oldStdout := os.Stdout // Save old stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call function
	err := commandMap(config)
	if err != nil {
		t.Fatalf("commandMap returned an error: %v", err)
	}

	// Close writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "eterna-city-area") {
		t.Errorf("commandMap() output does not contain expected area name.\nGot:\n%s", output)
	}

	if config.nextEndpoint == "" {
		t.Errorf("commandMap() did not update nextEndpoint")
	}
}
func TestCommandMapb(t *testing.T) {
	// Test case 1: config.previousEndpoint = valid url
	t.Run("Valid previousEndpoint", func(t *testing.T) {
		config := &configType{
			previousEndpoint: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
		}

		// Capture stdout
		oldStdout := os.Stdout // Save old stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Call function
		err := commandMapb(config)
		if err != nil {
			t.Fatalf("commandMapb returned an error: %v", err)
		}

		// Close writer and restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read output
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		output := buf.String()

		if !strings.Contains(output, "great-marsh-area-2") {
			t.Errorf("commandMapb() output does not contain expected area name.\nGot:\n%s", output)
		}

		if config.nextEndpoint == "" {
			t.Errorf("commandMapb() did not update nextEndpoint")
		}
	})

	// Test case 2: config.previousEndpoint = empty string
	t.Run("Empty previousEndpoint", func(t *testing.T) {
		config := &configType{
			previousEndpoint: "",
		}

		// Capture stdout
		oldStdout := os.Stdout // Save old stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Call function
		err := commandMapb(config)
		if err != nil {
			t.Fatalf("commandMapb returned an error: %v", err)
		}

		// Close writer and restore stdout
		w.Close()
		os.Stdout = oldStdout

		// Read output
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)
		output := buf.String()

		expectedOutput := "You're on the first page\n"
		if output != expectedOutput {
			t.Errorf("commandMapb() output mismatch.\nGot:\n%s\nExpected:\n%s", output, expectedOutput)
		}
	})
}
