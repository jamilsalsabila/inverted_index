package main

type Node struct {
	DocId int8
	Freq  int8
	Next  *Node
}

type Head struct {
	Total int8
	Front *Node
}

func CreateNode(doc int8) *Node {
	var node *Node

	node = new(Node)
	node.DocId = doc
	node.Freq++
	node.Next = nil

	return node
}

func Traverse(front *Node, docid int8) (bool, *Node) {
	var tmp *Node

	tmp = front
	for {
		if tmp.DocId == docid {
			return true, tmp
		} else {
			if tmp.Next == nil {
				return false, tmp
			} else {
				tmp = tmp.Next
			}
		}
	}

}

func Free(front *Head) {
	var tmp *Node

	for front.Front != nil {
		tmp = front.Front
		front.Total--
		front.Front = tmp.Next
		tmp.DocId = 0
		tmp.Freq = 0
		tmp.Next = nil
	}
}
