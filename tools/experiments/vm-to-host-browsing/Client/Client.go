package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	c, err := net.Dial("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()
	url := flag.String("url", "", "URL to send")
	flag.Parse()
	log.Print("Sending ", *url)
	c.Write([]byte(*url))
}
