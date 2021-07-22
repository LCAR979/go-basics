package gee

import (
	"strings"
)

type node struct {
	totalPattern string
	crtPattern   string // current node matching what
	children     []*node
	isExact      bool // if the current node is exact match
}

func (n *node) firstMatchInChildren(target string) *node {
	for _, child := range n.children {
		if child.crtPattern == target || !child.isExact {
			return child
		}
	}
	return nil
}

func (n *node) AllMatchesInChildren(target string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.crtPattern == target || !child.isExact {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(totalPattern string, parts []string, height int) {
	if len(parts) == height {
		n.totalPattern = totalPattern
		return
	}
	crtPattern := parts[height]
	child := n.firstMatchInChildren(crtPattern)
	if child == nil {
		child = &node{crtPattern: crtPattern, isExact: crtPattern[0] != '*' && crtPattern[0] != ':'}
		n.children = append(n.children, child)
	}
	child.insert(totalPattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.crtPattern, "*") {
		if n.totalPattern != "" {
			return n
		}
		return nil
	}
	crtSearch := parts[height]
	children := n.AllMatchesInChildren(crtSearch)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
