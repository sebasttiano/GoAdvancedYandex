package persistent

import (
	"persistent/store"
)

func Lookup(s store.Store, key string) ([]byte, error) {
	// ...
	return s.Get(key)
}
