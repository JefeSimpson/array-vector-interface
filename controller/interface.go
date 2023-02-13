package controller

type Functions interface {
	Len() int
	Push(int)
	GetItem(int) int
	Del(int) []int
	Print()
}
