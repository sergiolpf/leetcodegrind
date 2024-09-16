package main

import "fmt"

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	//"app"],["apple"],["beer"],["add"],["jam"],["rental"]
	//["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"]
	//"jam"],["apps"],["app"],["ad"],["applepie"],["rest"],["jan"],["rent"],["beer"],["jam"]]

	//output: false,true,false,false,false,false,false,true,true,false,true,true,false,false,false,true,true,true]
	obj := Constructor()
	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	fmt.Println("search apps: ", obj.Search("apps"))
	fmt.Println("search app: ", obj.Search("app"))
	fmt.Println("search ad: ", obj.Search("ad"))
	fmt.Println("search applepie: ", obj.Search("applepie"))
	fmt.Println("search rest: ", obj.Search("rest"))
	fmt.Println("search jan: ", obj.Search("jan"))
	fmt.Println("search rent: ", obj.Search("rent"))
	fmt.Println("search beer: ", obj.Search("beer"))
	fmt.Println("starts with jam: ", obj.StartsWith("jam"))
	fmt.Println("starts with apps: ", obj.StartsWith("apps"))
	fmt.Println("starts with app: ", obj.StartsWith("app"))
	fmt.Println("starts with ad: ", obj.StartsWith("ad"))
	fmt.Println("starts with applepie: ", obj.StartsWith("applepie"))
	fmt.Println("starts with rest: ", obj.StartsWith("rest"))
	fmt.Println("starts with jan: ", obj.StartsWith("jan"))
	fmt.Println("starts with rent: ", obj.StartsWith("rent"))
	fmt.Println("starts with beer: ", obj.StartsWith("beer"))
	fmt.Println("starts with jam: ", obj.StartsWith("jam"))

}

type Trie struct {
	letter rune
	next   map[rune]*Trie
	end    bool
}

func Constructor() Trie {
	return Trie{next: make(map[rune]*Trie)}

}

func (this *Trie) Insert(word string) {

	curr := this
	for _, letter := range word {

		if existingLetter, ok := curr.next[letter]; !ok {
			newLetter := Trie{
				letter: letter,
				next:   make(map[rune]*Trie),
			}
			curr.next[letter] = &newLetter
			curr = &newLetter
		} else {
			curr = existingLetter
		}
	}
	curr.end = true
}

func (this *Trie) Search(word string) bool {
	curr := this

	for _, letter := range word {
		if _, ok := curr.next[letter]; !ok {
			return false
		}
		curr = curr.next[letter]
	}

	return curr.end
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this

	for _, letter := range prefix {
		if _, ok := curr.next[letter]; !ok {
			return false
		}
		curr = curr.next[letter]
	}
	return true
}

/*
Complexity
Tempo: O(n)

*/
