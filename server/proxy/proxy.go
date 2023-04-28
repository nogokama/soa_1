package proxy

import (
	"fmt"
	"net"
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

	fmt.Println("Proxy listening on", conn.LocalAddr())

	for {
		listen(conn, workerPorts)
	}
}

func listen(conn net.PacketConn, workerPorts map[string]string) {
	var answer string
	buffer := make([]byte, 1024)
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
	port, ok := workerPorts[mode]
	if !ok {
		answer = fmt.Sprintf("port not found for %s, len %d ", messageParts[1], len(messageParts[1]))
		fmt.Println(answer)
		return
	}

	answer = getResult(port)
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

	buffer := make([]byte, 1024)
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to receive response:", err)
		return ""
	}

	response := buffer[:bytesRead]

	return string(response)
}
