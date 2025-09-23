package store

import "fmt"

// mockStore is a mock implementation for testing
type mockStore struct {
	store map[string]string
}

func NewMockStore() StoreInterface {
	return &mockStore{
		store: make(map[string]string),
	}
}

// Set a key-value pair
func (s *mockStore) Set(key string, value string) error {
	s.store[key] = value
	return nil
}

// Get the value for a key
func (s *mockStore) Get(key string) (string, error) {
	if value, ok := s.store[key]; ok {
		return value, nil
	}

	return "", fmt.Errorf("key %s not found", key)
}

// Delete a key-value pair
func (s *mockStore) Delete(key string) error {
	if _, ok := s.store[key]; ok {
		delete(s.store, key)
		return nil
	}

	return fmt.Errorf("key %s not found", key)
}

// List all key-value pairs
func (s *mockStore) List() (map[string]string, error) {
	if len(s.store) == 0 {
		return nil, fmt.Errorf("store is empty")
	}

	return s.store, nil
}

func (m *mockStore) Save() error {
	return nil
}

func (m *mockStore) Load() error {
	return nil
}
