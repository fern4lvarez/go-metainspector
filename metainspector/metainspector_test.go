package metainspector

import (
	"fmt"
	"net/url"
	"testing"
)

var msgFail = "%v function fails. Expects %v, returns %v"

func TestRoot(t *testing.T) {
	u, err := url.Parse("http://www.cloudcontrol.com/pricing")
	if err != nil {
		t.Errorf("%v", err)
	}
	if rootUrl := ExportRoot(u); rootUrl != "http://www.cloudcontrol.com" {
		t.Errorf(msgFail, "rootURL", "http://www.cloudcontrol.com", rootUrl)
	}
}

func TestFixURL(t *testing.T) {
	u, err := url.Parse("http://foo.bar")
	if err != nil {
		t.Errorf("%v", err)
	}
	u.Scheme, u.Host, u.Path = "", "", "www.cloudcontrol.com/pricing"
	fix := ExportFixURL(u)
	if fix.Scheme != "http" {
		t.Errorf(msgFail, "Scheme", "http", fix.Scheme)
	} else if fix.Host != "www.cloudcontrol.com" {
		t.Errorf(msgFail, "Host", "www.cloudcontrol.com", fix.Host)
	} else if fix.Path != "/pricing" {
		t.Errorf(msgFail, "Path", "/pricing", fix.Path)
	}
}

var mi = MetaInspector{url: "http://www.cloudcontrol.com",
	scheme:        "http",
	host:          "www.cloudontrol.com",
	rootUrl:       "http://www.cloudcontrol.com",
	title:         "CloudControl",
	language:      "en",
	author:        "CloudControl",
	description:   "PaaS company",
	generator:     "some generator",
	feed:          "http://www.cloudcontrol.com/feed",
	charset:       "utf-8",
	links:         []string{"http://foo.bar", "https://bar.foo"},
	images:        []string{"http://foo.jpg", "https://bar.png"},
	keywords:      []string{"cloud", "PaaS"},
	compatibility: map[string]string{"IE": "edge"},
}

func TestUrl(t *testing.T) {
	if url := mi.Url(); url != "http://www.cloudcontrol.com" {
		t.Errorf(msgFail, "Url", "www.google.com", url)
	}
}

func TestScheme(t *testing.T) {
	if scheme := mi.Scheme(); scheme != "http" {
		t.Errorf(msgFail, "Scheme", "http", scheme)
	}
}

func TestHost(t *testing.T) {
	if host := mi.Host(); host != "www.cloudontrol.com" {
		t.Errorf(msgFail, "Host", "www.cloudontrol.com", host)
	}
}

func TestRootURL(t *testing.T) {
	if rootUrl := mi.RootURL(); rootUrl != "http://www.cloudcontrol.com" {
		t.Errorf(msgFail, "RootURL", "http://www.cloudontrol.com", rootUrl)
	}
}

func TestTitle(t *testing.T) {
	if title := mi.Title(); title != "CloudControl" {
		t.Errorf(msgFail, "Title", "CloudControl", title)
	}
}

func TestLanguage(t *testing.T) {
	if language := mi.Language(); language != "en" {
		t.Errorf(msgFail, "Language", "en", language)
	}
}

func TestAuthor(t *testing.T) {
	if author := mi.Author(); author != "CloudControl" {
		t.Errorf(msgFail, "Author", "CloudControl", author)
	}
}

func TestDescription(t *testing.T) {
	if description := mi.Description(); description != "PaaS company" {
		t.Errorf(msgFail, "Description", "PaaS company", description)
	}
}

func TestGenerator(t *testing.T) {
	if generator := mi.Generator(); generator != "some generator" {
		t.Errorf(msgFail, "Generator", "some generator", generator)
	}
}

func TestFeed(t *testing.T) {
	if feed := mi.Feed(); feed != "http://www.cloudcontrol.com/feed" {
		t.Errorf(msgFail, "Feed", "http://www.cloudcontrol.com/feed", feed)
	}
}

func TestCharset(t *testing.T) {
	if charset := mi.Charset(); charset != "utf-8" {
		t.Errorf(msgFail, "Charset", "utf-8", charset)
	}
}

func TestLinks(t *testing.T) {
	l1 := fmt.Sprintf("%v", mi.Links())
	l2 := "[http://foo.bar https://bar.foo]"
	if l1 != l2 {
		t.Errorf(msgFail, "Links", l2, l1)
	}
}

func TestImages(t *testing.T) {
	i1 := fmt.Sprintf("%v", mi.Images())
	i2 := "[http://foo.jpg https://bar.png]"
	if i1 != i2 {
		t.Errorf(msgFail, "Images", i2, i1)
	}
}

func TestKeywords(t *testing.T) {
	k1 := fmt.Sprintf("%v", mi.Keywords())
	k2 := "[cloud PaaS]"
	if k1 != k2 {
		t.Errorf(msgFail, "Keywords", k2, k1)
	}
}

func TestCompatibility(t *testing.T) {
	c1 := fmt.Sprintf("%v", mi.Compatibility())
	c2 := "map[IE:edge]"
	if c1 != c2 {
		t.Errorf(msgFail, "Compatibility", c2, c1)
	}
}
