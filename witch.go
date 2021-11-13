package main

//start by importing the nessery libs for the app to work properly
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/gen2brain/dlgs"
)

type Certificate struct {
	EnableTLS bool
	Crt_file  string
	Key_file  string
}

func main() {
	//find the number of args there are
	argLength := len(os.Args[1:])

	//If ther program is exe'd with the help argument
	//display basic help
	if argLength == 1 && os.Args[1] == "help" {
		println("To start using witch, simply execute")
		println("./witch <port>")
		return
	}

	//Set the port and host varibals
	//so we can span a tcp listener
	_PORT := ""
	_host := "localhost"
	certFail := 0
	var cert tls.Certificate

	//shhhhhh
	_ = cert

	//if we did not input any args, display a gui box asking us for a port
	// I know nothing about gui stuff so i used this lib : gen2brain/dlgs
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

	var jsonCert Certificate
	rawJson, err := ioutil.ReadFile("cert.json")
	if err != nil {
		certFail = 1

	} else {
		err = json.Unmarshal(rawJson, &jsonCert)
		if err != nil {
			println("Failed to read json")
		}

		cert, err = tls.LoadX509KeyPair(jsonCert.Crt_file, jsonCert.Key_file)
		if err != nil {
			certFail = 1
			fmt.Println(err)
		}
	}

	//get a TCP4 listener goin

	var l net.Listener

	if certFail != 1 {
		config := &tls.Config{Certificates: []tls.Certificate{cert}}
		l, err = tls.Listen("tcp4", _host+":"+_PORT, config)
	} else {
		l, err = net.Listen("tcp4", _host+":"+_PORT)
	}

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	//FANCY
	println(" _       ____________________  __	")
	println("| |     / /  _/_  __/ ____/ / / /	")
	println("| | /| / // /  / / / /   / /_/ / 	")
	println("| |/ |/ // /  / / / /___/ __  /  	")
	println("|__/|__/___/ /_/  /____/_/ /_/   	\n")

	//tell the user that we are started up
	println("Witch Started")

	println("Now Listening For new Connections On Port " + _PORT)
	println("Acccess The Server by navigateing to http://localhost:" + _PORT + " In Your Web Browser")

	if certFail == 1 {
		println("No certificate was loaded")
	} else {
		println("Loaded certificate sucessfully")
	}

	println("----------------------------------------------------------------------------")

	//for ever connection spawn a new listener
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
