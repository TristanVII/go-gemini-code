package utils

import "fmt"

func GetArg[T any](args map[string]any, key string) (T, error) {
	val, ok := args[key]
	if !ok {
		var x T
		return x, fmt.Errorf("Args does not contain key '%s'", key)
	}
	typedVal, ok := val.(T)
	if !ok {
		var x T
		return x, fmt.Errorf("Value for '%s' is of wrong type", key)
	}
	return typedVal, nil
}
