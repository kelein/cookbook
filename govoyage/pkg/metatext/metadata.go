package metatext

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

// MetadataTextMap is a map of metadata key-value pairs that can
// be used to transport structured metadata between services.
type MetadataTextMap struct {
	metadata.MD
}

// FroeachKey iterates over the key-value pairs in the map,
// calling the provided function for each pair.
func (m MetadataTextMap) FroeachKey(fn func(key string, value string) error) error {
	for k, values := range m.MD {
		for _, v := range values {
			if err := fn(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Set sets the value for the given key
func (m MetadataTextMap) Set(key, value string) {
	key = strings.ToLower(key)
	m.MD[key] = append(m.MD[key], value)
}
