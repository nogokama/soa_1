package serializer

import "encoding/json"

type jsonSerializer struct {
}

func NewJsonSerializer() *jsonSerializer {
	return &jsonSerializer{}
}

func (s *jsonSerializer) Serialize(a TestStruct, _ int) []byte {
	res, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *jsonSerializer) Deserialize(data []byte, _ int) TestStruct {
	var res TestStruct
	err := json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}

	return res
}
