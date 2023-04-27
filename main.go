package main

import (
	"flag"
	"fmt"
	serializer "soa_hw/serializers"
)

const (
	jsonMode    = "json"
	xmlMode     = "xml"
	nativeMode  = "native"
	protoMode   = "proto"
	avroMode    = "avro"
	yamlMode    = "yaml"
	msgPackMode = "msg_pack"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", jsonMode, "select mode to use the programm")
	flag.Parse()

	var testSerializer serializer.Serializer

	switch mode {
	case jsonMode:
		testSerializer = serializer.NewJsonSerializer()
	case xmlMode:
		testSerializer = serializer.NewXmlSerializer()
	case nativeMode:
		testSerializer = serializer.NewGobSerializer()
	case protoMode:
		testSerializer = serializer.NewProtoSerializer()
	case avroMode:
		testSerializer = serializer.NewAvroSerializer()
	case yamlMode:
		testSerializer = serializer.NewYamlSerializer()
	case msgPackMode:
		testSerializer = serializer.NewMsgpackSerializer()
	default:
		panic("unknown mode")
	}

	size, ser, deser := serializer.Benchmark(testSerializer)

	fmt.Println(size, ser.Microseconds(), deser.Microseconds())

}
