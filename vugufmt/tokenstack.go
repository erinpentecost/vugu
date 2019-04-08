package vugufmt

import (
	"github.com/vugu/vugu/internal/htmlx"
)

// tokenStack is a stack of nodes.
type tokenStack []*htmlx.Token

// pop pops the stack. It will panic if s is empty.
func (s *tokenStack) pop() *htmlx.Token {
	i := len(*s)
	n := (*s)[i-1]
	*s = (*s)[:i-1]
	return n
}

// push inserts a node
func (s *tokenStack) push(n *htmlx.Token) {
	i := len(*s)
	(*s) = append(*s, nil)
	(*s)[i] = n
}

// top returns the most recently pushed node, or nil if s is empty.
func (s *tokenStack) top() *htmlx.Token {
	if i := len(*s); i > 0 {
		return (*s)[i-1]
	}
	return nil
}
