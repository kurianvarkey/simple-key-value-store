package app

import (
	"bufio"
	"errors"
	"fmt"
	"kurianvarkey/simple-key-value-store/cmd/store"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// runApp runs the application
func RunApp() {
	displayCommands()

	store, err := initStore()
	if err != nil {
		panic(err.Error())
	}

	checkForIntrupt(store)

	reader := bufio.NewReader(os.Stdin)
	for {
		displayOutput()

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
			if err := handleExit(store); err != nil {
				println(err.Error())
			}
			return
		}

		if err := handleOperations(cmd, store, inputs); err != nil {
			println(err.Error())
			continue
		}
	}
}

// Store the file if the user presses ctrl+c
func checkForIntrupt(store store.StoreInterface) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		// save the store
		_ = store.Save()
		os.Exit(0)
	}()
}

// initStore initialises the store
func initStore() (store.StoreInterface, error) {
	store, err := store.NewStore()
	if err != nil {
		return nil, err
	}

	if err := store.Load(); err != nil {
		return nil, err
	}

	return store, nil
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
func displayCommands() {
	commands := [5]string{
		"SET <key> <value> - Set a key-value pair",
		"GET <key> - Get the value for a key",
		"DELETE <key> - Delete a key-value pair",
		"LIST - List all key-value pairs",
		"EXIT - Terminate the program",
	}

	println("Welcome to the Key-Value Store!")
	for _, command := range commands {
		println(" - " + command)
	}
	println("")
}

// displayInput displays the input prompt
func displayOutput(str ...string) {
	if len(str) > 0 {
		println(str[0])
	}
	fmt.Print("Enter a command: ")
}

// handleExit handles the EXIT command
func handleExit(store store.StoreInterface) error {
	err := store.Save()
	if err != nil {
		return errors.New("please enter a valid command: SET <key> <value>")
	}

	println("Goodbye!")
	return nil
}

// handleOperations handles the operations
func handleOperations(cmd string, store store.StoreInterface, inputs []string) error {
	switch cmd {
	case "SET":
		return handleSet(store, inputs)

	case "GET":
		return handleGet(store, inputs)

	case "DELETE":
		return handleDelete(store, inputs)

	case "LIST":
		return handleList(store)

	default:
		return errors.New("please enter a valid command. Valid commands are: SET, GET, DELETE, LIST, EXIT")
	}
}

// handleSet handles the SET command
func handleSet(store store.StoreInterface, inputs []string) error {
	if len(inputs) != 3 {
		return errors.New("please enter a valid command: SET <key> <value>")
	}

	err := store.Set(inputs[1], inputs[2])
	if err != nil {
		return errors.New("Error setting value:" + err.Error())
	}

	println("Key " + inputs[1] + " set successfully")
	return nil
}

// handleGet handles the GET command
func handleGet(store store.StoreInterface, inputs []string) error {
	if len(inputs) != 2 {
		return errors.New("please enter a valid command: GET <key>")
	}

	value, err := store.Get(inputs[1])
	if err != nil {
		return errors.New("Error getting value:" + err.Error())
	}

	println("Key " + inputs[1] + ": " + value)
	return nil
}

// handleDelete handles the DELETE command
func handleDelete(store store.StoreInterface, inputs []string) error {
	if len(inputs) != 2 {
		return errors.New("please enter a valid command: DELETE <key>")
	}

	err := store.Delete(inputs[1])
	if err != nil {
		return errors.New("Error deleting value:" + err.Error())
	}

	println("Key " + inputs[1] + " deleted successfully")
	return nil
}

// handleList handles the LIST command
func handleList(store store.StoreInterface) error {
	values, err := store.List()
	if err != nil {
		return errors.New("Error listing values: " + err.Error())
	}

	for key, value := range values {
		println("Key:", key, "Value:", value)
	}

	return nil
}
