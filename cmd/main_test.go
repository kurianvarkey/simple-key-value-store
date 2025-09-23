package main

import (
	"bytes"
	"kurianvarkey/simple-key-value-store/cmd/app"
	"kurianvarkey/simple-key-value-store/cmd/store"
	"strings"
	"testing"
)

func TestRunApp(t *testing.T) {
	mockStore := store.NewMockStore()

	// Prepare a series of commands to simulate user input.
	input := `SET key1 value1
GET key1
LIST
DELETE key1
EXIT	
`
	// Create an in-memory buffer to simulate STDIN.
	in := strings.NewReader(input)

	// Create an in-memory buffer to capture STDOUT.
	var out bytes.Buffer

	// Run the application with mock I/O and store.
	app.RunApp(mockStore, in, &out)

	// Define the expected output.
	expectedOutput := `Welcome to the Key-Value Store!
 - SET <key> <value> - Set a key-value pair
 - GET <key> - Get the value for a key
 - DELETE <key> - Delete a key-value pair
 - LIST - List all key-value pairs
 - EXIT - Terminate the program

Enter a command: Key key1 set successfully
Enter a command: Key key1: value1
Enter a command: Key: key1 Value: value1
Enter a command: Key key1 deleted successfully
Enter a command: Goodbye!
`
	// Compare the captured output with the expected output.
	if out.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, out.String())
	}
}
