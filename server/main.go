package main

import (
	"flag"
	"fmt"
	"soa_hw/config"
	serializer "soa_hw/serializers"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", config.JsonMode, "select mode to use the programm")
	flag.Parse()

	testSerializer := config.GetSerializer(mode)

	size, ser, deser := serializer.Benchmark(testSerializer)

	fmt.Println(size, ser.Microseconds(), deser.Microseconds())

}
