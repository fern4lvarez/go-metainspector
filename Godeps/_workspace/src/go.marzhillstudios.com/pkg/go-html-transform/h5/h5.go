// Copyright 2011 Jeremy Wall (jeremy@marzhillstudios.com)
// Use of this source code is governed by the Artistic License 2.0.
// That License is included in the LICENSE file.

/*
Package h5 implements a wrapper and DSL for golang.org/x/net/html.

    tree, err := h5.New(rdr)

    tree.Walk(func(n *Node) {
       // do something with the node
    })

    tree2 := tree.Clone()
*/
package h5

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Partial(r io.Reader) ([]*html.Node, error) {
	b := &html.Node{}
	b.Data = "body"
	b.DataAtom = atom.Body
	b.Type = html.ElementNode
	return html.ParseFragment(r, b)
}

func PartialFromString(s string) ([]*html.Node, error) {
	return Partial(strings.NewReader(s))
}

func RenderNodes(w io.Writer, ns []*html.Node) error {
	for _, n := range ns {
		err := html.Render(w, n)
		if err != nil {
			return err
		}
	}
	return nil
}

func RenderNodesToString(ns []*html.Node) string {
	buf := bytes.NewBufferString("")
	RenderNodes(buf, ns)
	return string(buf.Bytes())
}

// Construct a new h5 parser from a string
func NewFromString(s string) (*Tree, error) {
	return New(strings.NewReader(s))
}

func Children(n *html.Node) []*html.Node {
	var cs []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cs = append(cs, c)
	}
	return cs
}

// Construct a new h5 parser from a io.Reader
func New(r io.Reader) (*Tree, error) {
	n, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	if n == nil {
		return nil, fmt.Errorf("Error parsing html from reader")
	}
	return &Tree{n}, nil
}
