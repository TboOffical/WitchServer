package main

//start by importing the nessery libs for the app to work properly
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/gen2brain/dlgs"
)

type Certificate struct {
	EnableTLS bool
	Crt_file  string
	Key_file  string
}

func main() {
	//call the exit listener
	exit_listener()

	//find the number of args there are
	argLength := len(os.Args[1:])

	//If ther program is exe'd with the help argument
	//display basic help
	if argLength == 1 && os.Args[1] == "help" {
		println("To start using witch, simply execute")
		println("./witch <port>")
		println("To set cutsom routes, create a witch.json file with the following structure")
		println(`
{
	"/route" : "something.html"
}
			`)
		println("To load a certificate, create cert.json with the following structure")
		println(`
{
	"enableTLS": true,
	"crt_file": "localhost.cert",
	"key_file": "localhost.key"
}
		`)
		return
	}

	//defineing some colors to use
	colorGreen := "\033[32m"
	colorCyan := "\033[36m"

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

	witch_script_in_use := 0
	witch_scripts := []string{}

	files, _ := ioutil.ReadDir(".")
	for file := range files {
		if strings.Contains(files[file].Name(), ".wba") {
			witch_script_in_use = 1
			witch_scripts = append(witch_scripts, files[file].Name())
		}
	}

	fmt.Print("\033[H\033[2J")
	println("\033[31m\r* Witch Pre-Init Process *\033[0m")

	if witch_script_in_use == 1 {
		os.Mkdir("./script-bin", 0755)
		for script := range witch_scripts {
			println("Compileing Script : " + witch_scripts[script])
			script_data, err := ioutil.ReadFile(witch_scripts[script])
			if err != nil {
				println("Failed!")
				break
			}
			script_temp := RandomString(20)
			ioutil.WriteFile(script_temp+".go", []byte(`
			package main
			import "os"
			`+string(script_data)+
				`
			const (
				request_get     = "get"
				request_post    = "post"
				requets_unknown = "unknown"
			)
			
			func request_type() string {
				if len(os.Args) < 2 {
					return "unknown"
				}
				if os.Args[1] == "get" {
					return "get"
				}
				if os.Args[1] == "post" {
					return "post"
				}
				return "unknown"
			}			

			`), 0644)
			cmd, err := exec.Command("go", "build", "-o", "./script-bin/"+witch_scripts[script]+".exe", script_temp+".go").Output()
			if err != nil {
				println("Failed!")
			}
			println(string(cmd))
			os.Remove(script_temp + ".go")
		}

		fmt.Print("\033[H\033[2J")
	}

	//load the witch.json config
	LoadConfig()

	//FANCY
	println(" _       ____________________  __	")
	println("| |     / /  _/_  __/ ____/ / / /	")
	println("| | /| / // /  / / / /   / /_/ / 	")
	println("| |/ |/ // /  / / / /___/ __  /  	")
	println("|__/|__/___/ /_/  /____/_/ /_/   	\n")

	//tell the user that we are started up
	println(string(colorGreen) + "Witch Started")

	println(string(colorCyan) + "Now Listening For new Connections On Port " + _PORT + "\033[0m")

	if certFail == 1 {
		println("Acccess The Server by navigateing to http://" + _host + ":" + _PORT + " In Your Web Browser")
	} else {
		println("Acccess The Server by navigateing to https://" + _host + ":" + _PORT + " In Your Web Browser")

	}

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
