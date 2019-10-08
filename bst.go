package main

import (
	"fmt"
	"strings"
)

type BST struct {
	Head *HeadBST
}

type HeadBST struct {
	First  *NodeBST
	Vocabs int8
}

type NodeBST struct {
	Data  string
	Total int8
	LL    *Head
	Left  *NodeBST
	Right *NodeBST
}

func insertBST(traverseNodeBST *NodeBST, data string, docId int8) *NodeBST {
	if traverseNodeBST == nil {
		var NodeBST = new(NodeBST)
		var newNode = CreateNode(docId)
		NodeBST.Data = data
		NodeBST.Total++
		NodeBST.Left = nil
		NodeBST.Right = nil

		NodeBST.LL = &Head{}
		NodeBST.LL.Front = newNode

		return NodeBST
	} else if strings.Compare(traverseNodeBST.Data, data) == 0 {
		var tmp *Node
		var exist bool

		exist, tmp = Traverse(traverseNodeBST.LL.Front, docId)

		if exist {
			tmp.Freq++
		} else {
			var newNode *Node
			newNode = CreateNode(int8(docId))
			tmp.Next = newNode
		}
		traverseNodeBST.Total++
	} else if strings.Compare(traverseNodeBST.Data, data) < 0 {
		traverseNodeBST.Left = insertBST(traverseNodeBST.Left, data, docId)
	} else {
		traverseNodeBST.Right = insertBST(traverseNodeBST.Right, data, docId)
	}

	return traverseNodeBST
}

func deleteBST(traverseNodeBST *NodeBST, data string) (*NodeBST, bool) {
	var exist bool
	exist = true
	if traverseNodeBST == nil { // data not found
		// fmt.Println(data, "Not Found")
		exist = false
		return nil, exist
	} else if strings.Compare(traverseNodeBST.Data, data) == 0 {
		// fmt.Println(data, "Found")
		var newTraverseNodeBST *NodeBST
		newTraverseNodeBST = traverseNodeBST
		if newTraverseNodeBST.Right != nil {
			for newTraverseNodeBST.Right.Right != nil {
				newTraverseNodeBST.Right.Data, newTraverseNodeBST.Data = newTraverseNodeBST.Data, newTraverseNodeBST.Right.Data
				newTraverseNodeBST = newTraverseNodeBST.Right
			}
			//swap final
			newTraverseNodeBST.Right.Data, newTraverseNodeBST.Data = newTraverseNodeBST.Data, newTraverseNodeBST.Right.Data
			if newTraverseNodeBST.Right.Left != nil {
				var temp *NodeBST
				temp = newTraverseNodeBST.Right.Left
				newTraverseNodeBST.Right.Left = nil
				newTraverseNodeBST.Right = temp
				temp = nil
			} else {
				newTraverseNodeBST.Right = nil
			}
		} else if newTraverseNodeBST.Left != nil {
			// pindahin HeadBST nya
			traverseNodeBST = newTraverseNodeBST.Left
			newTraverseNodeBST.Left = nil
		} else {
			newTraverseNodeBST = nil
			exist = true
			return nil, exist
		}
		newTraverseNodeBST = nil
	} else if traverseNodeBST.Data >= data {
		traverseNodeBST.Left, exist = deleteBST(traverseNodeBST.Left, data)
	} else {
		traverseNodeBST.Right, exist = deleteBST(traverseNodeBST.Right, data)
	}

	return traverseNodeBST, exist
}

func inorderBST(traverseNodeBST *NodeBST, tmp *Node) {
	if traverseNodeBST == nil {
		return
	}
	inorderBST(traverseNodeBST.Left, tmp)
	fmt.Printf("%s, %d, ", traverseNodeBST.Data, traverseNodeBST.Total)
	// fmt.Println(traverseNodeBST.LL)
	tmp = traverseNodeBST.LL.Front
	for tmp != nil {
		fmt.Printf("{%d %d} -> ", tmp.DocId, tmp.Freq)
		tmp = tmp.Next
	}
	fmt.Printf("\n")
	tmp = nil
	inorderBST(traverseNodeBST.Right, tmp)

}

func (HeadBST *HeadBST) insertBST(data string, docId int8) {
	HeadBST.First = insertBST(HeadBST.First, data, docId)
	HeadBST.Vocabs++
}

func (HeadBST *HeadBST) deleteBST(data string) {
	var exist bool
	HeadBST.First, exist = deleteBST(HeadBST.First, data)
	if exist {
		HeadBST.Vocabs--
	}

}

// postorder
func FreeBST(traverseNodeBST *NodeBST, tmp *Head) {
	if traverseNodeBST == nil {
		return
	}
	FreeBST(traverseNodeBST.Left, tmp)
	FreeBST(traverseNodeBST.Right, tmp)
	tmp = traverseNodeBST.LL
	Free(tmp)
	tmp = nil
	traverseNodeBST.Data = ""
	traverseNodeBST.Total = 0
	traverseNodeBST.Left = nil
	traverseNodeBST.Right = nil
}
