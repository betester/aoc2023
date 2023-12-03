package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TrieNode struct {
	Parent   *TrieNode
	Link     *TrieNode
	Children map[rune]*TrieNode
	Go       map[rune]*TrieNode
	IsRoot   bool
	IsLeaf   bool
	Edge     rune
	Value    int
}

type TrieIterator struct {
	CurrentNode *TrieNode
	Root        *TrieNode
}

func (tn *TrieNode) getLink() *TrieNode {

	if tn.Link == nil {
		if tn.IsRoot {
			tn.Link = tn
		} else if tn.Parent.IsRoot {
			tn.Link = tn.Parent
		} else {
			tn.Link = tn.Parent.getLink().next(tn.Edge)
		}
	}

	return tn.Link
}

func (tn *TrieNode) next(char rune) *TrieNode {

	if _, ok := tn.Go[char]; !ok {
		if _, ok := tn.Children[char]; ok {
			tn.Go[char] = tn.Children[char]
		} else if tn.IsRoot {
			tn.Go[char] = tn
		} else {
			tn.Go[char] = tn.getLink().next(char)
		}
	}

	return tn.Go[char]
}

func (t *TrieIterator) Next(char rune) *TrieNode {
	t.CurrentNode = t.CurrentNode.next(char)
	return t.CurrentNode
}

func (t *TrieIterator) Reset() {
	t.CurrentNode = t.Root
}

type Trie struct {
	Root TrieNode
}

func (t *Trie) Insert(word string, value int) {
	t.insert([]rune(word), 0, value, t.Root)
}

func (t *Trie) insert(word []rune, idx, value int, root TrieNode) {
	if idx == len(word) {
		return
	} else if _, ok := root.Children[word[idx]]; !ok {
		nodeValue := -1

		if idx == len(word)-1 {
			nodeValue = value
		}
		root.Children[word[idx]] = &TrieNode{
			Parent:   &root,
			Children: make(map[rune]*TrieNode),
			Go:       make(map[rune]*TrieNode),
			IsRoot:   false,
			IsLeaf:   idx == len(word)-1,
			Edge:     word[idx],
			Value:    nodeValue,
		}
	}

	t.insert(word, idx+1, value, *root.Children[word[idx]])
}

func (t *Trie) GetIterator() TrieIterator {
	return TrieIterator{
		CurrentNode: &t.Root,
		Root:        &t.Root,
	}
}

func getTrebuchetScore(trie TrieIterator, input string) int {
	leftSideValue, rightSideVal := 0, 0

	for _, char := range input {
		t := trie.Next(char)
		if t.IsLeaf {
			leftSideValue = t.Value
			break
		}
	}

	trie.Reset()

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		t := trie.Next(runes[i])
		if t.IsLeaf {
			rightSideVal = t.Value
			break
		}
	}

	return 10*leftSideValue + rightSideVal
}

func getTrebuchet(it TrieIterator, input []string) int {
	total := 0
	scores := ""
	for _, text := range input {
		score := getTrebuchetScore(it, text)
		scores += strconv.Itoa(score) + "\n"
		total += score
	}

	os.WriteFile("day1o.out", []byte(scores), 0644)

	return total
}

func main() {
	trie := Trie{
		Root: TrieNode{
			Children: make(map[rune]*TrieNode),
			Go:       make(map[rune]*TrieNode),
			IsRoot:   true,
			IsLeaf:   false,
			Value:    -1,
			Parent:   nil,
		},
	}
	trie.Insert("1", 1)
	trie.Insert("2", 2)
	trie.Insert("3", 3)
	trie.Insert("4", 4)
	trie.Insert("5", 5)
	trie.Insert("6", 6)
	trie.Insert("7", 7)
	trie.Insert("8", 8)
	trie.Insert("9", 9)

	trie.Insert("one", 1)
	trie.Insert("two", 2)
	trie.Insert("three", 3)
	trie.Insert("four", 4)
	trie.Insert("five", 5)
	trie.Insert("six", 6)
	trie.Insert("seven", 7)
	trie.Insert("eight", 8)
	trie.Insert("nine", 9)

	trie.Insert("eno", 1)
	trie.Insert("owt", 2)
	trie.Insert("eerht", 3)
	trie.Insert("ruof", 4)
	trie.Insert("evif", 5)
	trie.Insert("xis", 6)
	trie.Insert("neves", 7)
	trie.Insert("thgie", 8)
	trie.Insert("enin", 9)

	file, err := os.Open("day1.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := make([]string, 0)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	fmt.Println(getTrebuchet(trie.GetIterator(), array))
}
