package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Serialize(obj any) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func Deserialize(bytes []byte, obj any) error {
	return json.Unmarshal(bytes, obj)
}

func GenerateCacheKey(prefix string, identifier string) string {
	if identifier == "" {
		return prefix
	}
	key := fmt.Sprintf("%s:%s", prefix, identifier)
	return strings.ToLower(key)
}
