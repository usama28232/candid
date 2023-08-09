package shared

import (
	"encoding/json"
	"io"
	"strings"
)

func StringSplit(input, delim string) []string {
	return strings.Split(input, delim)
}

func EncodeResponse(data interface{}, w io.Writer) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err == nil {
		return nil
	} else {
		return err
	}
}

func CollectionContainsOrStartsWith(data []string, target string) bool {
	for _, value := range data {
		if value == target || strings.Contains(target, value) {
			return true
		}
	}
	return false
}
