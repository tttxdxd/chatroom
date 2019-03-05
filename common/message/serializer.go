package message

import (
	"encoding/json"
)

func Serializer(msg interface{}) (data []byte, err error) {
	data, err = json.Marshal(msg)
	return
}

func Unserializer(data []byte) (msg *interface{}, err error) {
	err = json.Unmarshal(data, msg)
	return
}
