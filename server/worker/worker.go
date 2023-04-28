package worker

import (
	"fmt"
	"net"
	"soa_hw/config"
	serializer "soa_hw/serializers"
	"time"
)

func Launch(mode string, port int) {
	testSerializer := config.GetSerializer(mode)

	conn, err := net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Server listening on", conn.LocalAddr())

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading message: ", err)
			continue
		}

		fmt.Printf("Received message from %v: %v\n", addr, string(buffer[:n]))

		size, serTime, deserTime := serializer.Benchmark(testSerializer)

		answer := fmt.Sprintf("%s - %d - %dus - %dus\n", mode, size, serTime/time.Microsecond, deserTime/time.Microsecond)

		_, err = conn.WriteTo([]byte(answer), addr)
		if err != nil {
			fmt.Println("Error sending message: ", err)
			continue
		}

		fmt.Println("Send answer:", answer)
	}
}
