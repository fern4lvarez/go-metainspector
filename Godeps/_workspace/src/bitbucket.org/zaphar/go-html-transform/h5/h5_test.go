package h5

import (
	"golang.org/x/net/html"
	"reflect"
	"testing"
)

type tester interface {
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
}

func assertEqual(t tester, val, expected interface{}) {
	if !reflect.DeepEqual(val, expected) {
		t.Errorf("val: %v not equal to %v", val, expected)
	}
}

func assertTrue(t tester, ok bool, msg string, args ...interface{}) {
	if !ok {
		t.Errorf(msg, args...)
	}
}

func assertOrDie(t tester, ok bool, msg string, args ...interface{}) {
	if !ok {
		t.Fatalf(msg, args...)
	}
}

func TestNodeClone(t *testing.T) {
	tree, err := NewFromString(
		"<html><body><a>foo</a><div>bar</div></body></html>")
	assertOrDie(t, err == nil, "error while parsing string: %s", err)
	n := tree.Clone()
	assertEqual(
		t, n.String(), tree.String())
}

func TestNodeWalk(t *testing.T) {
	tree, err := NewFromString(
		"<html><head></head><body><a>foo</a><div>bar</div></body></html>")
	assertOrDie(t, err == nil, "error while parsing string: %s", err)
	i := 0
	ns := make([]string, 0, 6)
	f := func(n *html.Node) {
		ns = append(ns, Data(n))
		i++
	}
	assertEqual(t, tree.String(),
		"<html><head></head><body><a>foo</a><div>bar</div></body></html>")
	tree.Walk(f)
	expected := 8
	assertOrDie(t, i == expected, "We didn't have %d collected %d", expected, i)
	assertEqual(t, ns[3:], []string{"body", "a", "foo", "div", "bar"})
}

//func TestSnippet(t *testing.T) {
//	p, err := NewParserFromString("<a></a><b>")
//	assertOrDie(t, err == nil, "we errored while parsing snippet %s", err)
//	assertTrue(
//		t, p.Top != nil, "We didn't get a node tree back while parsing snippet")
//	assertEqual(t, p.Tree()).String(), "<a></a><b>")
//}
