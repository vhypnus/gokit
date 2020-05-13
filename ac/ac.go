package ac





type Node struct {

	val string

	fail *Node

	children map[string]*Node
}