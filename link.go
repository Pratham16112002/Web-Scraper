package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"web_scraper/set"

	"golang.org/x/net/html"
)

type Link struct {
	schema string
	domain string
	path   string
}

type Links []Link

var DeadLinks []string
var ActiveLinks []string

func VisitLinks(input_url string, visited_links *set.HashSet[string]) {
	log.Printf("checking %v\n", input_url)
	visited_links.Add(input_url)
	parsed_url, err := url.Parse(input_url)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("parsed_url %v\n", parsed_url.String())
	resp, err := http.Get(parsed_url.String())
	if err != nil {
		log.Println(err)
		return
	}
	domain := parsed_url.Host
	schema := parsed_url.Scheme
	fmt.Println(domain)
	if 400 <= resp.StatusCode && resp.StatusCode < 600 {
		DeadLinks = append(DeadLinks, parsed_url.String())
	} else {
		ActiveLinks = append(ActiveLinks, parsed_url.String())
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal("HTML Parsing error : ", err)
		return
	}
	recursive_links := Links{}
	findAnchorTags(&recursive_links, doc)
	for _, link := range recursive_links {
		var new_link string
		var err error
		if link.domain == "" {
			link.domain = domain
			if link.schema == "" {
				link.schema = schema
				new_link, err = url.JoinPath(link.schema+"://", link.domain, link.path)
				if err != nil {
					log.Println("Error validing the URL", link.path)
					continue
				}
				log.Println(new_link)
			}
		}
		if !visited_links.Contains(new_link) {
			VisitLinks(new_link, visited_links)
		}
	}
}

func findAnchorTags(anchorLinks *Links, doc *html.Node) {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		for _, attr := range doc.Attr {
			if attr.Key == "href" {
				parsed_url, err := url.Parse(attr.Val)
				if err != nil {
					log.Println("Error parsing : ", parsed_url.String())
					break
				}
				*anchorLinks = append(*anchorLinks, Link{
					schema: parsed_url.Scheme,
					domain: parsed_url.Host,
					path:   parsed_url.Path,
				})
				break
			}
		}
	}
	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		findAnchorTags(anchorLinks, child)
	}
}
