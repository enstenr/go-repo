package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Server is running on localhost:8080...")
	fmt.Println("Send a GET request from your browser or terminal!")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)

		}
		go handleRequestEver(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, _ := reader.ReadString('\n')
	fmt.Printf("[%s] Received: %s", conn.RemoteAddr(), line)
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"Hello Rajesh!"
	conn.Write([]byte(response))
}

func handleRequestEver(conn net.Conn) {
	// We still want to close it eventually if the browser disappears
	defer conn.Close()

	// Create the reader ONCE for this connection
	reader := bufio.NewReader(conn)

	for {
		// This will block until the browser sends another request on the SAME port
		line, err := reader.ReadString('\n')
		if err != nil {
			// If the browser eventually hangs up, we exit the loop
			return
		}

		fmt.Printf("[%s] Received: %s", conn.RemoteAddr(), line)

		// Important: We need to drain the other headers sent by the browser
		// until we hit the empty line, otherwise the next 'ReadString' will
		// just read the next header (like Host:) instead of a new GET request.
		for {
			header, _ := reader.ReadString('\n')
			if header == "\r\n" || header == "\n" {
				break
			}
		}

		response := "HTTP/1.1 200 OK\r\n" +
			"Content-Type: text/plain\r\n" +
			"Content-Length: 13\r\n" +
			"Connection: keep-alive\r\n" +
			"\r\n" +
			"Hello Rajesh!"

		conn.Write([]byte(response))

		// The loop goes back to the top and waits for more data
		// WITHOUT closing the connection.
	}
}
