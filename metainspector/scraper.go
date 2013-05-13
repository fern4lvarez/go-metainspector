package metainspector

import (
	"code.google.com/p/go-html-transform/h5"
	"code.google.com/p/go.net/html"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type scraper struct {
	title         string
	language      string
	author        string
	description   string
	generator     string
	feed          string
	charset       string
	links         []string
	images        []string
	keywords      []string
	compatibility map[string]string
}

// findCharset returns the charset given the Content-Type
func findCharset(content string) string {
	if pos := strings.LastIndex(content, "charset="); pos != -1 {
		return content[pos+len("charset="):]
	}
	return ""
}

// mapifyStr converts a string with the format `foo=bar,ro=pi` 
// in a map like map[foo:bar ro:pi]
func mapifyStr(content string) map[string]string {
	m := make(map[string]string)
	a := strings.Split(content, ",")
	for i := range a {
		s := strings.Split(a[i], "=")
		m[s[0]] = s[1]
	}
	return m
}

// findAttribute returns the value of the key provided
// within the attributes of the given Node
func findAttribute(n *html.Node, key string) string {
	if a := n.Attr; a != nil {
		for i := range a {
			if a[i].Key == key {
				return a[i].Val
			}
		}
	}
	return ""
}

// addElement adds a string to a given string slice if matches the 
// given attribute within the given Node.
// If starts with slash means that is a external element, so is appended to 
// the root URL
func addElement(elems []string, u *url.URL, n *html.Node, attr string) []string {
	if val := findAttribute(n, attr); val != "" {
		if strings.Index(val, "/") == 0 {
			val = u.Scheme + "://" + u.Host + val
		}
		elems = append(elems, val)
	}
	return elems
}

// timeoutDialer sets a connection with a given timeout
func timeoutDialer(secs int) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		c, err := net.Dial(netw, addr)
		if err != nil {
			return nil, err
		}
		c.SetDeadline(time.Now().Add(time.Duration(secs) * time.Second))
		return c, nil
	}
}

// newScraper parses the given url and scrapes the site, returning 
// a scraper object with all the site info and meta tags
func newScraper(u *url.URL, timeout int) (*scraper, error) {
	var title string
	var language string
	var author string
	var description string
	var generator string
	var feed string
	charset := "utf-8"
	links := make([]string, 0)
	images := make([]string, 0)
	keywords := make([]string, 0)
	compatibility := make(map[string]string)

	scrpr := func(n *html.Node) {
		switch n.Data {
		case "html":
			language = findAttribute(n, "lang")
		case "title":
			title = n.FirstChild.Data
		case "a":
			links = addElement(links, u, n, "href")
		case "img":
			images = addElement(images, u, n, "src")
		case "link":
			typ := findAttribute(n, "type")
			switch typ {
			case "application/rss+xml":
				feed = findAttribute(n, "href")
			}
		case "meta":
			name := findAttribute(n, "name")
			switch name {
			case "author":
				author = findAttribute(n, "content")
			case "keywords":
				keywords = strings.Split(findAttribute(n, "content"), ", ")
			case "description":
				description = findAttribute(n, "content")
			case "generator":
				generator = findAttribute(n, "content")
			}

			httpEquiv := findAttribute(n, "http-equiv")
			switch httpEquiv {
			case "Content-Type":
				charset = findCharset(findAttribute(n, "content"))
			case "X-UA-Compatible":
				compatibility = mapifyStr(findAttribute(n, "content"))
			}
		}
	}

	cl := http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(timeout),
		},
	}

	resp, err := cl.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tree, err := h5.New(resp.Body)
	if err != nil {
		return nil, err
	}
	tree.Walk(scrpr)

	return &scraper{title,
		language,
		author,
		description,
		generator,
		feed,
		charset,
		links,
		images,
		keywords,
		compatibility}, nil
}
