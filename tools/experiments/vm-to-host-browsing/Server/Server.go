package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func browse(browser string, c net.Conn) {
	buf := make([]byte, 512)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}

	data := string(buf[0:nr])
	log.Print("Got ", data)
	cmd := exec.Command(browser, data)
	cmd.Run()
}

func main() {
	browser := flag.String("browser", "firefox", "Browser to use")
	flag.Parse()
	ln, err := net.Listen("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(ln, sigc)
	for {
		fd, _ := ln.Accept()
		browse(*browser, fd)
	}
}
