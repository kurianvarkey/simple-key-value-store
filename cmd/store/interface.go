package store

type StoreInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	List() (map[string]string, error)
	Save() error
	Load() error
}
