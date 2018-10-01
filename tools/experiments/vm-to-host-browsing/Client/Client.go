package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"

	"github.com/mitchellh/go-homedir"
)

// VTH: Configuration for vth
type VTH struct {
	Server  string
	Filters []VTHFilter
}

// VTHFilter: Filter configuration
type VTHFilter struct {
	Alias         string
	Filter        string
	Action        string
	ActionCommand []string
}

func main() {
	path, _ := homedir.Dir()
	url := flag.String("url", "", "URL to send")
	fconfig := flag.String("config", fmt.Sprintf("%s%c%s", path, os.PathSeparator, "vth.conf"), "Configuration location for setting up vth client")
	flag.Parse()

	f, err := os.OpenFile(fmt.Sprintf("%s%c%s", path, os.PathSeparator, "vth.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	cfgContent, cfgErr := ioutil.ReadFile(*fconfig)
	if cfgErr != nil {
		panic(cfgErr)
	}

	config := VTH{}
	parseErr := yaml.Unmarshal(cfgContent, &config)
	if parseErr != nil {
		panic(parseErr)
	}

	if *url == "" {
		*url = flag.Arg(0)
	}
	var server string

	// set default
	if config.Server == "" {
		config.Server = "10.0.2.2:7531"
	}
	server = config.Server
	c, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	filterMatch := false
	var matchingFilter VTHFilter
	for _, v := range config.Filters {
		log.Printf("%5s %20s %20s", " -> ", "Processing filter", v.Alias)
		if regexp.MustCompile(v.Filter).MatchString(*url) {
			filterMatch = true
			matchingFilter = v
			break
		}
	}

	if filterMatch {
		log.Printf("%5s %10s|%20s", " -> ", "Match", matchingFilter.Alias)
		if matchingFilter.Action == "drop" {
			log.Printf("Dropping %s", *url)
		} else if matchingFilter.Action == "local" {
			log.Printf("Sending to %s", matchingFilter.ActionCommand[0])
			cmd := exec.Command(matchingFilter.ActionCommand[0], strings.Join(matchingFilter.ActionCommand[1:], " "), *url)
			cmd.Run()
		}
	} else {
		log.Printf("Sending %s", *url)
		c.Write([]byte(*url))
	}
}
