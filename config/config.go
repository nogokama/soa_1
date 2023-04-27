package config

import serializer "soa_hw/serializers"

const (
	JsonMode    = "json"
	XmlMode     = "xml"
	NativeMode  = "native"
	ProtoMode   = "proto"
	AvroMode    = "avro"
	YamlMode    = "yaml"
	MsgPackMode = "msg_pack"
)

func GetSerializer(mode string) serializer.Serializer {
	var testSerializer serializer.Serializer
	switch mode {
	case JsonMode:
		testSerializer = serializer.NewJsonSerializer()
	case XmlMode:
		testSerializer = serializer.NewXmlSerializer()
	case NativeMode:
		testSerializer = serializer.NewGobSerializer()
	case ProtoMode:
		testSerializer = serializer.NewProtoSerializer()
	case AvroMode:
		testSerializer = serializer.NewAvroSerializer()
	case YamlMode:
		testSerializer = serializer.NewYamlSerializer()
	case MsgPackMode:
		testSerializer = serializer.NewMsgpackSerializer()
	default:
		panic("unknown mode")
	}

	return testSerializer
}

func GetPort(mode string) int {
	switch mode {
	case JsonMode:
		return 11001
	case XmlMode:
		return 11002
	case NativeMode:
		return 11003
	case ProtoMode:
		return 11004
	case AvroMode:
		return 11005
	case YamlMode:
		return 11006
	case MsgPackMode:
		return 11007
	default:
		panic("unknown mode")
	}
}
