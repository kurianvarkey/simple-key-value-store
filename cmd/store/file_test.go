package store

import (
	"reflect"
	"testing"
)

func TestNewFileStore(t *testing.T) {
	store := NewFileStore()
	expected := &fileStore{
		store: make(map[string]string),
	}

	if !reflect.DeepEqual(store, expected) {
		t.Errorf("NewFileStore() = %v, want %v", store, expected)
	}
}

// TestFileStore_Set tests the Set method
// Set should update if the key already exists
func TestFileStore_Set(t *testing.T) {
	store := NewFileStore()
	err := store.Set("key", "value")
	if err != nil {
		t.Errorf("Set() error = %v, want nil", err)
	}

	value, _ := store.Get("key")
	if value != "value" {
		t.Errorf("Set() value = %v, want %v", value, "value")
	}

	_ = store.Set("key", "value1")
	value, _ = store.Get("key")
	if value != "value1" {
		t.Errorf("Set() value = %v, want %v", value, "value1")
	}
}

func TestFileStore_Get(t *testing.T) {
	store := NewFileStore()
	_, err := store.Get("key")
	if err == nil {
		t.Errorf("Get() error = %v, want not nil", err)
	}

	_ = store.Set("key", "value")
	value, _ := store.Get("key")
	if value != "value" {
		t.Errorf("Set() value = %v, want %v", value, "value")
	}
}

func TestFileStore_Delete(t *testing.T) {
	store := NewFileStore()
	_ = store.Set("key", "value")
	err := store.Delete("key")
	if err != nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}

	_, err = store.Get("key")
	if err == nil {
		t.Errorf("Delete() error = %v, want not nil", err)
	}

	err = store.Delete("key")
	if err == nil {
		t.Errorf("Delete() error = %v, want nil", err)
	}
}

func TestFileStore_List(t *testing.T) {
	store := NewFileStore()
	_ = store.Set("key", "value")
	values, _ := store.List()
	if values["key"] != "value" {
		t.Errorf("List() value = %v, want %v", values["key"], "value")
	}

	_ = store.Delete("key")
	_, err := store.List()
	if err == nil {
		t.Errorf("List() error = %v, want not nil", err)
	}
}
