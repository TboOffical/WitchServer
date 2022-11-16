package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func exit_listener() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\033[31m\r* Cleaning Up... *\033[0m")
		err := os.RemoveAll("./script-bin")
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}

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

func genereate_err_html(err string, code int) string {
	html := `
		<html>
			<head>
				<title>Witch Server - Error</title>
			</head>
			<body>
				<center>
				<img src="https://raw.githubusercontent.com/TboOffical/WitchServer/main/logo.png" width="50%">
				<h1>` + fmt.Sprint(code) + `, Thats a error</h1>
				<h2>` + err + `</h2>
				</center>
				<style>
					@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@100&family=Roboto:wght@100&display=swap');

					body{
						font-family: 'Montserrat', sans-serif;
					}

				</style>
			</body>
		</html>
	`
	return html
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
		responceNoIndex := generate_status_headers(404) + genereate_err_html("That file could not be located. <br> If you are the owner, create the file. Or route '"+file+"' to a different file <br> in witch.json", 404)
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
