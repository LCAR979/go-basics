package gee

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

