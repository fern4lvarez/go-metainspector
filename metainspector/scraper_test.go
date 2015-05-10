package metainspector

import (
	"fmt"
	"net/url"
	"testing"

	"golang.org/x/net/html"
)

func TestFindCharset(t *testing.T) {
	content := "text/html;charset=utf-8"
	if charset := ExportFindCharset(content); charset != "utf-8" {
		t.Errorf(msgFail, "Charset", "utf-8", charset)
	}

	content = "text/html"
	if charset := ExportFindCharset(content); charset != "" {
		t.Errorf(msgFail, "Charset", "", charset)
	}
}

func TestMapifyStr(t *testing.T) {
	content := "foo=3,bar=sh,roo=pi"
	m := ExportMapifyStr(content)
	if m["foo"] != "3" {
		t.Errorf(msgFail, "m[foo]", "3", m["foo"])
	} else if m["bar"] != "sh" {
		t.Errorf(msgFail, "m[bar]", "sh", m["bar"])
	} else if m["roo"] != "pi" {
		t.Errorf(msgFail, "m[roo]", "pi", m["roo"])
	}
}

func TestFindAttribute(t *testing.T) {
	n := &html.Node{
		Type:      1,
		DataAtom:  1,
		Data:      "meta",
		Namespace: "",
		Attr: []html.Attribute{
			html.Attribute{
				Namespace: "",
				Key:       "name",
				Val:       "author"},
			html.Attribute{
				Namespace: "",
				Key:       "content",
				Val:       "cloudControl GmbH"}}}
	key := "content"

	if at := ExportFindAttribute(n, key); at != "cloudControl GmbH" {
		t.Errorf(msgFail, "Attribute", "cloudControl GmbH", at)
	}
}

func TestAddElement(t *testing.T) {
	elems := make([]string, 0)

	n1 := &html.Node{
		Type:      3,
		DataAtom:  1,
		Data:      "a",
		Namespace: "",
		Attr: []html.Attribute{
			html.Attribute{Namespace: "",
				Key: "href",
				Val: "/pricing"}}}
	n2 := &html.Node{
		Type:      3,
		DataAtom:  1,
		Data:      "a",
		Namespace: "",
		Attr: []html.Attribute{
			html.Attribute{
				Namespace: "",
				Key:       "href",
				Val:       "http://status.cloudcontrol.com"}}}
	n3 := &html.Node{
		Type:      3,
		DataAtom:  1,
		Data:      "a",
		Namespace: "",
		Attr: []html.Attribute{
			html.Attribute{
				Namespace: "",
				Key:       "href",
				Val:       "//console.cloudcontrol.com"}}}
	n4 := &html.Node{
		Type:      3,
		DataAtom:  1,
		Data:      "a",
		Namespace: "",
		Attr: []html.Attribute{
			html.Attribute{
				Namespace: "",
				Key:       "href",
				Val:       "#part"}}}

	u, err := url.Parse("http://www.cloudcontrol.com")
	if err != nil {
		t.Errorf("%v", err)
	}

	attr := "href"

	if e1 := fmt.Sprintf("%v", ExportAddElement(elems, u, n1, attr)); e1 != "[http://www.cloudcontrol.com/pricing]" {
		t.Errorf(msgFail, "Internal Link", "[http://www.cloudcontrol.com/pricing]", e1)
	}

	if e2 := fmt.Sprintf("%v", ExportAddElement(elems, u, n2, attr)); e2 != "[http://status.cloudcontrol.com]" {
		t.Errorf(msgFail, "External Link", "[http://status.cloudcontrol.com]", e2)
	}

	if e3 := fmt.Sprintf("%v", ExportAddElement(elems, u, n3, attr)); e3 != "[http://console.cloudcontrol.com]" {
		t.Errorf(msgFail, "External Link 2", "[http://console.cloudcontrol.com]", e3)
	}

	if e4 := fmt.Sprintf("%v", ExportAddElement(elems, u, n4, attr)); e4 != "[http://www.cloudcontrol.com#part]" {
		t.Errorf(msgFail, "Internal Link #", "[http://www.cloudcontrol.com#part]", e4)
	}
}

func TestTimeoutDialer(t *testing.T) {
	secs := 20
	f1 := ExportTimeoutDialer(secs)
	c, err := f1("tcp", "google.com:80")
	if err != nil {
		t.Errorf(msgFail, "Timeout Dialer", "TCPConn", c)
	}
}

func TestNewScraper(t *testing.T) {
	u, err := url.Parse("http://www.cloudcontrol.com")
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = ExportNewScraper(u, 20)
	if err != nil {
		t.Errorf(msgFail, "scraper", "Scraper", err)
	}
}

func TestHasProtocolAsPrefix(t *testing.T) {
	val1 := "http://example.com/image.jpg"
	b1 := ExportHasProtocolAsPrefix(val1)
	if !b1 {
		t.Errorf(msgFail, "HasProtocolAsPrefix", true, false)
	}

	val2 := "image.jpg"
	b2 := ExportHasProtocolAsPrefix(val2)
	if b2 {
		t.Errorf(msgFail, "HasProtocolAsPrefix", false, true)
	}

	val3 := "https://www.example.com/image.jpg"
	b3 := ExportHasProtocolAsPrefix(val3)
	if !b3 {
		t.Errorf(msgFail, "HasProtocolAsPrefix", true, false)
	}

}
