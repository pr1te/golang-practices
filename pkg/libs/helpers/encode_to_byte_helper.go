package helpers

import (
	"bytes"
	"encoding/gob"
	"log"
)

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
