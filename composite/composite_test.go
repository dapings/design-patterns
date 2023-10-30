package composite

func ExampleComposite() {
	root := NewComponent(CompositeNode, "root")
	branch1 := NewComponent(CompositeNode, "branch 1")
	branch2 := NewComponent(CompositeNode, "branch 2")
	branch3 := NewComponent(CompositeNode, "branch 3")

	leaf1 := NewComponent(LeafNode, "leaf 1")
	leaf2 := NewComponent(LeafNode, "leaf 2")
	leaf3 := NewComponent(LeafNode, "leaf 3")

	root.AddChild(branch1)
	root.AddChild(branch2)
	branch1.AddChild(branch3)
	branch1.AddChild(leaf1)
	branch2.AddChild(leaf2)
	branch2.AddChild(leaf3)

	root.Print("")

	// Output:
	// +root
	//  +branch 1
	//   +branch 3
	//   -leaf 1
	//  +branch 2
	//   -leaf 2
	//   -leaf 3
}
