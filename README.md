# go-metainspector [Documentation online](http://godoc.org/github.com/fern4lvarez/go-metainspector)

Simple web scrapping for Go.

**go-metainspector** is a web scraper package that provides access
to basic info and meta tags of a given URL.
It is inspired by the [metainspector gem](https://github.com/jaimeiniesta/metainspector) by [Jaime Iniesta](https://twitter.com/jaimeiniesta) and written completely in Go.

## Install (with GOPATH set on yout machine)

```
//Install the dependencies
go get code.google.com/p/go.net/html
go get code.google.com/p/go-html-transform/h5
```

```
// Get the package
go get github.com/fern4lvarez/go-metainspector
```

##Usage
```
package main

import (
  "fmt"
  mi "github.com/fern4lvarez/go-metainspector"
)

func main() {
  url := "http://www.cloudcontrol.com"

  MI, err := mi.NewMetaInspector(url)
  if err != nil {
    fmt.Printf("Error: %v", err)
  } else {
    fmt.Printf("\nThe url is: %s, the scheme is %s, the host is %s, root url is %s. Title is %s, written in %s, author is %s, description is %s. Charset is %s.", MI.Url(), MI.Scheme(), MI.Host(), MI.RootURL(), MI.Title(), MI.Language(), MI.Author(), MI.Description(), MI.Charset())

    fmt.Printf("\nSubscribe now! ->%s", MI.Feed())
    fmt.Printf("\nThe links are: %#v", MI.Links())
    fmt.Printf("\nThe images are: %#v", MI.Images())
    fmt.Printf("\nThe keywords are: %#v", MI.Keywords())
    fmt.Printf("\nCompatibility: %#v", MI.Compatibility())
  }
}

```

##TODO (aka Nice To Have)
* Mock http requests to speed up unit tests
* Internal links, External links
* Map() method wrapping all data
* Set Timeout optionally (now is 20 secs)

##License
go-metainspector is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector/blob/master/LICENSE)