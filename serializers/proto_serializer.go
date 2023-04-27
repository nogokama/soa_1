package serializer

import (
	"soa_hw/serializers/proto_gen"

	"github.com/golang/protobuf/proto"
)

type protoSerializer struct{}

func NewProtoSerializer() *protoSerializer {
	return &protoSerializer{}
}

func (s *protoSerializer) Serialize(a TestStruct, _ int) []byte {
	msg := &proto_gen.TestStruct{
		FloatField:  a.FloatField,
		IntField:    a.IntField,
		MapField:    a.MapField,
		StringField: a.StringField,
		ArrayField:  a.ArrayField,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return data
}

func (s *protoSerializer) Deserialize(data []byte, _ int) TestStruct {
	var msg proto_gen.TestStruct
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		panic(err)
	}
	return TestStruct{
		FloatField:  msg.FloatField,
		IntField:    msg.IntField,
		MapField:    msg.MapField,
		StringField: msg.StringField,
		ArrayField:  msg.ArrayField,
	}
}
