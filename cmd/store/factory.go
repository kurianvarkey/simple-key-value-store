package store

import "fmt"

// NewStore returns a new Store
func NewStore(store ...string) (StoreInterface, error) {
	storeName := "file"
	if len(store) > 0 {
		storeName = store[0]
	}

	switch storeName {
	case "file":
		return NewFileStore(), nil
	default:
		return nil, fmt.Errorf("store %s not found", storeName)
	}
}
