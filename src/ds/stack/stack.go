package stack

type Stack interface {
	CreateStack() Stack
	Peak() (string, error)
	Pop() (string, error)
	Push(string)
	IsEmpty() bool
	Size() int
}
