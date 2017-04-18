# go-metainspector [![Build Status](https://travis-ci.org/fern4lvarez/go-metainspector.png)](https://travis-ci.org/fern4lvarez/go-metainspector)[![GoDoc](https://godoc.org/github.com/fern4lvarez/go-metainspector/metainspector?status.png)](http://godoc.org/github.com/fern4lvarez/go-metainspector/metainspector)

Simple web metadata scraping in Go.

**go-metainspector** is a web scraper package that provides access
to basic info and meta tags of a given URL.
It is inspired by the [metainspector gem](https://github.com/jaimeiniesta/metainspector) by [Jaime Iniesta](https://twitter.com/jaimeiniesta) and completely written in Go.

## Install

* Step 1: Get the `metainspector` package

```
go get -u github.com/fern4lvarez/go-metainspector/metainspector
```

* Step 2 (Optional): Run tests

```
$ go test -v -cover ./...
```

## Usage

### API

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
  }
```

### CLI
You can use the go-metainspector as a command line tool.
*Note:* Once you have installed the package, you might want to set an alias for the CLI: add `alias metainspect=go-metainspector` or similar in your appropiated dotfile.

Simple to use!

```
$ go-metainspector -help
Usage of go-metainspector:
  -u="www.example.com": URL to metainspect.
  -all=false: Show full results.

$ go-metainspector -u www.cloudcontrol.com
www.cloudcontrol.com
----> Title: cloudControl » Cloud App Platform » supercharging development
----> Author: cloudControl GmbH
----> Description: Cloud hosting secure, easy and fair: Highly available and scalable cloud hosting with no administraton hassle and pay as you go billing
----> Charset: utf-8
----> Language: en
----> Feed URL: https://www.cloudcontrol.com/blog/feed
----> Keywords: cloudcontrol cloud control cloud hosting cloud computing cloud hosting web-hosting platform as a service paas 
----> Links: http://www.cloudcontrol.com/ http://www.cloudcontrol.com/pricing http://www.cloudcontrol.com/dev-center http://www.cloudcontrol.com/add-ons https://console.cloudcontrolled.com http://www.cloudcontrol.com/for-isvs http://www.cloudcontrol.com/for-developers http://www.cloudcontrol.com/use-cases/vamos http://www.cloudcontrol.com/use-cases/searchbox http://www.cloudcontrol.com/use-cases/snipclip ...
----> Images: http://www.cloudcontrol.com/assets/spinner-9bde1d21899a52974160da652c0a6622.gif https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/006/thumb/vamos-logo-rgb-vertical.png?1360313785 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/008/thumb/logo.png?1362736875 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/002/thumb/snipclip_logo_2011_rgb.png?1359481602 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/004/thumb/ormigo.png?1359481203 https://s3-eu-west-1.amazonaws.com/cctrl-www-assets/add-ons/blitz.png https://s3-eu-west-1.amazonaws.com/cctrl-www-assets/add-ons/sendgrid.png https://s3-eu-west-1.amazonaws.com/cctrl-www-production/solution_providers/logos/000/000/013/small/emind_klein.png?1364402125 ...
```

`-u` is optional though:

```
$ go-metainspector www.cloudcontrol.com
----> Title: cloudControl » Cloud App Platform » supercharging development
----> Author: cloudControl GmbH
...
```

Sometimes one site contains so many links and images that makes it ugly to print them all. By default it's cut on 10 links and images. Use `-all` to print them all:

```
$ go-metainspector -u www.cloudcontrol.com -all
----> Title: cloudControl » Cloud App Platform » supercharging development
----> Author: cloudControl GmbH
----> Description: Cloud hosting secure, easy and fair: Highly available and scalable cloud hosting with no administraton hassle and pay as you go billing
----> Charset: utf-8
----> Language: en
----> Feed URL: https://www.cloudcontrol.com/blog/feed
----> Keywords: cloudcontrol cloud control cloud hosting cloud computing cloud hosting web-hosting platform as a service paas 
----> Links: http://www.cloudcontrol.com/ http://www.cloudcontrol.com/pricing http://www.cloudcontrol.com/dev-center http://www.cloudcontrol.com/add-ons https://console.cloudcontrolled.com http://www.cloudcontrol.com/for-isvs http://www.cloudcontrol.com/for-developers http://www.cloudcontrol.com/add-ons/sendgrid http://www.cloudcontrol.com/add-ons/newrelic http://www.cloudcontrol.com/use-cases/ormigo http://www.cloudcontrol.com/use-cases/vamos http://www.cloudcontrol.com/use-cases/afp http://www.cloudcontrol.com/use-cases/adcloud http://www.cloudcontrol.com/use-cases/kinderfee http://www.cloudcontrol.com/use-cases/snipclip http://www.cloudcontrol.com/dev-center/Quickstart http://www.cloudcontrol.com/dev-center/Platform Documentation http://status.cloudcontrol.com http://www.cloudcontrol.com/dev-center/support https://console.cloudcontrolled.com http://www.cloudcontrol.com/team http://www.cloudcontrol.com/jobs http://www.cloudcontrol.com/blog http://www.cloudcontrol.com/contact http://www.cloudcontrol.com/add-on-provider-program http://www.cloudcontrol.com/solution-providers http://www.cloudcontrol.com/tos http://www.cloudcontrol.com/privacy-policy http://www.cloudcontrol.com/imprint 
----> Images: http://www.cloudcontrol.com/assets/spinner-9bde1d21899a52974160da652c0a6622.gif https://s3-eu-west-1.amazonaws.com/cctrl-www-assets/add-ons/sendgrid.png https://s3-eu-west-1.amazonaws.com/cctrl-www-assets/add-ons/newrelic.png https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/004/thumb/ormigo.png?1359481203 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/006/thumb/vamos-logo-rgb-vertical.png?1360313785 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/003/thumb/AFP-logo.png?1359481038 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/001/thumb/adcloud.png?1359563276 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/005/thumb/Kinderfee_Logo_Claim.png?1359481378 https://s3-eu-west-1.amazonaws.com/cctrl-www-production/use_cases/logos/000/000/002/thumb/snipclip_logo_2011_rgb.png?1359481602 
```

### Web
You can use go-metainspector via web. Check out and run the [`go-metainspector-site`](https://github.com/fern4lvarez/go-metainspector-site) or just try it out: http://gometainspector.cloudcontrolled.com

## Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Everything is appreciated!

## TODO (aka Nice To Have)
* Extend documentation
* Write a CHANGELOG
* Mock http requests to speed up unit tests
* Internal links, External links
* Map() method wrapping all data
* Set Timeout optionally (now is 20 secs)
* Your suggestion <HERE>

## License
go-metainspector is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector/blob/master/LICENSE)