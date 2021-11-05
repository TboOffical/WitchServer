package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/dlgs"
)

func main() {
	argLength := len(os.Args[1:])

	if argLength == 1 && os.Args[1] == "help" {
		println("To start using witch, simply execute")
		println("./witch <port>")
		return
	}

	_PORT := ""
	_host := "localhost"

	if argLength == 0 {
		port, _, _ := dlgs.Entry("Witch", "Please enter a port to start witch on", "9000")
		if port == "" {
			dlgs.Info("Witch", "Please Enter A Value")
			return
		}
		_PORT = port
	} else {
		_PORT = os.Args[1]
	}

	l, err := net.Listen("tcp4", _host+":"+_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	println(" _       ____________________  __	")
	println("| |     / /  _/_  __/ ____/ / / /	")
	println("| | /| / // /  / / / /   / /_/ / 	")
	println("| |/ |/ // /  / / / /___/ __  /  	")
	println("|__/|__/___/ /_/  /____/_/ /_/   	")

	println("Starting Witch...")
	defer l.Close()

	println("Now Listening For new Connections On Port " + _PORT)
	println("Acccess The Server by navigateing to http://localhost:" + _PORT + " In Your Web Browser")
	println("----------------------------------------------------------------------------")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("A Web Web Browser has connected -> " + c.RemoteAddr().String())
		go handleConnection(c)
	}

}

func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {

			return s[i:]
		}
	}
	return ""
}

func handleConnection(conn net.Conn) {
	dt := time.Now()
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	headers := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\n\n"
	headerscss := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/css,*/*;q=0.1;\n\n"

	if err != nil {
		fmt.Println("Client Has Dissconnected")
		conn.Close()
		return
	}

	log.Println("Client Request -> ", string(buffer[:len(buffer)-1]))
	req_str := string(buffer[:len(buffer)-1])

	ereq := strings.Split(req_str, " ")

	request_type := ereq[0]

	if request_type == "GET" && ereq[1] != "/" {
		println("Loading In file -> " + ereq[1])
		content, err := ioutil.ReadFile(trimFirstRune(ereq[1]))
		if err != nil {
			fmt.Println("No File Found " + ereq[1])
			responceNoIndex := headers + "<title>Witch -> Cant Find file</title><center><img style='border-top: groove;' src='https://i.imgur.com/j5GlneF.png' </center>"
			conn.Write([]byte(responceNoIndex))
			conn.Close()
			return
		}
		responce_text := headers + string(content)
		//println(strings.Contains(ereq[1], ".css"))
		if strings.Contains(ereq[1], ".css") {
			responce_texts := headerscss + string(content)
			conn.Write([]byte(responce_texts))
			conn.Close()
			return
		}

		conn.Write([]byte(responce_text))
		conn.Close()
		return
	}

	//check if and load index
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("No Index.html File Found")
		responceNoIndex := headers + "<title>Witch -> Cant Find file</title><center><img style='border-top: groove;' src='https://i.imgur.com/GJIUMVC.png' </center>"
		conn.Write([]byte(responceNoIndex))
		conn.Close()
		return
	}
	//Found Index.html Sending

	TextRES := "HTTP/1.1 200 OK\nDate:" + dt.String() + "\nServer:WitchFX\nContent-Type: text/html;\n\n" + string(content)
	conn.Write([]byte(TextRES))

	conn.Close()
}
