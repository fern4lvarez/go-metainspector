// Copyright 2011 Jeremy Wall (jeremy@marzhillstudios.com)
// Use of this source code is governed by the Artistic License 2.0.
// That License is included in the LICENSE file.

package h5

import (
	exphtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"io"
	"strings"
)

// String form of an html nodes data.
// (eg: The Tagname for ElementNodes or text for TextNodes)
func Data(n *exphtml.Node) string {
	if n.Data == "" {
		return n.DataAtom.String()
	}
	return n.Data
}

// CloneNode makes a copy of a Node with all descendants.
func CloneNode(n *exphtml.Node) *exphtml.Node {
	clone := new(exphtml.Node)
	clone.Type = n.Type
	clone.DataAtom = n.DataAtom
	clone.Data = n.Data
	clone.Attr = make([]exphtml.Attribute, len(n.Attr))
	copy(clone.Attr, n.Attr)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nc := CloneNode(c)
		clone.AppendChild(nc)
	}
	return clone
}

func NewTree(n *exphtml.Node) Tree {
	return Tree{n}
}

// Tree represents an html5 Node tree.
type Tree struct {
	n *exphtml.Node
}

func (t Tree) Top() *exphtml.Node {
	return t.n
}

// Walk a Tree applying a given function to each Node.
func (t Tree) Walk(f func(*exphtml.Node)) {
	WalkNodes(t.n, f)
}

func (t Tree) Render(w io.Writer) error {
	return RenderNodes(w, []*exphtml.Node{t.n})
}

func (t Tree) String() string {
	return RenderNodesToString([]*exphtml.Node{t.n})
}

// Walk walks a Node with all descendants applying a given function to each one.
func WalkNodes(n *exphtml.Node, f func(*exphtml.Node)) {
	if n != nil {
		f(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			WalkNodes(c, f)
		}
	}
}

// Clone clones an html5 nodetree to get a detached copy
// the parent of the node we are cloning will not be copied.
func (t Tree) Clone() Tree {
	return Tree{CloneNode(t.n)}
}

// Text constructs a TextNode
func Text(str string) *exphtml.Node {
	return &exphtml.Node{
		Data: str,
		Type: exphtml.TextNode,
	}
}

// Anchor constructs an Anchor Node
func Anchor(url, content string) *exphtml.Node {
	a := &exphtml.Node{
		DataAtom: atom.A,
		Data:     "a",
		Type:     exphtml.ElementNode,
	}
	if url != "" {
		a.Attr = []exphtml.Attribute{exphtml.Attribute{Key: "href", Val: url}}
	}
	if content != "" {
		a.AppendChild(Text(content))
	}
	return a
}

func Div(id string, class []string, children ...*exphtml.Node) *exphtml.Node {
	var attr []exphtml.Attribute
	if id != "" {
		attr = append(attr, exphtml.Attribute{Key: "id", Val: id})
	}
	if class != nil {
		attr = append(attr, exphtml.Attribute{
			Key: "class",
			Val: strings.Join(class, " "),
		})
	}
	return Element("div", attr, children...)
}

func Element(name string, attrs []exphtml.Attribute, children ...*exphtml.Node) *exphtml.Node {
	n := &exphtml.Node{
		Data: name,
		Type: exphtml.ElementNode,
		Attr: attrs,
	}
	for _, c := range children {
		n.AppendChild(c)
	}
	return n
}
