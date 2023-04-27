package serializer

import (
	"bytes"
	"encoding/gob"
)

type gobSerializer struct{}

func NewGobSerializer() *gobSerializer {
	return &gobSerializer{}
}

func (s *gobSerializer) Serialize(a TestStruct, _ int) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(a)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func (s *gobSerializer) Deserialize(data []byte, _ int) TestStruct {
	var res TestStruct
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&res)
	if err != nil {
		panic(err)
	}
	return res
}
