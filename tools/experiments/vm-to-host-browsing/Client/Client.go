package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/mitchellh/go-homedir"
)

func main() {
	url := flag.String("url", "", "URL to send")
	sock := flag.String("server", "10.0.2.2:7531", "Host to connect to")
	flag.Parse()
	if *url == "" {
		*url = flag.Arg(0)
	}
	var server string
	path, _ := homedir.Dir()
	if *sock == "" {
		s, _ := ioutil.ReadFile(fmt.Sprintf("%s%s%s", path, os.PathSeparator, "vth.conf"))
		server = string(s)
	} else {
		server = *sock
		ioutil.WriteFile(fmt.Sprintf("%s%s%s", path, os.PathSeparator, "vth.conf"), []byte(*sock), 777)
	}
	c, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()
	log.Print("Sending ", *url)
	c.Write([]byte(*url))
}
