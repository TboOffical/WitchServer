package main

import (
	"net"
	"time"
)

func handleWBA(conn net.Conn, file string, requestType string) {

	dt := time.Now()
	headers := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\nx-frame-options: SAMEORIGIN\n\n"

	println("WBA -> Handleing file " + file)
	if requestType == "POST" {

		conn.Write([]byte(generate_status_headers(405) + string("WBA is not yet finished")))
		conn.Close()
		return
	} else {

		conn.Write([]byte(headers))
		conn.Close()
		return
	}
}
