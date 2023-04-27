package serializer

import (
	"github.com/linkedin/goavro/v2"
)

type avroSerializer struct{}

func NewAvroSerializer() *avroSerializer {
	return &avroSerializer{}
}

func (s *avroSerializer) Serialize(a TestStruct, _ int) []byte {
	codec, err := goavro.NewCodec(AvroSchemaJSON)
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	m["float"] = a.FloatField
	m["int"] = a.IntField
	m["map"] = a.MapField
	m["string"] = a.StringField
	m["array"] = a.ArrayField

	binaryData, err := codec.BinaryFromNative(nil, m)
	if err != nil {
		panic(err)
	}

	return binaryData
}

func (s *avroSerializer) Deserialize(data []byte, _ int) TestStruct {
	codec, err := goavro.NewCodec(AvroSchemaJSON)
	if err != nil {
		panic(err)
	}
	nativeData, _, err := codec.NativeFromBinary(data)
	if err != nil {
		panic(err)
	}
	record, ok := nativeData.(map[string]interface{})
	if !ok {
		panic("failed to convert")
	}

	floatField := record["float"].(float64)
	intField := int64(record["int"].(int64))
	stringField := record["string"].(string)
	arrayField := record["array"].([]interface{})
	var intArray []int32
	for _, v := range arrayField {
		intArray = append(intArray, v.(int32))
	}
	mapField := make(map[string]string)
	mapData := record["map"].(map[string]interface{})
	for k, v := range mapData {
		mapField[k] = v.(string)
	}
	return TestStruct{
		FloatField:  floatField,
		IntField:    intField,
		MapField:    mapField,
		StringField: stringField,
		ArrayField:  intArray,
	}
}
