package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"regexp"
	"web_scraper/set"
)

type CmdFlags struct {
	Url string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Url, "url", "", "URL")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(deadLinks *Links) {
	switch {
	case cf.Url != "":
		re, err := regexp.Compile(cf.Url)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		parsed_url, err := url.Parse(re.String())
		if err != nil {
			log.Println("URL parsing error : ", err)
			os.Exit(1)
		}
		visited_links := set.New[string]()
		VisitLinks(parsed_url.String(), visited_links)
	default:
		log.Println("Invalid option")
		os.Exit(1)
	}
}
