package ds

type TreeNode struct {
	Children []*TreeNode
	Parent   *TreeNode
	Value    interface{}
	Name     string
}

func NewTree(rootValue interface{}) *TreeNode {
	return &TreeNode{
		Children: make([]*TreeNode, 0),
		Value:    rootValue,
	}
}

func (n *TreeNode) GetChild(name string) *TreeNode {
	for _, c := range n.Children {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (n *TreeNode) AddChild(name string, value interface{}) *TreeNode {
	newChild := TreeNode{
		Name:     name,
		Parent:   n,
		Children: make([]*TreeNode, 0),
		Value:    value,
	}

	n.Children = append(n.Children, &newChild)
	return &newChild
}
