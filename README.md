# go-metainspector [![Build Status](https://travis-ci.org/fern4lvarez/go-metainspector.png)](https://travis-ci.org/fern4lvarez/go-metainspector) 
[Documentation online](http://godoc.org/github.com/fern4lvarez/go-metainspector/metainspector)

Simple web scraping for Go.

**go-metainspector** is a web scraper package that provides access
to basic info and meta tags of a given URL.
It is inspired by the [metainspector gem](https://github.com/jaimeiniesta/metainspector) by [Jaime Iniesta](https://twitter.com/jaimeiniesta) and completely written in Go.

## Install (with GOPATH set on yout machine)

* Step 1: Get the `metainspector` package

```
go get github.com/fern4lvarez/go-metainspector/metainspector
```

* Step 2 (Optional): Run tests

```
$ go test -v ./...
```

##Usage
```
package main

import (
  "fmt"
  "github.com/fern4lvarez/go-metainspector/metainspector"
)

func main() {
  url := "http://www.cloudcontrol.com/pricing"
  MI, err := metainspector.New(url)
  if err != nil {
    fmt.Printf("Error: %v", err)
  } else {
    fmt.Printf("\nURL: %s\n", MI.Url())
    fmt.Printf("Scheme: %s\n", MI.Scheme())
    fmt.Printf("Host: %s\n", MI.Host())
    fmt.Printf("Root: %s\n", MI.RootURL())
    fmt.Printf("Title: %s\n", MI.Title())
    fmt.Printf("Language: %s\n", MI.Language())
    fmt.Printf("Author: %s\n", MI.Author())
    fmt.Printf("Description: %s\n", MI.Description())
    fmt.Printf("Charset: %s\n", MI.Charset())
    fmt.Printf("Feed URL: %s\n", MI.Feed())
    fmt.Printf("Links: %v\n", MI.Links())
    fmt.Printf("Images: %v\n", MI.Images())
    fmt.Printf("Keywords: %v\n", MI.Keywords())
    fmt.Printf("Compatibility: %v\n", MI.Compatibility())
    // URL: http://www.cloudcontrol.com/pricing
    // Scheme: http
    // Host: www.cloudcontrol.com
    // Root: http://www.cloudcontrol.com
    // Title: cloudControl » Cloud App Platform » Pricing & Sign-Up
    // Language: en
    // Author: cloudControl GmbH
    // Description: Cloud hosting secure, easy and fair: Highly available and scalable cloud hosting with no administraton hassle and pay as you go billing
    // Charset: utf-8
    // Feed URL: https://www.cloudcontrol.com/blog/feed
    // Links: [http://www.cloudcontrol.com/ http://www.cloudcontrol.com/pricing http://www.cloudcontrol.com/dev-center http://www.cloudcontrol.com/add-ons https://console.cloudcontrolled.com http://www.cloudcontrol.com/pricing/modal/load_template_free http://www.cloudcontrol.com/pricing/modal/load_template_small http://www.cloudcontrol.com/pricing/modal/load_template_medium http://www.cloudcontrol.com/pricing/modal/load_template_large http://www.cloudcontrol.com/pricing/modal/web_container_show http://www.cloudcontrol.com/pricing/modal/worker_container_show http://www.cloudcontrol.com/pricing/modal/addon_show http://www.cloudcontrol.com/pricing/modal/support_show https://console.cloudcontrolled.com/register https://console.cloudcontrolled.com/register http://www.cloudcontrol.com/contact http://www.cloudcontrol.com/dev-center/Quickstart http://www.cloudcontrol.com/dev-center/Platform Documentation http://status.cloudcontrol.com http://www.cloudcontrol.com/dev-center/support https://console.cloudcontrolled.com http://www.cloudcontrol.com/team http://www.cloudcontrol.com/jobs http://www.cloudcontrol.com/blog http://www.cloudcontrol.com/contact http://www.cloudcontrol.com/add-on-provider-program http://www.cloudcontrol.com/solution-providers http://www.cloudcontrol.com/tos http://www.cloudcontrol.com/privacy-policy http://www.cloudcontrol.com/imprint]
    // Images: [http://www.cloudcontrol.com/assets/spinner-9bde1d21899a52974160da652c0a6622.gif]
    // Keywords: [cloudcontrol cloud control cloud hosting cloud computing cloud hosting web-hosting platform as a service paas]
    // Compatibility: map[IE:edge chrome:1]
  }
```

##Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Everything is appreciated!

##TODO (aka Nice To Have)
* Extend documentation
* Write a CHANGELOG
* Command line interface
* Example site (written in Go, of course)
* Mock http requests to speed up unit tests
* Internal links, External links
* Map() method wrapping all data
* Set Timeout optionally (now is 20 secs)
* Your suggestion <HERE>

##License
go-metainspector is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector/blob/master/LICENSE)