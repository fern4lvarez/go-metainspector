/*
go-metainspector is a web scraper package that provides access
to basic info and meta tags of a given URL. It includes a command line tool as well.
It is inspired by the metainspector gem (https://github.com/jaimeiniesta/metainspector)
and completely written in Go.
You will find more info on the README file and the `examples` directory.
*/
package metainspector

import (
	"errors"
	"net/url"
	"strings"
)

// root returns the root URL, with no path on it
func root(u *url.URL) string {
	return u.Scheme + "://" + u.Host
}

// fixURL adds the Scheme to the URL and 
// and fixes the Host and Path 
// in case Scheme is empty
func fixURL(u *url.URL) *url.URL {
	if u.Scheme == "" {
		s := strings.Split(u.String(), "/")
		u.Scheme = "http"
		u.Host = s[0]
		u.Path = u.Path[len(s[0]):]
	}
	return u
}

type MetaInspector struct {
	url           string
	scheme        string
	host          string
	rootUrl       string
	title         string
	language      string
	author        string
	description   string
	feed          string
	charset       string
	links         []string
	images        []string
	keywords      []string
	compatibility map[string]string
}

// New creates an object with all the URL scraped info
func New(uri string) (*MetaInspector, error) {
	if uri == "" {
		return nil, errors.New("Url is empty!")
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	scraper, err := newScraper(fixURL(u), 20)
	if err != nil {
		return nil, err
	}

	return &MetaInspector{uri,
		u.Scheme,
		u.Host,
		root(u),
		scraper.title,
		scraper.language,
		scraper.author,
		scraper.description,
		scraper.feed,
		scraper.charset,
		scraper.links,
		scraper.images,
		scraper.keywords,
		scraper.compatibility}, nil
}

// Url provided by the Metainspector
func (m MetaInspector) Url() string {
	return m.url
}

// Scheme of the URL, `http` by default
func (m MetaInspector) Scheme() string {
	return m.scheme
}

// Host of the URL
func (m MetaInspector) Host() string {
	return m.host
}

// RootURL, with no path and slashes
func (m MetaInspector) RootURL() string {
	return m.rootUrl
}

// Title of the site
func (m MetaInspector) Title() string {
	return m.title
}

// Language of the site
func (m MetaInspector) Language() string {
	return m.language
}

// Author of the site
func (m MetaInspector) Author() string {
	return m.author
}

// Description of the site
func (m MetaInspector) Description() string {
	return m.description
}

// Feed link of the site
func (m MetaInspector) Feed() string {
	return m.feed
}

//Charset of the site
func (m MetaInspector) Charset() string {
	return m.charset
}

// Links found on the site
func (m MetaInspector) Links() []string {
	return m.links
}

// Images located on the site
func (m MetaInspector) Images() []string {
	return m.images
}

// Keywords of the site
func (m MetaInspector) Keywords() []string {
	return m.keywords
}

// Compatibility con browsers. specially with IE
func (m MetaInspector) Compatibility() map[string]string {
	return m.compatibility
}
