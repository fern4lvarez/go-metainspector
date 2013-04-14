# go-metainspector [![Build Status](https://travis-ci.org/fern4lvarez/go-metainspector.png)](https://travis-ci.org/fern4lvarez/go-metainspector) [Documentation online](http://godoc.org/github.com/fern4lvarez/go-metainspector/metainspector)

Simple web scrapping for Go.

**go-metainspector** is a web scraper package that provides access
to basic info and meta tags of a given URL.
It is inspired by the [metainspector gem](https://github.com/jaimeiniesta/metainspector) by [Jaime Iniesta](https://twitter.com/jaimeiniesta) and completely written in Go.

## Install (with GOPATH set on yout machine)

#### Note: This package relies on the [go-html-transform](http://code.google.com/p/go-html-transform) package's latest source, although the `go1` tag is set to an old commit, so it's necessary to checkout the repository in order to get the source code compatible with this package. As soon as the `go1` tag is updated, this package will be available just using `go get`. So far is needed to follow these steps to install it properly. 

* Step 1: Change dir to your GOPATH

```
$ cd GOPATH
```

* Step 2: Install ´go.net/html´

```
$ go get code.google.com/p/go.net/html
```

* Step 3: Install Mercurial

* Step 4: Install ´go-html-transform/5´ and checkout to the default branch

```
$ go get code.google.com/p/go-html-transform/h5
$ cd GOPATH/src/code.google.com/p/go-html-transform
$ hg checkout default
$ cd GOPATH
```

* Step 5: Get the `go-metainspector` package

```
go get github.com/fern4lvarez/go-metainspector/metainspector
```

* Step 6 (Optional): Run tests

```
$ cd GOPATH/src/github.com/fern4lvarez/go-metainspector
$ go test -v ./...
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
    fmt.Printf("\nThe url is: %s, the scheme is %s, the host is %s, root url is %s. Title is %s, written in %s, author is %s, description is %s. Charset is %s.",
      MI.Url(),
      MI.Scheme(),
      MI.Host(),
      MI.RootURL(),
      MI.Title(),
      MI.Language(),
      MI.Author(),
      MI.Description(),
      MI.Charset())
    fmt.Printf("\nSubscribe now! ->%s", MI.Feed())
    fmt.Printf("\nThe links are: %#v", MI.Links())
    fmt.Printf("\nThe images are: %#v", MI.Images())
    fmt.Printf("\nThe keywords are: %#v", MI.Keywords())
    fmt.Printf("\nCompatibility: %#v", MI.Compatibility())
  }
}


```

##Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Anything is appreciated!

##TODO (aka Nice To Have)
* Example site as SaaS
* Mock http requests to speed up unit tests
* Internal links, External links
* Map() method wrapping all data
* Set Timeout optionally (now is 20 secs)
* Your suggestion <HERE>

##License
go-metainspector is MIT licensed, see [here](https://github.com/fern4lvarez/go-metainspector/blob/master/LICENSE)