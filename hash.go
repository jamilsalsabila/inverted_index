package main

import "fmt"

type Hash struct {
	NumOfVocabs int8
	H           map[string]*Head
}

func (hash *Hash) PrintHash(tmp *Node) {
	fmt.Println(hash.NumOfVocabs)
	fmt.Println("----------- HASH --------------")
	for key, value := range hash.H {
		fmt.Printf("%s, %d, ", key, value.Total)
		tmp = value.Front
		for tmp != nil {
			fmt.Printf("{%d %d} -> ", tmp.DocId, tmp.Freq)
			tmp = tmp.Next
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------- HASH --------------")
}

func (hash *Hash) FreeHash(tmp *Head) {
	for _, value := range hash.H {
		tmp = value
		Free(tmp)
		tmp = nil
	}

}
