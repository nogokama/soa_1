package serializer

import "encoding/xml"

type xmlSerializer struct{}

func NewXmlSerializer() *xmlSerializer {
	return &xmlSerializer{}
}

func (s *xmlSerializer) Serialize(a TestStruct, _ int) []byte {
	res, err := xml.Marshal(a)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *xmlSerializer) Deserialize(data []byte, attempt int) TestStruct {
	var res TestStruct
	err := xml.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}

	return res
}
