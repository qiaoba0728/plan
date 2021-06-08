package test

import "encoding/json"

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
