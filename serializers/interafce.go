package serializer

import (
	"encoding/xml"
	"io"
)

type Serializer interface {
	Serialize(a TestStruct, attempt int) []byte
	Deserialize(data []byte, attempt int) TestStruct
}

type StringMap map[string]string

type TestStruct struct {
	FloatField  float64   `json:"float" xml:"float" yaml:"float"`
	IntField    int64     `json:"int" xml:"int" yaml:"int"`
	MapField    StringMap `json:"map" xml:"map" yaml:"map"`
	StringField string    `json:"string" xml:"string" yaml:"string"`
	ArrayField  []int32   `json:"array" xml:"array" yaml:"array"`
}

var AvroSchemaJSON = `
{
	"type": "record",
	"name": "TestStruct",
	"fields": [
	  {
		"name": "float",
		"type": "double"
	  },
	  {
		"name": "int",
		"type": "long"
	  },
	  {
		"name": "map",
		"type": {
		  "type": "map",
		  "values": "string"
		}
	  },
	  {
		"name": "string",
		"type": "string"
	  },
	  {
		"name": "array",
		"type": {
		  "type": "array",
		  "items": "int"
		}
	  }
	]
  }
`

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = StringMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

var TestStructBenchmark = TestStruct{
	FloatField: 134626.1234,
	IntField:   2134,
	MapField: StringMap(map[string]string{
		"key_1":  "value_1",
		"key_2":  "value_2",
		"key_3":  "value_3",
		"key_4":  "value_4",
		"key_5":  "value_5",
		"key_6":  "value_6",
		"key_7":  "value_7",
		"key_8":  "value_8",
		"key_9":  "value_9",
		"key_10": "value_10",
		"key_11": "value_11",
		"key_12": "value_12",
		"key_13": "value_13",
		"key_14": "value_14",
		"key_15": "value_15",
		"key_16": "value_16",
		"key_17": "value_17",
		"key_18": "value_18",
		"key_19": "value_19",
		"key_20": "value_20",
		"key_21": "value_21",
		"key_22": "value_22",
		"key_23": "value_23",
		"key_24": "value_24",
		"key_25": "value_25",
		"key_26": "value_26",
		"key_27": "value_27",
		"key_28": "value_28",
		"key_29": "value_29",
	}),
	StringField: "Some big string",
	ArrayField:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
}
