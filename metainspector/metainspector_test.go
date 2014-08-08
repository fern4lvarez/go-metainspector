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

func ExampleMetaInspector() {
	url := "http://www.cloudcontrol.com/pricing"
	MI, err := New(url)
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
		// Output:
		//URL: http://www.cloudcontrol.com/pricing
		//Scheme: http
		//Host: www.cloudcontrol.com
		//Root: http://www.cloudcontrol.com
		//Title: cloudControl » Cloud App Platform » Pricing
		//Language: en
		//Author: cloudControl GmbH
		//Description: Cloud hosting secure, easy and fair: Highly available and scalable cloud hosting with no administraton hassle and pay as you go billing
		//Charset: utf-8
		//Feed URL: https://www.cloudcontrol.com/blog.rss
		//Links: [http://www.cloudcontrol.com/console/account/{{user.username}} http://www.cloudcontrol.com/ http://www.cloudcontrol.com/pricing http://www.cloudcontrol.com/dev-center http://www.cloudcontrol.com/add-ons http://www.cloudcontrol.com/blog http://www.cloudcontrol.com/console http://www.cloudcontrol.com/pricing/calculator http://www.cloudcontrol.com#included http://www.cloudcontrol.com#memoryhours http://www.cloudcontrol.com#included http://www.cloudcontrol.com#memoryhours http://www.cloudcontrol.com#included http://www.cloudcontrol.com#memoryhours http://www.cloudcontrol.com#included http://www.cloudcontrol.com#memoryhours http://www.cloudcontrol.com/sign-up http://www.cloudcontrol.com/pricing/calculator http://www.cloudcontrol.com/sign-up?plan=Start-up http://www.cloudcontrol.com/pricing/calculator?plan=startup http://www.cloudcontrol.com/sign-up?plan=Business http://www.cloudcontrol.com/pricing/calculator?plan=business http://www.cloudcontrol.com/sign-up?plan=Business%2B http://www.cloudcontrol.com/pricing/calculator?plan=businessplus http://www.cloudcontrol.com/contact http://www.cloudcontrol.com#plantable http://www.cloudcontrol.com#plantable http://www.cloudcontrol.com/dev-center/Quickstart http://www.cloudcontrol.com/dev-center/Platform Documentation http://status.cloudcontrol.com http://www.cloudcontrol.com/dev-center/support http://www.cloudcontrol.com/console http://www.cloudcontrol.com/team http://www.cloudcontrol.com/jobs http://www.cloudcontrol.com/blog http://www.cloudcontrol.com/contact http://www.cloudcontrol.com/add-on-provider-program http://www.cloudcontrol.com/solution-provider-program http://www.whitelabelpaas.info http://www.cloudcontrol.com/tos http://www.cloudcontrol.com/privacy-policy http://www.cloudcontrol.com/imprint]
		//Images: [http://www.cloudcontrol.com/assets/spinner-6f9309f477dcc1d3c21cd21ce44dc8b2.gif]
		//Keywords: [cloudcontrol cloud control cloud hosting cloud computing cloud hosting web-hosting platform as a service paas]
		//Compatibility: map[IE:edge chrome:1]
	}
}
