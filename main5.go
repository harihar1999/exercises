package main

import (
	"fmt"

)

type node struct {
	value string
	left *node
	right *node

}
func inorder(root *node){

	if root==nil {
		return
	}
	inorder(root.left)
	fmt.Println(root.value)
	inorder(root.right)

}

func postorder(root *node){

	if root==nil {
		return
	}
	postorder(root.left)
	postorder(root.right)
	fmt.Println( root.value )

}
func preorder(root *node){

	if root==nil {
		return
	}
	fmt.Println( root.value )
	preorder( root.left )
	preorder( root.right )

}


func main() {
	root:=node{"+",nil,nil}
	root.left=&node{"a",nil,nil}
	root.right=&node{"-",nil,nil}
	root.right.left=&node{"b",nil,nil}
	root.right.right=&node{"c",nil,nil}
	fmt.Println("inorder")
    inorder(&root)
	fmt.Println("postorder")
	postorder(&root)
	fmt.Println("preorder")
	preorder(&root)

}