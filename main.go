/*
REFERENSI:
	1. https://sites.google.com/site/kevinbouge/stopwords-lists
	2. https://github.com/agonopol/go-stem
	3. https://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i
	4. https://stackoverflow.com/questions/34383705/how-do-i-compare-strings-in-golang
	5. https://elearning.unsyiah.ac.id/pluginfile.php/389951/mod_resource/content/1/slide-inverted-indeks.pdf
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var doc1 [8]string
	var doc1Tokenized [8][]string
	var doc1Stemmed [8][]string
	var doc1Clean [8][]string
	var hash = new(Hash)
	var bst = new(BST)
	var strTemp string
	var stopwords map[string]rune
	var stopwordFile = "stopwords_en.txt"
	var file *os.File
	var bufRead *bufio.Reader
	var line []byte
	var tmp *Node
	var tmpHead *Head
	var newNode *Node
	var exist bool
	var e error

	stopwords = make(map[string]rune)

	hash.H = make(map[string]*Head)
	hash.NumOfVocabs = 0

	doc1[0] = "New home sales top forecasts"
	doc1[1] = "Home sales rise in july"
	doc1[2] = "Increase in home sales in july"
	doc1[3] = "July new home sales rise"
	doc1[4] = "Breakthrough drug for schizophrenia"
	doc1[5] = "New schizophrenia drug"
	doc1[6] = "new approach for treatment of schizophrenia"
	doc1[7] = "New hopes for schizophrenia patients"

	// step 1: Tokenizing
	for i := 0; i < len(doc1); i++ {
		doc1Tokenized[i] = Tokenizer(doc1[i])
	}

	// step 2: Stemming (using Porter Algorithm)
	for i := 0; i < len(doc1Tokenized); i++ {
		// word by word
		for j := 0; j < len(doc1Tokenized[i]); j++ {
			strTemp = strTemp + string(Stem([]byte(doc1Tokenized[i][j]))) + " "
		}

		doc1Stemmed[i] = strings.Split(strings.TrimSpace(strTemp), " ")
		strTemp = ""
	}

	// step 3: Delete Common Word
	// 3.1. Load stopwords for english into memory
	file, e = os.OpenFile(stopwordFile, os.O_RDONLY, 0755)
	if e != nil {
		panic(e)
	}
	bufRead = bufio.NewReader(file)
	line, _, e = bufRead.ReadLine()
	for e == nil {
		// overcome duplicate
		if stopwords[string(line)] == 0 {
			stopwords[string(line)] = 1
		}
		line, _, e = bufRead.ReadLine()
	}
	// 3.2. remove stopword from doc (if any)
	strTemp = ""
	for i := 0; i < len(doc1Stemmed); i++ {
		// word by word
		for j := 0; j < len(doc1Stemmed[i]); j++ {
			if stopwords[doc1Stemmed[i][j]] == 0 {
				strTemp += doc1Stemmed[i][j] + " "
			}
		}
		doc1Clean[i] = strings.Split(strings.TrimSpace(strTemp), " ")
		strTemp = ""
	}

	// step 4.1. Insert into hash table
	start := time.Now()
	for i := 0; i < len(doc1Clean); i++ { // docs
		for j := 0; j < len(doc1Clean[i]); j++ { // words
			if hash.H[doc1Clean[i][j]] == nil {
				hash.H[doc1Clean[i][j]] = &Head{}
				newNode = CreateNode(int8(i))
				hash.H[doc1Clean[i][j]].Front = newNode
			} else {
				exist, tmp = Traverse(hash.H[doc1Clean[i][j]].Front, int8(i))
				if exist {
					tmp.Freq++

				} else {
					newNode = CreateNode(int8(i))
					tmp.Next = newNode
				}
			}
			hash.H[doc1Clean[i][j]].Total++
			hash.NumOfVocabs++
		}
	}
	stop := time.Since(start)
	fmt.Println("Waktu membangun Inv. Idx menggunakan Hash: ", stop.Seconds())
	// step 4.2. Insert into BST
	start = time.Now()
	bst.Head = new(HeadBST)
	for i := 0; i < len(doc1Clean); i++ { // docs
		for j := 0; j < len(doc1Clean[i]); j++ { // words
			bst.Head.First = insertBST(bst.Head.First, doc1Clean[i][j], int8(i))
		}
	}
	stop = time.Since(start)
	fmt.Println("Waktu membangun Inv. Idx menggunakan BST: ", stop.Seconds())
	// Print
	for i := 0; i < len(doc1Clean); i++ {
		fmt.Println(doc1Clean[i])
	}
	hash.PrintHash(tmp)
	fmt.Println("----------- BST --------------")
	inorderBST(bst.Head.First, tmp)
	fmt.Println("----------- BST --------------")

	// Free Memory
	hash.FreeHash(tmpHead)
	FreeBST(bst.Head.First, tmpHead)

	// proof
	hash.H["dede"] = &Head{}
	hash.PrintHash(tmp)
	inorderBST(bst.Head.First, tmp)

	file.Close()
}
