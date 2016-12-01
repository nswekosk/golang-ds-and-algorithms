package stack

import "errors"

type Node struct {
	item string
	next *Node
}

type StringLinkedListStack struct {
	first *Node
	last  *Node
	size  int
}

func (s *StringLinkedListStack) CreateStack() Stack {

	s = &StringLinkedListStack{
		first: nil,
		last:  nil,
		size:  0,
	}

	return s

}

func (s *StringLinkedListStack) Push(str string) {

	oldList := s.first

	newNode := &Node{
		item: str,
		next: oldList,
	}

	s.size++

	if s.first == nil {
		s.last = newNode
	}

	s.first = newNode
}

func (s *StringLinkedListStack) Pop() (string, error) {
	if !s.IsEmpty() {
		item := s.first.item

		s.first = s.first.next

		s.size--

		return item, nil
	}

	return "", errors.New("Stack is empty")
}

func (s *StringLinkedListStack) IsEmpty() bool {
	return s.size == 0
}

func (s *StringLinkedListStack) Size() int {

	return s.size

}

func (s *StringLinkedListStack) Peak() (string, error) {
	if !s.IsEmpty() {
		return s.last.item, nil

	}
	return "", errors.New("Stack is empty")
}
