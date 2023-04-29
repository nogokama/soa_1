package proxy

import (
	"fmt"
	"net"
	"os"
	"soa_hw/config"
	"strings"
)

const (
	getResultPart = "get_result"
)

func Launch(port int, workerPorts map[string]string) {
	conn, err := net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	broadcastConn, err := makeBroadcastConn(os.Getenv(config.MulticastGroupAddr))
	if err != nil {
		panic(fmt.Errorf("failed to make broadcast conn: %w", err))
	}
	fmt.Println("Proxy listening on", conn.LocalAddr())

	answersChan := make(chan string, 100)

	go listenAnswers(answersChan)

	for {
		listen(conn, broadcastConn, workerPorts, answersChan)
	}
}

func listenAnswers(answersChan chan string) {
	proxyAddress := os.Getenv(config.ProxyAnswersAddr)
	addr, err := net.ResolveUDPAddr("udp", proxyAddress)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, config.MaxDGRAMSize)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		answersChan <- string(buffer[:n])
	}
}

func makeBroadcastConn(address string) (net.Conn, error) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

func listen(conn net.PacketConn, broadcastConn net.Conn, workerPorts map[string]string, answersChan chan string) {
	var answer string
	buffer := make([]byte, config.MaxDGRAMSize)
	n, addr, err := conn.ReadFrom(buffer)
	if err != nil {
		answer = fmt.Sprintf("Error reading message: %v", err)
		fmt.Println(answer)
		return
	}

	defer func() {
		_, err := conn.WriteTo([]byte(answer), addr)
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}
	}()

	message := string(buffer[:n])
	fmt.Printf("Received message from %v: %v\n", addr, message)

	messageParts := strings.Split(message, " ")
	if len(messageParts) != 2 || messageParts[0] != getResultPart {
		answer = "Unknown format of request"
		fmt.Println(answer)
		return
	}
	mode := strings.TrimSpace(messageParts[1])

	if mode == config.MulticastMode {
		answer = getMulticastResult(broadcastConn, answersChan)
		return
	}

	port, ok := workerPorts[mode]
	if !ok {
		answer = fmt.Sprintf("port not found for %q\n", messageParts[1])
		fmt.Println(answer)
		return
	}

	answer = getResult(port)
}

func getMulticastResult(conn net.Conn, answersChan chan string) string {
	fmt.Println("getting multicast result")

	answer := strings.Builder{}

	_, err := conn.Write([]byte("multicast hello"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < config.FormatsCount; i++ {
		answer.WriteString(<-answersChan)
	}

	return answer.String()
}

func getResult(port string) string {
	fmt.Println("getting result from: ", port)
	conn, err := net.Dial("udp", port)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return ""
	}
	defer conn.Close()

	message := []byte("Hello, server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return ""
	}

	buffer := make([]byte, config.MaxDGRAMSize)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to receive response:", err)
		return ""
	}

	response := buffer[:bytesRead]

	return string(response)
}
