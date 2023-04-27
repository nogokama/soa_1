package serializer

import (
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

type msgpackSerializer struct{}

func NewMsgpackSerializer() *msgpackSerializer {
	return &msgpackSerializer{}
}

func (s *msgpackSerializer) Serialize(a TestStruct, _ int) []byte {
	res, err := msgpack.Marshal(a)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *msgpackSerializer) Deserialize(data []byte, _ int) TestStruct {
	var res TestStruct
	err := msgpack.Unmarshal(data, &res)
	if err != nil {
		panic(fmt.Errorf("could not deserialize MessagePack data: %s", err))
	}

	return res
}
