package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func FlattenJSON(raw []byte, separator string) (map[string]string, error) {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, err
	}

	out := make(map[string]string)
	flatten("", v, separator, out)
	return out, nil
}

func flatten(prefix string, v any, separator string, out map[string]string) {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			np := k
			if prefix != "" {
				np = prefix + separator + k
			}
			flatten(np, val, separator, out)
		}
	case []any:
		for i, val := range x {
			idx := strconv.Itoa(i)
			np := idx
			if prefix != "" {
				np = prefix + separator + idx
			}
			flatten(np, val, separator, out)
		}
	default:
		out[prefix] = fmt.Sprint(x)
	}

}

func main() {
	in := []byte(`{"user":{"id":42,"name":"jim","tags":["eng","guitar"],"prefs":{"darkMode":true}}}`)
	m, err := FlattenJSON(in, ".")
	if err != nil {
		panic(err)
	}

	for k, v := range m {
		fmt.Printf("%s=%s\n", k, v)
	}
}