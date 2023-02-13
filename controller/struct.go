package controller

import "fmt"

type ArrayListSample struct {
	items []int
}

type VectorSample struct {
	items []int
}

func (a *ArrayListSample) Len() int {
	return len(a.items)
}

func (v *VectorSample) Len() int {
	return len(v.items)
}

func (a *ArrayListSample) Push(data int) {
	a.items = append(a.items, data)
}

func (v *VectorSample) Push(data int) {
	v.items = append(v.items, data)
}

func (a *ArrayListSample) GetItem(index int) int {
	return a.items[index]
}

func (v *VectorSample) GetItem(index int) int {
	return v.items[index]
}

func (a *ArrayListSample) Del(index int) []int {
	return append(a.items[:index], a.items[index+1:]...)
}

func (v *VectorSample) Del(index int) []int {
	return append(v.items[:index], v.items[index+1:]...)
}

func (a *ArrayListSample) Print() {
	fmt.Print("ArrayList: {")
	for _, elem := range a.items {
		fmt.Print(elem, ", ")
	}
	fmt.Print("}\n")
}

func (v *VectorSample) Print() {
	fmt.Print("Vector: {")
	for _, elem := range v.items {
		fmt.Print(elem, ", ")
	}
	fmt.Print("}\n")
}
