package test

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

//TODO this is a mock and needs no tests ;)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// Send initial SMTP greeting
	writer.WriteString("220 Simple SMTP Server\r\n")
	writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
			}
			return
		}
		line = strings.TrimSpace(line)
		log.Printf("Received: %s", line)

		if strings.HasPrefix(line, "HELO") || strings.HasPrefix(line, "EHLO") {
			writer.WriteString("250 Hello\r\n")
		} else if strings.HasPrefix(line, "MAIL FROM:") {
			writer.WriteString("250 OK\r\n")
		} else if strings.HasPrefix(line, "RCPT TO:") {
			writer.WriteString("250 OK\r\n")
		} else if strings.HasPrefix(line, "DATA") {
			writer.WriteString("354 End data with <CR><LF>.<CR><LF>\r\n")
			writer.Flush()

			var dataLines []string
			for {
				dataLine, err := reader.ReadString('\n')
				if err != nil {
					log.Printf("Error reading data: %v", err)
					return
				}
				if dataLine == ".\r\n" {
					break
				}
				dataLines = append(dataLines, dataLine)
			}
			log.Printf("Received email data: %s", strings.Join(dataLines, ""))

			writer.WriteString("250 OK\r\n")
		} else if strings.HasPrefix(line, "QUIT") {
			writer.WriteString("221 Bye\r\n")
			writer.Flush()
			return
		} else {
			writer.WriteString("500 Unrecognized command\r\n")
		}
		writer.Flush()
	}
}

// StartMockSmtpServer starts the mock SMTP server and listens for a stop signal.
func StartMockSmtpServer(stop chan struct{}) {
	listenAddr := ":1025"
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Printf("SMTP server listening on %s", listenAddr)

	for {
		select {
		case <-stop:
			log.Println("Stopping SMTP server...")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
				continue
			}
			go handleConnection(conn)
		}
	}
}
