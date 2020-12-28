package main

import ("fmt")

func main() {
	createBtree()
}


type Node struct {
    Val string
    Left *Node
    Right *Node
}


func createBtree() {
	rootNode := Node{Val: "M", Left:nil, Right:nil}
	aNode := Node{Val: "A", Left:nil, Right:nil}
	bNode := Node{Val: "R", Left:nil, Right:nil}
	insertNode(aNode, &rootNode)
	insertNode(bNode, &rootNode)
	debugNode(&rootNode)
}

func debugNode(rootNode *Node) {
    fmt.Println("debugging", rootNode.Val)
    left := rootNode.Left.Val
    right := rootNode.Right.Val
    fmt.Println(left)
    fmt.Println(right)
}

func insertNode(aNode Node, rootNode *Node){
    if aNode.Val < rootNode.Val {
	// if this is the last node in the chain
	// then insert here
	if (rootNode.Left == nil) {
	    rootNode.Left = &aNode
	    return
	}
	left := rootNode.Left
	if aNode.Val < left.Val {
	// we need to go down one more level
	    insertNode(aNode, left)
	} else {
	    rootNode.Left = &aNode
	    aNode.Left = left
	}
    } else {
	if  (rootNode.Right == nil) {
	    rootNode.Right = &aNode
	    return
	}
	right := rootNode.Right
	if aNode.Val > right.Val {
	    // go down one more level till the place is found
	    insertNode(aNode, right)
	} else {
	    // of the root on the right is bigger than the current val,
	    // then this is the proper place for the node
	    rootNode.Right = &aNode
	    aNode.Right = right
	}
    }

}
