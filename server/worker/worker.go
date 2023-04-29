package worker

import (
	"fmt"
	"log"
	"net"
	"os"
	"soa_hw/config"
	serializer "soa_hw/serializers"
	"sync"
	"time"
)

func Launch(mode string, port int) {
	testSerializer := config.GetSerializer(mode)

	multicastAddr := os.Getenv(config.MulticastGroupAddr)
	listenAddr, err := net.ResolveUDPAddr("udp4", multicastAddr)
	if err != nil {
		panic(fmt.Errorf("failed to resolve udp addr: %w", err))
	}

	multicastConn, err := net.ListenMulticastUDP("udp4", nil, listenAddr)
	if err != nil {
		panic(err)
	}

	multicastConn.SetReadBuffer(config.MaxDGRAMSize)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		listenConnection(mode, multicastConn, testSerializer)
	}()

	ordinaryAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	ordinaryConn, err := net.ListenUDP("udp4", ordinaryAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer wg.Done()
		listenConnection(mode, ordinaryConn, testSerializer)
	}()

	wg.Wait()
}

func listenConnection(mode string, conn *net.UDPConn, testSerializer serializer.Serializer) {
	fmt.Println("Server listening on", conn.LocalAddr())

	for {
		buffer := make([]byte, config.MaxDGRAMSize)
		n, src, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		fmt.Printf("Received message from %v: %v\n", src, string(buffer[:n]))

		size, serTime, deserTime := serializer.Benchmark(testSerializer)

		answer := fmt.Sprintf("%s - %d - %dus - %dus\n", mode, size, serTime/time.Microsecond, deserTime/time.Microsecond)

		_, err = conn.WriteToUDP([]byte(answer), src)
		if err != nil {
			fmt.Println("Error sending message: ", err)
			continue
		}

		fmt.Println("Send answer:", answer)
	}
}
