package main

import (
	"fmt"
	"strings"
)

type MapValue struct {
	parent map[string]any
	key    string
}

func (mv *MapValue) Get() any {
	return mv.parent[mv.key]
}

func (mv *MapValue) GetHcl() ([]byte, error) {
	switch value := mv.Get().(type) {
	case map[string]any:
		hclBytes, err := ConvertMapToHcl(value)
		if err != nil {
			return nil, err
		}
		return hclBytes, nil
	case string:
		return fmt.Appendf(nil, "%s = %q", mv.key, value), nil
	default:
		return fmt.Appendf(nil, "%s = %v", mv.key, value), nil
	}
}

func (mv *MapValue) Set(value any) {
	mv.parent[mv.key] = value
}

func FindByPath(data map[string]any, path string) (*MapValue, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path")
	}

	splitPath := strings.Split(path, ".")
	if len(splitPath) == 0 {
		return nil, fmt.Errorf("empty path")
	}

	var current any
	current = data

	for _, key := range splitPath[:len(splitPath)-1] {
		switch v := current.(type) {
		case map[string]any:
			var exists bool
			current, exists = v[key]
			if !exists {
				return nil, fmt.Errorf("key %s not found", key)
			}
		default:
			return nil, fmt.Errorf("cannot access key %s on non-object", key)
		}
	}

	parent := current.(map[string]any)

	finalKey := splitPath[len(splitPath)-1]
	if _, exists := parent[finalKey]; !exists {
		return nil, fmt.Errorf("key %s not found", finalKey)
	}

	return &MapValue{parent: parent, key: finalKey}, nil
}
