type Trie struct {
    child [26]*Trie
    isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this
    for _, ch := range word {
        ch = ch - 'a'
        if node.child[ch] == nil {
            node.child[ch] = &Trie{}
        }
        node = node.child[ch]
    }
    node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
    node := this
    for _, ch := range prefix {
        ch -= 'a'
        if node.child[ch] == nil {
            return nil
        }
        node = node.child[ch]
    }
    return node
}



/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.SearchPrefix(word)
    return node != nil && node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    return this.SearchPrefix(prefix) != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */