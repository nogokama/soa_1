package serializer

import (
	"reflect"
	"time"
)

const (
	benchMarkAttempts = 1000
)

func Benchmark(serializer Serializer) (serializedSize int, serializeTime time.Duration, deserializeTime time.Duration) {
	test := TestStructBenchmark
	result := serializer.Serialize(test, 0)
	serializedSize = len(result)

	startTime := time.Now()

	for i := 0; i < benchMarkAttempts; i++ {
		result = serializer.Serialize(test, i)
	}

	serializeTime = time.Since(startTime) / benchMarkAttempts

	startTime = time.Now()
	var dRes TestStruct
	for i := 0; i < benchMarkAttempts; i++ {
		dRes = serializer.Deserialize(result, i)
	}

	if !reflect.DeepEqual(dRes, test) {
		panic("serialize and deserialize values are different")
	}

	deserializeTime = time.Since(startTime) / benchMarkAttempts

	return
}
