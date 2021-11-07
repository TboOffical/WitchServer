package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	dt := time.Now()

	//create a new scanner to scan the connections
	scanner := bufio.NewScanner(conn)
	scanner.Scan()

	//get the first line of the http request
	//This will tell us the type of request and the route
	request_first_line := scanner.Text()

	//defineing some differnt headers for different content types
	headers := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\n\n"
	headerscss := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/css,*/*;q=0.1;\n\n"

	//split up the first line on spaces so we can get the differnt values
	ereq := strings.Split(request_first_line, " ")

	//the request type is the first string
	request_type := ereq[0]

	//log it becaquse it is usefull for debug purpuses
	println("Request Type is " + request_type)

	//if we have a get request or a post request
	if request_type == "GET" && ereq[1] != "/" || request_type == "POST" && ereq[1] != "/" {
		println("Loading In file -> " + ereq[1])

		//try to access the file the route is pointing to
		content, err := ioutil.ReadFile(trimFirstRune(ereq[1]))
		if err != nil {
			//if we cant find it error
			fmt.Println("No File Found " + ereq[1])
			responceNoIndex := headers + "<title>Witch -> Cant Find file</title><center><img style='border-top: groove;' src='https://i.imgur.com/j5GlneF.png' </center>"
			conn.Write([]byte(responceNoIndex))
			conn.Close()
			return
		}

		//genearate a responce
		responce_text := headers + string(content)

		//if it is a css file add the css content type
		if strings.Contains(ereq[1], ".css") {
			responce_texts := headerscss + string(content)

			//send it over to the client and end the connection.
			conn.Write([]byte(responce_texts))
			conn.Close()
			return
		}

		//send it over to the client and end the connection.
		conn.Write([]byte(responce_text))
		conn.Close()
		return
	}

	//check if and load index
	contenti, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("No Index.html File Found")
		responceNoIndex := headers + "<title>Witch -> Cant Find file</title><center><img style='border-top: groove;' src='https://i.imgur.com/GJIUMVC.png' </center>"
		conn.Write([]byte(responceNoIndex))
		conn.Close()
		return
	}
	//Found Index.html Sending

	TextRES := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\n\n" + string(contenti)
	conn.Write([]byte(TextRES))

	conn.Close()
}
