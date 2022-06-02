package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {

			return s[i:]
		}
	}
	return ""
}

func generate_status_headers(code int) string {
	dt := time.Now()
	return "HTTP/1.1 " + fmt.Sprint(code) + " OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\nx-frame-options: SAMEORIGIN\n\n"
}

func serverFile(conn net.Conn, file string) {
	dt := time.Now()

	//defineing some differnt headers for different content types
	headers := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\nx-frame-options: SAMEORIGIN\n\n"
	headerscss := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/css,*/*;q=0.1;\nx-frame-options: SAMEORIGIN\n\n"
	headersjs := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: application/javascript,*/*;q=0.1;\nx-frame-options: SAMEORIGIN\n\n"

	println("Loading In file -> " + file)

	//try to access the file the route is pointing to
	content, err := ioutil.ReadFile(trimFirstRune(file))
	if err != nil {
		//if we cant find it error
		fmt.Println("No File Found " + file)
		responceNoIndex := generate_status_headers(404) + "<title>Witch -> Cant Find file</title><center><img style='border-top: groove;' src='https://i.imgur.com/j5GlneF.png' </center>"
		conn.Write([]byte(responceNoIndex))
		conn.Close()
		return
	}

	//genearate a responce
	responce_text := headers + string(content)

	//if it is a css file add the css content type
	if strings.Contains(file, ".css") {
		responce_texts := headerscss + string(content)

		//send it over to the client and end the connection.
		conn.Write([]byte(responce_texts))
		conn.Close()
		return
	}

	if strings.Contains(file, ".js") {
		responce_texts := headersjs + string(content)

		//send it over to the client and end the connection.
		conn.Write([]byte(responce_texts))
		conn.Close()
		return
	}

	//send it over to the client and end the connection.
	conn.Write([]byte(responce_text))
	conn.Close()
}
