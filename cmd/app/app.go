package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"kurianvarkey/simple-key-value-store/cmd/store"
	"strings"
)

// runApp runs the application
func RunApp(store store.StoreInterface, in io.Reader, out io.Writer) {
	displayCommands(out)
	reader := bufio.NewReader(in)

	for {
		displayOutput(out)

		input, err := readInput(reader)
		if err != nil {
			println(err.Error())
			return
		}

		if input == "" {
			continue
		}

		inputs := strings.Fields(input)
		cmd := strings.ToUpper(inputs[0]) // input[0] is the command
		if cmd == "EXIT" {
			if err := handleExit(store, out); err != nil {
				println(err.Error())
			}
			return
		}

		if err := handleOperations(cmd, store, inputs, out); err != nil {
			println(err.Error())
			continue
		}
	}
}

// readInput reads input from the user
func readInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		if err.Error() == "EOF" {
			return "", errors.New("exiting due to EOF")
		}
		println(err.Error())
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// displayCommands displays the available commands
func displayCommands(out io.Writer) {
	commands := [5]string{
		"SET <key> <value> - Set a key-value pair",
		"GET <key> - Get the value for a key",
		"DELETE <key> - Delete a key-value pair",
		"LIST - List all key-value pairs",
		"EXIT - Terminate the program",
	}

	fmt.Fprintln(out, "Welcome to the Key-Value Store!")
	for _, command := range commands {
		fmt.Fprintln(out, " - "+command)
	}
	fmt.Fprintln(out, "")
}

// displayInput displays the input prompt
func displayOutput(out io.Writer, str ...string) {
	if len(str) > 0 {
		fmt.Fprint(out, str[0])
	}
	fmt.Fprint(out, "Enter a command: ")
}

// handleExit handles the EXIT command
func handleExit(store store.StoreInterface, out io.Writer) error {
	err := store.Save()
	if err != nil {
		return errors.New("please enter a valid command: SET <key> <value>")
	}

	fmt.Fprint(out, "Goodbye!\n")
	return nil
}

// handleOperations handles the operations
func handleOperations(cmd string, store store.StoreInterface, inputs []string, out io.Writer) error {
	switch cmd {
	case "SET":
		return handleSet(store, inputs, out)

	case "GET":
		return handleGet(store, inputs, out)

	case "DELETE":
		return handleDelete(store, inputs, out)

	case "LIST":
		return handleList(store, out)

	default:
		return errors.New("please enter a valid command. Valid commands are: SET, GET, DELETE, LIST, EXIT")
	}
}

// handleSet handles the SET command
func handleSet(store store.StoreInterface, inputs []string, out io.Writer) error {
	if len(inputs) != 3 {
		return errors.New("please enter a valid command: SET <key> <value>")
	}

	err := store.Set(inputs[1], inputs[2])
	if err != nil {
		return errors.New("Error setting value:" + err.Error())
	}

	fmt.Fprintln(out, "Key "+inputs[1]+" set successfully")
	return nil
}

// handleGet handles the GET command
func handleGet(store store.StoreInterface, inputs []string, out io.Writer) error {
	if len(inputs) != 2 {
		return errors.New("please enter a valid command: GET <key>")
	}

	value, err := store.Get(inputs[1])
	if err != nil {
		return errors.New("Error getting value:" + err.Error())
	}

	fmt.Fprintln(out, "Key "+inputs[1]+": "+value)
	return nil
}

// handleDelete handles the DELETE command
func handleDelete(store store.StoreInterface, inputs []string, out io.Writer) error {
	if len(inputs) != 2 {
		return errors.New("please enter a valid command: DELETE <key>")
	}

	err := store.Delete(inputs[1])
	if err != nil {
		return errors.New("Error deleting value:" + err.Error())
	}

	fmt.Fprintln(out, "Key "+inputs[1]+" deleted successfully")
	return nil
}

// handleList handles the LIST command
func handleList(store store.StoreInterface, out io.Writer) error {
	values, err := store.List()
	if err != nil {
		return errors.New("Error listing values: " + err.Error())
	}

	for key, value := range values {
		fmt.Fprintln(out, "Key:", key, "Value:", value)
	}

	return nil
}
