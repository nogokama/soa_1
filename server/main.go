package main

import (
	"flag"
	"fmt"
	"soa_hw/config"
	"soa_hw/server/proxy"
	"soa_hw/server/worker"
)

func main() {
	var programType string
	var mode string
	var port int

	var jsonPort string
	var xmlPort string
	var nativePort string
	var protoPort string
	var avroPort string
	var yamlPort string
	var msgPackPort string

	flag.StringVar(&programType, "program", config.ProgramWorker, "select mode of program")
	flag.StringVar(&mode, "mode", config.JsonMode, "select format to use in the worker program")
	flag.IntVar(&port, "port", -1, "port for program listening udp connections")

	for _, p := range []struct {
		name string
		ptr  *string
	}{
		{"json-port", &jsonPort},
		{"xml-port", &xmlPort},
		{"native-port", &nativePort},
		{"proto-port", &protoPort},
		{"avro-port", &avroPort},
		{"yaml-port", &yamlPort},
		{"msgpack-port", &msgPackPort},
	} {
		flag.StringVar(p.ptr, p.name, "", fmt.Sprintf("port for %s format", p.name))
	}

	flag.Parse()

	if port == -1 {
		panic("port is obligatory parameter")
	}

	workerPorts := map[string]string{
		config.JsonMode:    jsonPort,
		config.XmlMode:     xmlPort,
		config.NativeMode:  nativePort,
		config.ProtoMode:   protoPort,
		config.AvroMode:    avroPort,
		config.YamlMode:    yamlPort,
		config.MsgPackMode: msgPackPort,
	}

	fmt.Println(workerPorts)

	switch programType {
	case config.ProgramProxy:
		proxy.Launch(port, workerPorts)
	case config.ProgramWorker:
		worker.Launch(mode, port)
	default:
		panic(fmt.Sprintf("unkown mode %q", mode))
	}
}
