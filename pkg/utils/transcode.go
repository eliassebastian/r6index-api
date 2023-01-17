package utils

import (
	"bytes"
	"encoding/json"
)

func Transcode(in, out interface{}) (error, error) {
	buf := new(bytes.Buffer)
	a := json.NewEncoder(buf).Encode(in)
	b := json.NewDecoder(buf).Decode(out)

	return a, b
}
