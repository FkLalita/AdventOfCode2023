package main

import (
	"os"
	"testing"
)

func TestCubeCount(t *testing.T) {
	// Create a temporary file with test input
	file, err := os.CreateTemp("", "input*.txt")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(file.Name())
	file.WriteString("game: 1, red 5; blue 10; green 8\n")
	file.WriteString("game: 2, red 12; blue 15; green 13\n")
	file.Close()

	// Redirect standard input to the temporary file
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin, err = os.Open(file.Name())
	if err != nil {
		t.Fatal("Error opening temporary file for stdin:", err)
	}
	defer os.Stdin.Close()

	// Redirect standard output to capture printed output
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("Error creating pipe for stdout:", err)
	}
	os.Stdout = w

	// Run the CubeCount function
	CubeCount()

	// Close the write end of the pipe, and read the captured output
	w.Close()
	capturedOutput := make([]byte, 1024)
	n, err := r.Read(capturedOutput)
	if err != nil {
		t.Fatal("Error reading captured output:", err)
	}

	// Check if the output contains the expected values
	expectedOutput := "Sum of Possible IDs: 1\nSum of Powers: 80\n"
	if string(capturedOutput[:n]) != expectedOutput {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", string(capturedOutput[:n]), expectedOutput)
	}
}
