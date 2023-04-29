package config

import serializer "soa_hw/serializers"

const (
	JsonMode      = "json"
	XmlMode       = "xml"
	NativeMode    = "native"
	ProtoMode     = "proto"
	AvroMode      = "avro"
	YamlMode      = "yaml"
	MsgPackMode   = "msg_pack"
	MulticastMode = "all"
)

const (
	ProgramProxy  = "proxy"
	ProgramWorker = "worker"
)

const (
	MulticastGroupAddr = "MULTICAST_GROUP_ADDR"
)

const (
	MaxDGRAMSize = 1024
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
