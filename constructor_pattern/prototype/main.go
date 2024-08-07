// 原型模式
package main

import (
	"bytes"
	"fmt"
)

// Node a document object model node
type Node interface {
	// Strings returns nodes text representation
	String() string
	// Parent returns the node parent
	Parent() Node
	// SetParent sets the node parent
	SetParent(node Node)
	// Children returns the node children nodes
	Children() []Node
	// AddChild adds a child node
	AddChild(child Node)
	// Clone clones a node
	Clone() Node
}

// Element represents an element in document object model
type Element struct {
	text     string
	parent   Node
	children []Node
}

// NewElement makes a new element
func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]Node, 0),
	}
}

// Parent returns the element parent
func (e *Element) Parent() Node {
	return e.parent
}

// SetParent sets the element parent
func (e *Element) SetParent(node Node) {
	e.parent = node
}

// Children returns the element children elements
func (e *Element) Children() []Node {
	return e.children
}

// AddChild adds a child element
func (e *Element) AddChild(child Node) {
	clone := child.Clone()
	clone.SetParent(e)
	e.children = append(e.children, clone)
}

// Clone makes a copy of particular element. Note that the element becomes a
// root of new orphan tree
func (e *Element) Clone() Node {
	c := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]Node, 0),
	}
	for _, child := range e.children {
		c.AddChild(child)
	}
	return c
}

// String returns string representation of element
func (e *Element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.Children() {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}

func main() {
	// 职级节点--总监
	directorNode := NewElement("Director of Engineering")
	// 职级节点--研发经理
	engManagerNode := NewElement("Engineering Manager")
	engManagerNode.AddChild(NewElement("Lead Software Engineer"))
	// 研发经理是总监的下级
	directorNode.AddChild(engManagerNode)
	directorNode.AddChild(engManagerNode)
	// 办公室经理也是总监的下级
	officeManagerNode := NewElement("Office Manager")
	directorNode.AddChild(officeManagerNode)
	fmt.Println("")
	fmt.Println("# Company Hierarchy")
	fmt.Print(directorNode)
	fmt.Println("")
	// 从研发经理节点克隆出一颗新的树
	fmt.Println("# Team Hiearachy")
	fmt.Print(engManagerNode.Clone())
}
