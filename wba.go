package main

import (
	"bytes"
	"net"
	"os/exec"
	"time"
)

func handleWBA(conn net.Conn, file string, requestType string) {

	dt := time.Now()
	headers := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\nx-frame-options: SAMEORIGIN\n\n"

	println("WBA -> Handleing file " + file)
	if requestType == "POST" {

		cmd := exec.Command("./script-bin/"+file+".exe", "post")

		pipe, err := cmd.StdoutPipe()
		if err != nil {
			println("Failed to run script " + file)
			println(err.Error())
			conn.Close()
			return
		}
		err = cmd.Start()
		if err != nil {
			print("Failed to run script " + file)
			conn.Close()
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(pipe)

		conn.Write([]byte(headers + string(buf.String())))
		conn.Close()
		return
	} else {

		cmd := exec.Command("./script-bin/"+file+".exe", "get")

		pipe, err := cmd.StdoutPipe()
		if err != nil {
			println("Failed to run script " + file)
			println(err.Error())
			conn.Close()
			return
		}
		err = cmd.Start()
		if err != nil {
			print("Failed to run script " + file)
			conn.Close()
			return
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(pipe)

		conn.Write([]byte(headers + string(buf.String())))
		conn.Close()
		return
	}
}
