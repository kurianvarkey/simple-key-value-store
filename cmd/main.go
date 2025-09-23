package main

import (
	"kurianvarkey/simple-key-value-store/cmd/app"
	"kurianvarkey/simple-key-value-store/cmd/store"
	"os"
	"os/signal"
	"syscall"
)

// main function
func main() {
	store, err := initStore()
	if err != nil {
		panic(err.Error())
	}

	checkForIntrupt(store)
	app.RunApp(store, os.Stdin, os.Stdout)
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
