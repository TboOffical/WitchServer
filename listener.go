package main

import (
	"bufio"
	"encoding/json"
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
	//headerscss := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/css,*/*;q=0.1;\n\n"

	//split up the first line on spaces so we can get the differnt values
	ereq := strings.Split(request_first_line, " ")

	//the request type is the first string
	request_type := ereq[0]

	//log it becaquse it is usefull for debug purpuses
	println("Request Type is " + request_type)

	//if we have a get request or a post request
	if request_type == "GET" && ereq[1] != "/" || request_type == "POST" && ereq[1] != "/" {
		rMap := make(map[string]string)
		JsonData, _ := ioutil.ReadFile("witch.json")

		json.Unmarshal(JsonData, &rMap)
		for route, file := range rMap {
			if ereq[1] == route {
				serverFile(conn, "/"+file)
			}
		}

		serverFile(conn, ereq[1])
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
