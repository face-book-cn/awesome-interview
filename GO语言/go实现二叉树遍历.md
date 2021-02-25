```GO
package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

func main() {
	nodeG := Node{data: "g", left: nil,    right: nil}
	nodeF := Node{data: "f", left: &nodeG, right: nil}
	nodeE := Node{data: "e", left: nil,    right: nil}
	nodeD := Node{data: "d", left: &nodeE, right: nil}
	nodeC := Node{data: "c", left: nil,    right: nil}
	nodeB := Node{data: "b", left: &nodeD, right: &nodeF}
	nodeA := Node{data: "a", left: &nodeB, right: &nodeC}

	preOrderRecursive(nodeA)
}
//递归遍历
func preOrderRecursive(node Node) {
	fmt.Print(node.data, " ")
	if node.left != nil {
		preOrderRecursive(*node.left)
	}
	if node.right != nil {
		preOrderRecursive(*node.right)
	}
}
```

